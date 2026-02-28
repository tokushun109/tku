package middleware

import (
	"net/http"
	"strings"
	"time"

	"github.com/tokushun109/tku/backend/internal/shared/id"
	"github.com/tokushun109/tku/backend/internal/shared/logger"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		reqID := getOrCreateRequestID(r)

		// ResponseWriterをラップして status を取得
		rw := NewResponseWriter(w)
		rw.Header().Set("X-Request-ID", reqID)

		next.ServeHTTP(rw, r)

		latency := time.Since(start).Milliseconds()
		status := rw.Status()
		path := sanitizeLogValue(r.URL.Path, 256)
		msgFormat := "request_id=%s method=%s path=%s status=%d latency_ms=%d"

		switch {
		case status >= 500:
			logger.Errorf(msgFormat, reqID, r.Method, path, status, latency)
		case status >= 400:
			logger.Warnf(msgFormat, reqID, r.Method, path, status, latency)
		default:
			logger.Infof(msgFormat, reqID, r.Method, path, status, latency)
		}
	})
}

func getOrCreateRequestID(r *http.Request) string {
	if v := r.Header.Get("X-Request-ID"); v != "" {
		return sanitizeLogValue(v, 128)
	}
	newID := id.GenerateUUID()
	r.Header.Set("X-Request-ID", newID)
	return newID
}

func sanitizeLogValue(v string, maxLen int) string {
	v = strings.ReplaceAll(v, "\r", "")
	v = strings.ReplaceAll(v, "\n", "")
	if maxLen > 0 && len(v) > maxLen {
		v = v[:maxLen]
	}
	return v
}
