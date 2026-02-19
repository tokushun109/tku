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
	domain "github.com/tokushun109/tku/clean-backend/internal/domain/sales_site"
	"github.com/tokushun109/tku/clean-backend/internal/interface/http/response"
	"github.com/tokushun109/tku/clean-backend/internal/shared/id"
)

var salesSiteTestUUID = id.GenerateUUID()

type stubSalesSiteUC struct {
	listRes   []*domain.SalesSite
	listErr   error
	createErr error
	updateErr error
	deleteErr error
}

func (s *stubSalesSiteUC) List(ctx context.Context) ([]*domain.SalesSite, error) {
	return s.listRes, s.listErr
}

func (s *stubSalesSiteUC) Create(ctx context.Context, name string, rawURL string, icon string) error {
	return s.createErr
}

func (s *stubSalesSiteUC) Update(ctx context.Context, uuid string, name string, rawURL string, icon string) error {
	return s.updateErr
}

func (s *stubSalesSiteUC) Delete(ctx context.Context, uuid string) error {
	return s.deleteErr
}

type salesSiteResp struct {
	UUID string `json:"uuid"`
	Name string `json:"name"`
	URL  string `json:"url"`
	Icon string `json:"icon"`
}

func TestSalesSiteGet(t *testing.T) {
	t.Run("有効な入力を渡したとき処理に成功する", func(t *testing.T) {

		s := mustSalesSite(salesSiteTestUUID, "Creema", "https://www.creema.jp", "icon")
		salesSiteUC := &stubSalesSiteUC{listRes: []*domain.SalesSite{s}}
		h := NewSalesSiteHandler(salesSiteUC)

		req := httptest.NewRequest(http.MethodGet, "/api/sales_site", nil)
		rr := httptest.NewRecorder()

		h.List(rr, req)

		if rr.Code != http.StatusOK {
			t.Fatalf("expected 200, got %d", rr.Code)
		}

		var resp []salesSiteResp
		if err := json.NewDecoder(rr.Body).Decode(&resp); err != nil {
			t.Fatalf("decode error: %v", err)
		}
		if len(resp) != 1 || resp[0].Name != "Creema" || resp[0].URL != "https://www.creema.jp" {
			t.Fatalf("unexpected response: %+v", resp)
		}
	})

}

func TestSalesSitePost(t *testing.T) {
	t.Run("JSONが不正なときバリデーションエラーで失敗する", func(t *testing.T) {

		salesSiteUC := &stubSalesSiteUC{}
		h := NewSalesSiteHandler(salesSiteUC)

		body := bytes.NewBufferString(`{invalid}`)
		req := httptest.NewRequest(http.MethodPost, "/api/sales_site", body)
		rr := httptest.NewRecorder()

		h.Create(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Fatalf("expected 400, got %d", rr.Code)
		}
	})
	t.Run("有効な入力を渡したとき処理に成功する", func(t *testing.T) {

		salesSiteUC := &stubSalesSiteUC{}
		h := NewSalesSiteHandler(salesSiteUC)

		body := bytes.NewBufferString(`{"name":"Creema","url":"https://www.creema.jp"}`)
		req := httptest.NewRequest(http.MethodPost, "/api/sales_site", body)
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
	})
}

func TestSalesSitePut(t *testing.T) {
	t.Run("JSONが不正なときバリデーションエラーで失敗する", func(t *testing.T) {

		salesSiteUC := &stubSalesSiteUC{}
		h := NewSalesSiteHandler(salesSiteUC)

		body := bytes.NewBufferString(`{invalid}`)
		req := httptest.NewRequest(http.MethodPut, "/api/sales_site/uuid", body)
		req = mux.SetURLVars(req, map[string]string{"sales_site_uuid": salesSiteTestUUID})
		rr := httptest.NewRecorder()

		h.Update(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Fatalf("expected 400, got %d", rr.Code)
		}
	})
	t.Run("有効な入力を渡したとき処理に成功する", func(t *testing.T) {

		salesSiteUC := &stubSalesSiteUC{}
		h := NewSalesSiteHandler(salesSiteUC)

		body := bytes.NewBufferString(`{"name":"Creema","url":"https://www.creema.jp"}`)
		req := httptest.NewRequest(http.MethodPut, "/api/sales_site/uuid", body)
		req = mux.SetURLVars(req, map[string]string{"sales_site_uuid": salesSiteTestUUID})
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
	})
}

func TestSalesSiteDelete(t *testing.T) {
	t.Run("有効な入力を渡したとき処理に成功する", func(t *testing.T) {

		salesSiteUC := &stubSalesSiteUC{}
		h := NewSalesSiteHandler(salesSiteUC)

		req := httptest.NewRequest(http.MethodDelete, "/api/sales_site/uuid", nil)
		req = mux.SetURLVars(req, map[string]string{"sales_site_uuid": salesSiteTestUUID})
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
	})

}

func mustSalesSite(uuidStr, name, rawURL, icon string) *domain.SalesSite {
	u, err := primitive.NewUUID(uuidStr)
	if err != nil {
		panic(err)
	}
	n, err := domain.NewSalesSiteName(name)
	if err != nil {
		panic(err)
	}
	u2, err := primitive.NewURL(rawURL)
	if err != nil {
		panic(err)
	}
	return &domain.SalesSite{UUID: u, Name: n, URL: u2, Icon: icon}
}
