package handler

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

type stubHealthUC struct {
	err error
}

func (s *stubHealthUC) Check(ctx context.Context) error {
	return s.err
}

type errorResp struct {
	Message string `json:"message"`
}

func TestHealthCheck_OK(t *testing.T) {
	uc := &stubHealthUC{}
	h := NewHealthHandler(uc)

	req := httptest.NewRequest(http.MethodGet, "/api/health_check", nil)
	rr := httptest.NewRecorder()

	h.Check(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rr.Code)
	}
	var resp successResp
	if err := json.NewDecoder(rr.Body).Decode(&resp); err != nil {
		t.Fatalf("decode error: %v", err)
	}
	if !resp.Success {
		t.Fatalf("expected success true")
	}
}

func TestHealthCheck_Error(t *testing.T) {
	uc := &stubHealthUC{err: errors.New("db down")}
	h := NewHealthHandler(uc)

	req := httptest.NewRequest(http.MethodGet, "/api/health_check", nil)
	rr := httptest.NewRecorder()

	h.Check(rr, req)

	if rr.Code != http.StatusInternalServerError {
		t.Fatalf("expected 500, got %d", rr.Code)
	}
	var resp errorResp
	if err := json.NewDecoder(rr.Body).Decode(&resp); err != nil {
		t.Fatalf("decode error: %v", err)
	}
	if resp.Message == "" {
		t.Fatalf("expected error message")
	}
}
