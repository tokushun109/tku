package middleware

import (
	"bytes"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestLoggingMiddleware_LevelByStatus(t *testing.T) {
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
			var buf bytes.Buffer
			prevOut := log.Writer()
			prevFlags := log.Flags()
			log.SetOutput(&buf)
			log.SetFlags(0)
			defer func() {
				log.SetOutput(prevOut)
				log.SetFlags(prevFlags)
			}()

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
			logStr := buf.String()
			if !strings.Contains(logStr, "["+tc.level+"]") {
				t.Fatalf("expected level %s in log, got %q", tc.level, logStr)
			}
			if !strings.Contains(logStr, "request_id=fixed-id") {
				t.Fatalf("expected request_id in log, got %q", logStr)
			}
		})
	}
}

func TestLoggingMiddleware_GeneratesRequestID(t *testing.T) {
	var buf bytes.Buffer
	prevOut := log.Writer()
	prevFlags := log.Flags()
	log.SetOutput(&buf)
	log.SetFlags(0)
	defer func() {
		log.SetOutput(prevOut)
		log.SetFlags(prevFlags)
	}()

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
	logStr := buf.String()
	if !strings.Contains(logStr, "request_id=") {
		t.Fatalf("expected request_id in log, got %q", logStr)
	}
}

func TestSanitizeLogValue_RemovesCRLF(t *testing.T) {
	in := "abc\r\ndef\nghi\r"
	got := sanitizeLogValue(in, 128)
	if strings.ContainsAny(got, "\r\n") {
		t.Fatalf("expected no CR/LF, got %q", got)
	}
	if got != "abcdefghi" {
		t.Fatalf("expected sanitized string, got %q", got)
	}
}

func TestSanitizeLogValue_Truncates(t *testing.T) {
	in := strings.Repeat("a", 200)
	got := sanitizeLogValue(in, 128)
	if len(got) != 128 {
		t.Fatalf("expected length 128, got %d", len(got))
	}
}
