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
