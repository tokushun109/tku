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

func TestAdminMiddleware(t *testing.T) {
	t.Run("コンテキストに認証情報がないなら未認証エラーを返す", func(t *testing.T) {

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
	})
	t.Run("管理者権限がないなら権限エラーを返す", func(t *testing.T) {

		admin := NewAdminMiddleware()
		h := admin.RequireAdmin(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		}))

		req := httptest.NewRequest(http.MethodPost, "/api/category", nil)
		req = req.WithContext(ContextWithAuthenticatedUser(req.Context(), AuthenticatedUser{IsAdmin: false}))
		rr := httptest.NewRecorder()

		h.ServeHTTP(rr, req)

		if rr.Code != http.StatusForbidden {
			t.Fatalf("expected 403, got %d", rr.Code)
		}
	})
	t.Run("有効な入力を渡したとき処理に成功する", func(t *testing.T) {

		admin := NewAdminMiddleware()
		h := admin.RequireAdmin(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		}))

		req := httptest.NewRequest(http.MethodPost, "/api/category", nil)
		req = req.WithContext(ContextWithAuthenticatedUser(req.Context(), AuthenticatedUser{IsAdmin: true}))
		rr := httptest.NewRecorder()

		h.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK {
			t.Fatalf("expected 200, got %d", rr.Code)
		}
	})
}
