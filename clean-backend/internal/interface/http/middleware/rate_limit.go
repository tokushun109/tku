package middleware

import (
	"container/list"
	"net"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/tokushun109/tku/clean-backend/internal/interface/http/response"
)

const (
	defaultRateLimitPerMinute       = 5
	defaultRateLimitWindow          = time.Minute
	defaultRateLimitMaxEntries      = 10000
	defaultCleanupBatchSize         = 128
	defaultRateLimitExceededMessage = "リクエスト回数が多すぎます。しばらくしてからお試しください。"
)

type rateLimitOptions struct {
	limitPerWindow   int
	window           time.Duration
	maxEntries       int
	cleanupBatchSize int
	message          string
	keyFunc          func(*http.Request) string
}

type rateLimitEntry struct {
	windowStart time.Time
	lastSeen    time.Time
	count       int
	key         string
	node        *list.Element
}

type rateLimiter struct {
	limit            int
	window           time.Duration
	maxEntries       int
	cleanupBatchSize int
	cleanupInterval  time.Duration
	staleTTL         time.Duration
	message          string
	keyFunc          func(*http.Request) string
	now              func() time.Time

	mu          sync.Mutex
	lastCleanup time.Time
	entries     map[string]*rateLimitEntry
	lru         *list.List
}

// NewRateLimitMiddleware はIPなどのキー単位で固定ウィンドウ型のレート制限ミドルウェアを返す。
// 第1引数は1分あたりの上限回数で、省略時はデフォルト値を使用する。
func NewRateLimitMiddleware(limitPerMinute ...int) func(http.Handler) http.Handler {
	limit := defaultRateLimitPerMinute
	if len(limitPerMinute) > 0 && limitPerMinute[0] > 0 {
		limit = limitPerMinute[0]
	}

	return newRateLimitMiddlewareWithOptions(rateLimitOptions{
		limitPerWindow:   limit,
		window:           defaultRateLimitWindow,
		maxEntries:       defaultRateLimitMaxEntries,
		cleanupBatchSize: defaultCleanupBatchSize,
		message:          defaultRateLimitExceededMessage,
		keyFunc:          ClientIPKey,
	})
}

func newRateLimitMiddlewareWithOptions(opts rateLimitOptions) func(http.Handler) http.Handler {
	limiter := newRateLimiter(opts)

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			key := limiter.keyFunc(r)
			allowed, retryAfter := limiter.allow(key)
			if !allowed {
				retryAfterSeconds := int((retryAfter + time.Second - 1) / time.Second)
				if retryAfterSeconds < 1 {
					retryAfterSeconds = 1
				}
				w.Header().Set("Retry-After", strconv.Itoa(retryAfterSeconds))
				response.WriteError(w, http.StatusTooManyRequests, limiter.message)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}

func newRateLimiter(opts rateLimitOptions) *rateLimiter {
	limit := opts.limitPerWindow
	if limit <= 0 {
		limit = defaultRateLimitPerMinute
	}

	window := opts.window
	if window <= 0 {
		window = defaultRateLimitWindow
	}

	maxEntries := opts.maxEntries
	if maxEntries <= 0 {
		maxEntries = defaultRateLimitMaxEntries
	}

	cleanupBatchSize := opts.cleanupBatchSize
	if cleanupBatchSize <= 0 {
		cleanupBatchSize = defaultCleanupBatchSize
	}

	keyFunc := opts.keyFunc
	if keyFunc == nil {
		keyFunc = ClientIPKey
	}

	message := strings.TrimSpace(opts.message)
	if message == "" {
		message = defaultRateLimitExceededMessage
	}

	return &rateLimiter{
		limit:            limit,
		window:           window,
		maxEntries:       maxEntries,
		cleanupBatchSize: cleanupBatchSize,
		cleanupInterval:  window,
		staleTTL:         window * 5,
		message:          message,
		keyFunc:          keyFunc,
		now:              time.Now,
		lastCleanup:      time.Now(),
		entries:          make(map[string]*rateLimitEntry),
		lru:              list.New(),
	}
}

func (l *rateLimiter) allow(key string) (bool, time.Duration) {
	now := l.now()

	l.mu.Lock()
	defer l.mu.Unlock()

	if now.Sub(l.lastCleanup) >= l.cleanupInterval {
		l.cleanup(now)
		l.lastCleanup = now
	}

	entry, exists := l.entries[key]
	if !exists {
		if len(l.entries) >= l.maxEntries {
			return false, l.window
		}

		l.addEntry(key, now)
		return true, 0
	}

	entry.lastSeen = now
	l.touchEntry(entry)
	if now.Sub(entry.windowStart) >= l.window {
		entry.windowStart = now
		entry.count = 1
		return true, 0
	}

	if entry.count >= l.limit {
		retryAfter := l.window - now.Sub(entry.windowStart)
		if retryAfter < 0 {
			retryAfter = 0
		}
		return false, retryAfter
	}

	entry.count++
	return true, 0
}

func (l *rateLimiter) cleanup(now time.Time) {
	expireBefore := now.Add(-l.staleTTL)
	for i := 0; i < l.cleanupBatchSize; i++ {
		front := l.lru.Front()
		if front == nil {
			return
		}

		entry, ok := front.Value.(*rateLimitEntry)
		if !ok || entry == nil {
			l.lru.Remove(front)
			continue
		}
		if !entry.lastSeen.Before(expireBefore) {
			return
		}

		l.removeEntry(entry)
	}
}

func (l *rateLimiter) addEntry(key string, now time.Time) {
	entry := &rateLimitEntry{
		windowStart: now,
		lastSeen:    now,
		count:       1,
		key:         key,
	}
	entry.node = l.lru.PushBack(entry)
	l.entries[key] = entry
}

func (l *rateLimiter) touchEntry(entry *rateLimitEntry) {
	if entry == nil {
		return
	}
	if entry.node == nil {
		entry.node = l.lru.PushBack(entry)
		return
	}
	l.lru.MoveToBack(entry.node)
}

func (l *rateLimiter) removeEntry(entry *rateLimitEntry) {
	if entry == nil {
		return
	}
	if entry.node != nil {
		l.lru.Remove(entry.node)
		entry.node = nil
	}
	delete(l.entries, entry.key)
}

// ClientIPKey はリクエストからクライアントIPを抽出してレート制限キーを生成する。
func ClientIPKey(r *http.Request) string {
	if ip := parseIP(strings.TrimSpace(r.Header.Get("CF-Connecting-IP"))); ip != "" {
		return ip
	}

	if ip := parseXForwardedFor(r.Header.Get("X-Forwarded-For")); ip != "" {
		return ip
	}

	if ip := parseIP(strings.TrimSpace(r.Header.Get("X-Real-IP"))); ip != "" {
		return ip
	}

	remoteAddr := strings.TrimSpace(r.RemoteAddr)
	if remoteAddr == "" {
		return "unknown"
	}

	host, _, err := net.SplitHostPort(remoteAddr)
	if err == nil {
		if ip := parseIP(host); ip != "" {
			return ip
		}
	}

	if ip := parseIP(remoteAddr); ip != "" {
		return ip
	}

	return "unknown"
}

func parseXForwardedFor(value string) string {
	for _, part := range strings.Split(value, ",") {
		if ip := parseIP(strings.TrimSpace(part)); ip != "" {
			return ip
		}
	}
	return ""
}

func parseIP(value string) string {
	ip := net.ParseIP(strings.TrimSpace(value))
	if ip == nil {
		return ""
	}
	return ip.String()
}
