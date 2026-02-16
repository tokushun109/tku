package middleware

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/tokushun109/tku/clean-backend/internal/usecase"
)

type stubSessionUC struct{}

func (s *stubSessionUC) Validate(ctx context.Context, token string) error {
	if token == "" {
		return usecase.NewAppError(usecase.ErrUnauthorized)
	}
	return nil
}

type errorResp struct {
	Message string `json:"message"`
}

func TestAuthMiddleware_Unauthorized(t *testing.T) {
	auth := NewAuthMiddleware(&stubSessionUC{})
	h := auth.RequireSession(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	req := httptest.NewRequest(http.MethodPost, "/api/category", nil)
	rr := httptest.NewRecorder()

	h.ServeHTTP(rr, req)

	if rr.Code != http.StatusUnauthorized {
		t.Fatalf("expected 401, got %d", rr.Code)
	}

	var resp errorResp
	if err := json.NewDecoder(rr.Body).Decode(&resp); err != nil {
		t.Fatalf("decode error: %v", err)
	}
	if resp.Message == "" {
		t.Fatalf("expected error message")
	}
}

func TestAuthMiddleware_OK(t *testing.T) {
	auth := NewAuthMiddleware(&stubSessionUC{})
	h := auth.RequireSession(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	req := httptest.NewRequest(http.MethodPost, "/api/category", nil)
	req.AddCookie(&http.Cookie{Name: "__sess__", Value: "token"})
	rr := httptest.NewRecorder()

	h.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rr.Code)
	}
}
