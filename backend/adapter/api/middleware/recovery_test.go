package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/tokushun109/tku/backend/internal/testutil"
)

func TestRecoveryMiddleware(t *testing.T) {
	log := &testutil.Logger{}

	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		panic("boom")
	})

	wrapped := NewRecovery(log)(h)

	req := httptest.NewRequest(http.MethodGet, "/api/test", nil)
	rr := httptest.NewRecorder()

	wrapped.ServeHTTP(rr, req)

	if rr.Code != http.StatusInternalServerError {
		t.Fatalf("expected status 500, got %d", rr.Code)
	}
	if len(log.Errors) == 0 {
		t.Fatalf("expected error log to be written")
	}
}
