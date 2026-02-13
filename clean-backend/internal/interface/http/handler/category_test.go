package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	domain "github.com/tokushun109/tku/clean-backend/internal/domain/category"
	"github.com/tokushun109/tku/clean-backend/internal/shared/id"
)

var testUUID = id.NewUUID().String()

type stubCategoryUC struct {
	listRes   []*domain.Category
	listErr   error
	createErr error
}

func (s *stubCategoryUC) List(ctx context.Context, mode string) ([]*domain.Category, error) {
	return s.listRes, s.listErr
}

func (s *stubCategoryUC) Create(ctx context.Context, name string) error {
	return s.createErr
}

type categoryResp struct {
	UUID string `json:"uuid"`
	Name string `json:"name"`
}

type successResp struct {
	Success bool `json:"success"`
}

func TestCategoryGet_OK(t *testing.T) {
	cat := mustCategory(testUUID, "a")
	catUC := &stubCategoryUC{listRes: []*domain.Category{cat}}
	h := NewCategoryHandler(catUC)

	req := httptest.NewRequest(http.MethodGet, "/api/category?mode=all", nil)
	rr := httptest.NewRecorder()

	h.List(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rr.Code)
	}

	var resp []categoryResp
	if err := json.NewDecoder(rr.Body).Decode(&resp); err != nil {
		t.Fatalf("decode error: %v", err)
	}
	if len(resp) != 1 || resp[0].Name != "a" {
		t.Fatalf("unexpected response: %+v", resp)
	}
}

func TestCategoryGet_InvalidMode(t *testing.T) {
	catUC := &stubCategoryUC{}
	h := NewCategoryHandler(catUC)

	req := httptest.NewRequest(http.MethodGet, "/api/category?mode=", nil)
	rr := httptest.NewRecorder()

	h.List(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Fatalf("expected 400, got %d", rr.Code)
	}
}

func TestCategoryPost_InvalidJSON(t *testing.T) {
	catUC := &stubCategoryUC{}
	h := NewCategoryHandler(catUC)

	body := bytes.NewBufferString(`{invalid}`)
	req := httptest.NewRequest(http.MethodPost, "/api/category", body)
	rr := httptest.NewRecorder()

	h.Create(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Fatalf("expected 400, got %d", rr.Code)
	}
}

func TestCategoryPost_OK(t *testing.T) {
	catUC := &stubCategoryUC{}
	h := NewCategoryHandler(catUC)

	body := bytes.NewBufferString(`{"name":"a"}`)
	req := httptest.NewRequest(http.MethodPost, "/api/category", body)
	rr := httptest.NewRecorder()

	h.Create(rr, req)

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

func mustCategory(uuidStr, name string) *domain.Category {
	u, err := domain.ParseCategoryUUID(uuidStr)
	if err != nil {
		panic(err)
	}
	n, err := domain.NewCategoryName(name)
	if err != nil {
		panic(err)
	}
	return &domain.Category{UUID: u, Name: n}
}
