package middleware

import (
	"log"
	"net/http"
	"time"

	"github.com/tokushun109/tku/clean-backend/internal/shared/id"
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

		level := "INFO"
		if status >= 500 {
			level = "ERROR"
		} else if status >= 400 {
			level = "WARN"
		}

		log.Printf("[%s] request_id=%s method=%s path=%s status=%d latency_ms=%d",
			level, reqID, r.Method, r.URL.Path, status, latency,
		)
	})
}

func getOrCreateRequestID(r *http.Request) string {
	if v := r.Header.Get("X-Request-ID"); v != "" {
		return v
	}
	newID := id.NewUUID().String()
	r.Header.Set("X-Request-ID", newID)
	return newID
}
