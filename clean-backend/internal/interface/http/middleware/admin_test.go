package middleware

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

type adminErrorResp struct {
	Message string `json:"message"`
}

func TestAdminMiddleware_Unauthorized_NoContext(t *testing.T) {
	admin := NewAdminMiddleware()
	h := admin.RequireAdmin(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	req := httptest.NewRequest(http.MethodPost, "/api/category", nil)
	rr := httptest.NewRecorder()

	h.ServeHTTP(rr, req)

	if rr.Code != http.StatusUnauthorized {
		t.Fatalf("expected 401, got %d", rr.Code)
	}

	var resp adminErrorResp
	if err := json.NewDecoder(rr.Body).Decode(&resp); err != nil {
		t.Fatalf("decode error: %v", err)
	}
	if resp.Message == "" {
		t.Fatalf("expected error message")
	}
}

func TestAdminMiddleware_Forbidden_NotAdmin(t *testing.T) {
	admin := NewAdminMiddleware()
	h := admin.RequireAdmin(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	req := httptest.NewRequest(http.MethodPost, "/api/category", nil)
	req = req.WithContext(ContextWithAuthenticatedUser(req.Context(), AuthenticatedUser{UserID: 1, IsAdmin: false}))
	rr := httptest.NewRecorder()

	h.ServeHTTP(rr, req)

	if rr.Code != http.StatusForbidden {
		t.Fatalf("expected 403, got %d", rr.Code)
	}
}

func TestAdminMiddleware_OK(t *testing.T) {
	admin := NewAdminMiddleware()
	h := admin.RequireAdmin(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	req := httptest.NewRequest(http.MethodPost, "/api/category", nil)
	req = req.WithContext(ContextWithAuthenticatedUser(req.Context(), AuthenticatedUser{UserID: 1, IsAdmin: true}))
	rr := httptest.NewRecorder()

	h.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rr.Code)
	}
}
