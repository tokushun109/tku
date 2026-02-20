package middleware

import (
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
	defaultRateLimitExceededMessage = "リクエスト回数が多すぎます。しばらくしてからお試しください。"
)

type rateLimitOptions struct {
	limitPerWindow int
	window         time.Duration
	message        string
	keyFunc        func(*http.Request) string
}

type rateLimitEntry struct {
	windowStart time.Time
	lastSeen    time.Time
	count       int
}

type rateLimiter struct {
	limit           int
	window          time.Duration
	cleanupInterval time.Duration
	staleTTL        time.Duration
	message         string
	keyFunc         func(*http.Request) string
	now             func() time.Time

	mu          sync.Mutex
	lastCleanup time.Time
	entries     map[string]*rateLimitEntry
}

// NewRateLimitMiddleware はIPなどのキー単位で固定ウィンドウ型のレート制限ミドルウェアを返す。
// 第1引数は1分あたりの上限回数で、省略時はデフォルト値を使用する。
func NewRateLimitMiddleware(limitPerMinute ...int) func(http.Handler) http.Handler {
	limit := defaultRateLimitPerMinute
	if len(limitPerMinute) > 0 && limitPerMinute[0] > 0 {
		limit = limitPerMinute[0]
	}

	return newRateLimitMiddlewareWithOptions(rateLimitOptions{
		limitPerWindow: limit,
		window:         defaultRateLimitWindow,
		message:        defaultRateLimitExceededMessage,
		keyFunc:        ClientIPKey,
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

	keyFunc := opts.keyFunc
	if keyFunc == nil {
		keyFunc = ClientIPKey
	}

	message := strings.TrimSpace(opts.message)
	if message == "" {
		message = defaultRateLimitExceededMessage
	}

	return &rateLimiter{
		limit:           limit,
		window:          window,
		cleanupInterval: window,
		staleTTL:        window * 5,
		message:         message,
		keyFunc:         keyFunc,
		now:             time.Now,
		lastCleanup:     time.Now(),
		entries:         make(map[string]*rateLimitEntry),
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
		l.entries[key] = &rateLimitEntry{
			windowStart: now,
			lastSeen:    now,
			count:       1,
		}
		return true, 0
	}

	entry.lastSeen = now
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
	for key, entry := range l.entries {
		if entry.lastSeen.Before(expireBefore) {
			delete(l.entries, key)
		}
	}
}

// ClientIPKey はリクエストからクライアントIPを抽出してレート制限キーを生成する。
func ClientIPKey(r *http.Request) string {
	if xff := strings.TrimSpace(r.Header.Get("X-Forwarded-For")); xff != "" {
		parts := strings.Split(xff, ",")
		if len(parts) > 0 {
			first := strings.TrimSpace(parts[0])
			if first != "" {
				return first
			}
		}
	}

	if xrip := strings.TrimSpace(r.Header.Get("X-Real-IP")); xrip != "" {
		return xrip
	}

	remoteAddr := strings.TrimSpace(r.RemoteAddr)
	if remoteAddr == "" {
		return "unknown"
	}

	host, _, err := net.SplitHostPort(remoteAddr)
	if err == nil && host != "" {
		return host
	}

	return remoteAddr
}
