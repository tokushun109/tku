package middleware

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/tokushun109/tku/backend/internal/testutil"
)

func TestLoggingMiddleware(t *testing.T) {
	log := &testutil.Logger{}

	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusCreated)
	})

	wrapped := NewLogger(log)(h)

	req := httptest.NewRequest(http.MethodGet, "/api/test", nil)
	req.RemoteAddr = "127.0.0.1:1234"
	rr := httptest.NewRecorder()

	wrapped.ServeHTTP(rr, req)

	if rr.Code != http.StatusCreated {
		t.Fatalf("expected status 201, got %d", rr.Code)
	}
	if len(log.Infos) == 0 {
		t.Fatalf("expected info log to be written")
	}
	if !strings.Contains(log.Infos[0], "status=201") {
		t.Fatalf("expected log to contain status=201, got %s", log.Infos[0])
	}
	if !strings.Contains(log.Infos[0], "path=/api/test") {
		t.Fatalf("expected log to contain path, got %s", log.Infos[0])
	}
}

func TestLoggingMiddleware_XForwardedFor(t *testing.T) {
	log := &testutil.Logger{}

	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	wrapped := NewLogger(log)(h)

	req := httptest.NewRequest(http.MethodGet, "/api/test", nil)
	req.RemoteAddr = "127.0.0.1:1234"
	req.Header.Set("X-Forwarded-For", "203.0.113.10, 198.51.100.2")
	rr := httptest.NewRecorder()

	wrapped.ServeHTTP(rr, req)

	if len(log.Infos) == 0 {
		t.Fatalf("expected info log to be written")
	}
	if !strings.Contains(log.Infos[0], "remote=203.0.113.10") {
		t.Fatalf("expected log to contain forwarded remote, got %s", log.Infos[0])
	}
}

func TestLoggingMiddleware_RemoteAddrNoPort(t *testing.T) {
	log := &testutil.Logger{}

	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	wrapped := NewLogger(log)(h)

	req := httptest.NewRequest(http.MethodGet, "/api/test", nil)
	req.RemoteAddr = "203.0.113.10"
	rr := httptest.NewRecorder()

	wrapped.ServeHTTP(rr, req)

	if len(log.Infos) == 0 {
		t.Fatalf("expected info log to be written")
	}
	if !strings.Contains(log.Infos[0], "remote=203.0.113.10") {
		t.Fatalf("expected log to contain remote addr, got %s", log.Infos[0])
	}
}
