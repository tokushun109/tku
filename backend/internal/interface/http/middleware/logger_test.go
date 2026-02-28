package middleware

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/tokushun109/tku/backend/internal/shared/logger"
)

func TestLoggingMiddleware(t *testing.T) {
	t.Run("ステータスコードに応じたとき適切なログレベルで出力する", func(t *testing.T) {

		cases := []struct {
			name   string
			status int
			level  string
		}{
			{name: "info", status: http.StatusOK, level: "INFO"},
			{name: "warn", status: http.StatusBadRequest, level: "WARN"},
			{name: "error", status: http.StatusInternalServerError, level: "ERROR"},
		}

		for _, tc := range cases {
			t.Run(tc.name, func(t *testing.T) {
				var stdoutBuf bytes.Buffer
				var stderrBuf bytes.Buffer
				logger.SetOutputs(&stdoutBuf, &stderrBuf)
				logger.SetFlags(0)
				defer logger.Reset()

				h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(tc.status)
				})
				wrapped := LoggingMiddleware(h)

				req := httptest.NewRequest(http.MethodGet, "/api/test", nil)
				req.Header.Set("X-Request-ID", "fixed-id")
				rr := httptest.NewRecorder()

				wrapped.ServeHTTP(rr, req)

				if got := rr.Header().Get("X-Request-ID"); got != "fixed-id" {
					t.Fatalf("expected X-Request-ID fixed-id, got %q", got)
				}
				logStr := stdoutBuf.String()
				if tc.status >= http.StatusInternalServerError {
					logStr = stderrBuf.String()
					if stdoutBuf.Len() != 0 {
						t.Fatalf("expected no stdout log, got %q", stdoutBuf.String())
					}
				} else if stderrBuf.Len() != 0 {
					t.Fatalf("expected no stderr log, got %q", stderrBuf.String())
				}
				if !strings.Contains(logStr, "["+tc.level+"]") {
					t.Fatalf("expected level %s in log, got %q", tc.level, logStr)
				}
				if !strings.Contains(logStr, "request_id=fixed-id") {
					t.Fatalf("expected request_id in log, got %q", logStr)
				}
			})
		}
	})
	t.Run("リクエストIDがないなら新しいIDを返す", func(t *testing.T) {

		var stdoutBuf bytes.Buffer
		var stderrBuf bytes.Buffer
		logger.SetOutputs(&stdoutBuf, &stderrBuf)
		logger.SetFlags(0)
		defer logger.Reset()

		h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		})
		wrapped := LoggingMiddleware(h)

		req := httptest.NewRequest(http.MethodGet, "/api/test", nil)
		rr := httptest.NewRecorder()

		wrapped.ServeHTTP(rr, req)

		if got := rr.Header().Get("X-Request-ID"); got == "" {
			t.Fatalf("expected X-Request-ID to be set")
		}
		if stderrBuf.Len() != 0 {
			t.Fatalf("expected no stderr log, got %q", stderrBuf.String())
		}
		logStr := stdoutBuf.String()
		if !strings.Contains(logStr, "request_id=") {
			t.Fatalf("expected request_id in log, got %q", logStr)
		}
	})
}

func TestSanitizeLogValue(t *testing.T) {
	t.Run("値にCRLFが含まれるなら除去した値を返す", func(t *testing.T) {

		in := "abc\r\ndef\nghi\r"
		got := sanitizeLogValue(in, 128)
		if strings.ContainsAny(got, "\r\n") {
			t.Fatalf("expected no CR/LF, got %q", got)
		}
		if got != "abcdefghi" {
			t.Fatalf("expected sanitized string, got %q", got)
		}
	})
	t.Run("値が長すぎるなら切り詰めた値を返す", func(t *testing.T) {

		in := strings.Repeat("a", 200)
		got := sanitizeLogValue(in, 128)
		if len(got) != 128 {
			t.Fatalf("expected length 128, got %d", len(got))
		}
	})
}
