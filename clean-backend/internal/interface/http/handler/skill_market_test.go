package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/tokushun109/tku/clean-backend/internal/domain/primitive"
	domain "github.com/tokushun109/tku/clean-backend/internal/domain/skill_market"
	"github.com/tokushun109/tku/clean-backend/internal/interface/http/response"
	"github.com/tokushun109/tku/clean-backend/internal/shared/id"
)

var skillMarketTestUUID = id.GenerateUUID()

type stubSkillMarketUC struct {
	listRes   []*domain.SkillMarket
	listErr   error
	createErr error
	updateErr error
	deleteErr error
}

func (s *stubSkillMarketUC) List(ctx context.Context) ([]*domain.SkillMarket, error) {
	return s.listRes, s.listErr
}

func (s *stubSkillMarketUC) Create(ctx context.Context, name string, rawURL string, icon string) error {
	return s.createErr
}

func (s *stubSkillMarketUC) Update(ctx context.Context, uuid string, name string, rawURL string, icon string) error {
	return s.updateErr
}

func (s *stubSkillMarketUC) Delete(ctx context.Context, uuid string) error {
	return s.deleteErr
}

type skillMarketResp struct {
	UUID string `json:"uuid"`
	Name string `json:"name"`
	URL  string `json:"url"`
	Icon string `json:"icon"`
}

func TestSkillMarketGet_OK(t *testing.T) {
	s := mustSkillMarket(skillMarketTestUUID, "minne", "https://minne.com", "icon")
	skillMarketUC := &stubSkillMarketUC{listRes: []*domain.SkillMarket{s}}
	h := NewSkillMarketHandler(skillMarketUC)

	req := httptest.NewRequest(http.MethodGet, "/api/skill_market", nil)
	rr := httptest.NewRecorder()

	h.List(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rr.Code)
	}

	var resp []skillMarketResp
	if err := json.NewDecoder(rr.Body).Decode(&resp); err != nil {
		t.Fatalf("decode error: %v", err)
	}
	if len(resp) != 1 || resp[0].Name != "minne" || resp[0].URL != "https://minne.com" {
		t.Fatalf("unexpected response: %+v", resp)
	}
}

func TestSkillMarketPost_InvalidJSON(t *testing.T) {
	skillMarketUC := &stubSkillMarketUC{}
	h := NewSkillMarketHandler(skillMarketUC)

	body := bytes.NewBufferString(`{invalid}`)
	req := httptest.NewRequest(http.MethodPost, "/api/skill_market", body)
	rr := httptest.NewRecorder()

	h.Create(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Fatalf("expected 400, got %d", rr.Code)
	}
}

func TestSkillMarketPost_OK(t *testing.T) {
	skillMarketUC := &stubSkillMarketUC{}
	h := NewSkillMarketHandler(skillMarketUC)

	body := bytes.NewBufferString(`{"name":"minne","url":"https://minne.com"}`)
	req := httptest.NewRequest(http.MethodPost, "/api/skill_market", body)
	rr := httptest.NewRecorder()

	h.Create(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rr.Code)
	}
	var resp response.SuccessResponse
	if err := json.NewDecoder(rr.Body).Decode(&resp); err != nil {
		t.Fatalf("decode error: %v", err)
	}
	if !resp.Success {
		t.Fatalf("expected success true")
	}
}

func TestSkillMarketPut_InvalidJSON(t *testing.T) {
	skillMarketUC := &stubSkillMarketUC{}
	h := NewSkillMarketHandler(skillMarketUC)

	body := bytes.NewBufferString(`{invalid}`)
	req := httptest.NewRequest(http.MethodPut, "/api/skill_market/uuid", body)
	req = mux.SetURLVars(req, map[string]string{"skill_market_uuid": skillMarketTestUUID})
	rr := httptest.NewRecorder()

	h.Update(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Fatalf("expected 400, got %d", rr.Code)
	}
}

func TestSkillMarketPut_OK(t *testing.T) {
	skillMarketUC := &stubSkillMarketUC{}
	h := NewSkillMarketHandler(skillMarketUC)

	body := bytes.NewBufferString(`{"name":"minne","url":"https://minne.com"}`)
	req := httptest.NewRequest(http.MethodPut, "/api/skill_market/uuid", body)
	req = mux.SetURLVars(req, map[string]string{"skill_market_uuid": skillMarketTestUUID})
	rr := httptest.NewRecorder()

	h.Update(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rr.Code)
	}
	var resp response.SuccessResponse
	if err := json.NewDecoder(rr.Body).Decode(&resp); err != nil {
		t.Fatalf("decode error: %v", err)
	}
	if !resp.Success {
		t.Fatalf("expected success true")
	}
}

func TestSkillMarketDelete_OK(t *testing.T) {
	skillMarketUC := &stubSkillMarketUC{}
	h := NewSkillMarketHandler(skillMarketUC)

	req := httptest.NewRequest(http.MethodDelete, "/api/skill_market/uuid", nil)
	req = mux.SetURLVars(req, map[string]string{"skill_market_uuid": skillMarketTestUUID})
	rr := httptest.NewRecorder()

	h.Delete(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rr.Code)
	}
	var resp response.SuccessResponse
	if err := json.NewDecoder(rr.Body).Decode(&resp); err != nil {
		t.Fatalf("decode error: %v", err)
	}
	if !resp.Success {
		t.Fatalf("expected success true")
	}
}

func mustSkillMarket(uuidStr, name, rawURL, icon string) *domain.SkillMarket {
	u, err := primitive.NewUUID(uuidStr)
	if err != nil {
		panic(err)
	}
	n, err := domain.NewSkillMarketName(name)
	if err != nil {
		panic(err)
	}
	u2, err := primitive.NewURL(rawURL)
	if err != nil {
		panic(err)
	}
	return &domain.SkillMarket{UUID: u, Name: n, URL: u2, Icon: icon}
}
