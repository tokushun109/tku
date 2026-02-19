package middleware

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/tokushun109/tku/clean-backend/internal/domain/primitive"
	domainSession "github.com/tokushun109/tku/clean-backend/internal/domain/session"
	domainUser "github.com/tokushun109/tku/clean-backend/internal/domain/user"
	"github.com/tokushun109/tku/clean-backend/internal/usecase"
)

type stubUserUC struct {
	user *domainUser.User
}

func (s *stubUserUC) Login(ctx context.Context, email string, password string) (*domainSession.Session, error) {
	return nil, nil
}

func (s *stubUserUC) GetBySessionToken(ctx context.Context, token string) (*domainUser.User, error) {
	if token == "" {
		return nil, usecase.NewAppError(usecase.ErrUnauthorized)
	}
	return s.user, nil
}

func (s *stubUserUC) Logout(ctx context.Context, token string) error {
	return nil
}

type errorResp struct {
	Message string `json:"message"`
}

func TestAuthMiddleware_Unauthorized(t *testing.T) {
	auth := NewAuthMiddleware(&stubUserUC{})
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
	uuid, err := primitive.NewUUID("11111111-1111-4111-8111-111111111111")
	if err != nil {
		t.Fatalf("unexpected uuid error: %v", err)
	}
	auth := NewAuthMiddleware(&stubUserUC{user: &domainUser.User{ID: 1, UUID: uuid, Name: "admin", Email: "admin@example.com", IsAdmin: true}})
	h := auth.RequireSession(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authUser, ok := AuthenticatedUserFromContext(r.Context())
		if !ok {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if authUser.UserID != 1 || !authUser.IsAdmin {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
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
