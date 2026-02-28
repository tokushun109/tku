package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestOriginMiddleware(t *testing.T) {
	t.Run("CLIENT_URL未設定ならPOSTでもそのまま通す", func(t *testing.T) {
		origin, err := NewOriginMiddleware("")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		h := origin.RequireTrustedOrigin(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		}))

		req := httptest.NewRequest(http.MethodPost, "/api/category", nil)
		rr := httptest.NewRecorder()

		h.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK {
			t.Fatalf("expected 200, got %d", rr.Code)
		}
	})

	t.Run("許可されたOriginのPOSTは通す", func(t *testing.T) {
		origin, err := NewOriginMiddleware("https://tocoriri.com")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		h := origin.RequireTrustedOrigin(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		}))

		req := httptest.NewRequest(http.MethodPost, "/api/category", nil)
		req.Header.Set("Origin", "https://tocoriri.com")
		rr := httptest.NewRecorder()

		h.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK {
			t.Fatalf("expected 200, got %d", rr.Code)
		}
	})

	t.Run("不正なOriginのPOSTは403を返す", func(t *testing.T) {
		origin, err := NewOriginMiddleware("https://tocoriri.com")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		h := origin.RequireTrustedOrigin(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		}))

		req := httptest.NewRequest(http.MethodPost, "/api/category", nil)
		req.Header.Set("Origin", "https://evil.example")
		rr := httptest.NewRecorder()

		h.ServeHTTP(rr, req)

		if rr.Code != http.StatusForbidden {
			t.Fatalf("expected 403, got %d", rr.Code)
		}
	})

	t.Run("OriginがなくてもRefererが許可されていれば通す", func(t *testing.T) {
		origin, err := NewOriginMiddleware("https://tocoriri.com")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		h := origin.RequireTrustedOrigin(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		}))

		req := httptest.NewRequest(http.MethodDelete, "/api/category/test", nil)
		req.Header.Set("Referer", "https://tocoriri.com/admin/classification")
		rr := httptest.NewRecorder()

		h.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK {
			t.Fatalf("expected 200, got %d", rr.Code)
		}
	})

	t.Run("GETはOriginがなくても通す", func(t *testing.T) {
		origin, err := NewOriginMiddleware("https://tocoriri.com")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		h := origin.RequireTrustedOrigin(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		}))

		req := httptest.NewRequest(http.MethodGet, "/api/user/me", nil)
		rr := httptest.NewRecorder()

		h.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK {
			t.Fatalf("expected 200, got %d", rr.Code)
		}
	})
}
