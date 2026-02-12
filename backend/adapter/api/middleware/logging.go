package middleware

import (
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/tokushun109/tku/backend/adapter/logger"
)

type loggingResponseWriter struct {
	http.ResponseWriter
	status int
}

func (w *loggingResponseWriter) WriteHeader(code int) {
	w.status = code
	w.ResponseWriter.WriteHeader(code)
}

func NewLogger(log logger.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			lrw := &loggingResponseWriter{ResponseWriter: w, status: http.StatusOK}

			next.ServeHTTP(lrw, r)

			remoteIP := ""
			if xff := r.Header.Get("X-Forwarded-For"); xff != "" {
				remoteIP = strings.TrimSpace(strings.Split(xff, ",")[0])
			}
			if remoteIP == "" {
				if ip, _, err := net.SplitHostPort(r.RemoteAddr); err == nil {
					remoteIP = ip
				} else {
					remoteIP = r.RemoteAddr
				}
			}
			duration := time.Since(start)
			log.Infof(
				"method=%s path=%s status=%d duration=%s remote=%s",
				r.Method,
				r.URL.Path,
				lrw.status,
				duration.String(),
				remoteIP,
			)
		})
	}
}
