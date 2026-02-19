package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/tokushun109/tku/clean-backend/internal/domain/primitive"
	domainSession "github.com/tokushun109/tku/clean-backend/internal/domain/session"
	domainUser "github.com/tokushun109/tku/clean-backend/internal/domain/user"
	"github.com/tokushun109/tku/clean-backend/internal/interface/http/middleware"
	"github.com/tokushun109/tku/clean-backend/internal/interface/http/response"
	"github.com/tokushun109/tku/clean-backend/internal/shared/id"
)

type stubUserUC struct {
	loginRes     *domainSession.Session
	loginErr     error
	logoutErr    error
	sessionToken string
}

func (s *stubUserUC) Login(ctx context.Context, email string, password string) (*domainSession.Session, error) {
	return s.loginRes, s.loginErr
}

func (s *stubUserUC) GetBySessionToken(ctx context.Context, token string) (*domainUser.User, error) {
	return nil, nil
}

func (s *stubUserUC) Logout(ctx context.Context, token string) error {
	s.sessionToken = token
	return s.logoutErr
}

type loginSessionResp struct {
	UUID string `json:"uuid"`
}

type loginUserResp struct {
	UUID    string `json:"uuid"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	IsAdmin bool   `json:"isAdmin"`
}

func TestUserLogin_InvalidJSON(t *testing.T) {
	h := NewUserHandler(&stubUserUC{})

	req := httptest.NewRequest(http.MethodPost, "/api/user/login", bytes.NewBufferString(`{invalid}`))
	rr := httptest.NewRecorder()

	h.Login(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Fatalf("expected 400, got %d", rr.Code)
	}
}

func TestUserLogin_OK(t *testing.T) {
	uuid := id.GenerateUUID()
	h := NewUserHandler(&stubUserUC{loginRes: &domainSession.Session{UUID: mustUUID(uuid)}})

	reqBody := bytes.NewBufferString(`{"Email":"admin@example.com","Password":"pass"}`)
	req := httptest.NewRequest(http.MethodPost, "/api/user/login", reqBody)
	rr := httptest.NewRecorder()

	h.Login(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rr.Code)
	}
	var resp loginSessionResp
	if err := json.NewDecoder(rr.Body).Decode(&resp); err != nil {
		t.Fatalf("decode error: %v", err)
	}
	if resp.UUID != uuid {
		t.Fatalf("expected uuid %s, got %s", uuid, resp.UUID)
	}
}

func TestGetLoginUser_Unauthorized_NoContext(t *testing.T) {
	h := NewUserHandler(&stubUserUC{})

	req := httptest.NewRequest(http.MethodGet, "/api/user/login", nil)
	rr := httptest.NewRecorder()

	h.GetLoginUser(rr, req)

	if rr.Code != http.StatusUnauthorized {
		t.Fatalf("expected 401, got %d", rr.Code)
	}
}

func TestGetLoginUser_OK(t *testing.T) {
	h := NewUserHandler(&stubUserUC{})

	req := httptest.NewRequest(http.MethodGet, "/api/user/login", nil)
	req = req.WithContext(middleware.ContextWithAuthenticatedUser(req.Context(), middleware.AuthenticatedUser{
		UserID:       1,
		UUID:         "11111111-1111-4111-8111-111111111111",
		Name:         "admin",
		Email:        "admin@example.com",
		IsAdmin:      true,
		SessionToken: "token",
	}))
	rr := httptest.NewRecorder()

	h.GetLoginUser(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rr.Code)
	}

	var resp loginUserResp
	if err := json.NewDecoder(rr.Body).Decode(&resp); err != nil {
		t.Fatalf("decode error: %v", err)
	}
	if resp.Email != "admin@example.com" || !resp.IsAdmin {
		t.Fatalf("unexpected response: %+v", resp)
	}
}

func TestLogout_Unauthorized_NoContext(t *testing.T) {
	h := NewUserHandler(&stubUserUC{})

	req := httptest.NewRequest(http.MethodPost, "/api/user/logout", nil)
	rr := httptest.NewRecorder()

	h.Logout(rr, req)

	if rr.Code != http.StatusUnauthorized {
		t.Fatalf("expected 401, got %d", rr.Code)
	}
}

func TestLogout_OK(t *testing.T) {
	stub := &stubUserUC{}
	h := NewUserHandler(stub)

	req := httptest.NewRequest(http.MethodPost, "/api/user/logout", nil)
	req = req.WithContext(middleware.ContextWithAuthenticatedUser(req.Context(), middleware.AuthenticatedUser{
		UserID:       1,
		UUID:         "11111111-1111-4111-8111-111111111111",
		Name:         "admin",
		Email:        "admin@example.com",
		IsAdmin:      true,
		SessionToken: "token",
	}))
	rr := httptest.NewRecorder()

	h.Logout(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rr.Code)
	}
	if stub.sessionToken != "token" {
		t.Fatalf("expected token to be passed to usecase")
	}
	var resp response.SuccessResponse
	if err := json.NewDecoder(rr.Body).Decode(&resp); err != nil {
		t.Fatalf("decode error: %v", err)
	}
	if !resp.Success {
		t.Fatalf("expected success true")
	}
}

func mustUUID(s string) primitive.UUID {
	u, err := primitive.NewUUID(s)
	if err != nil {
		panic(err)
	}
	return u
}
