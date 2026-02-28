package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	domain "github.com/tokushun109/tku/backend/internal/domain/category"
	"github.com/tokushun109/tku/backend/internal/interface/http/response"
	"github.com/tokushun109/tku/backend/internal/shared/id"
)

var testUUID = id.GenerateUUID()

type stubCategoryUC struct {
	listRes   []*domain.Category
	listErr   error
	createErr error
	updateErr error
	deleteErr error
}

func (s *stubCategoryUC) List(ctx context.Context, mode string) ([]*domain.Category, error) {
	return s.listRes, s.listErr
}

func (s *stubCategoryUC) Create(ctx context.Context, name string) error {
	return s.createErr
}

func (s *stubCategoryUC) Update(ctx context.Context, uuid string, name string) error {
	return s.updateErr
}

func (s *stubCategoryUC) Delete(ctx context.Context, uuid string) error {
	return s.deleteErr
}

type categoryResp struct {
	UUID string `json:"uuid"`
	Name string `json:"name"`
}

func TestCategoryGet(t *testing.T) {
	t.Run("有効な入力を渡したとき処理に成功する", func(t *testing.T) {

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
	})
	t.Run("モードが不正なときバリデーションエラーで失敗する", func(t *testing.T) {

		catUC := &stubCategoryUC{}
		h := NewCategoryHandler(catUC)

		req := httptest.NewRequest(http.MethodGet, "/api/category?mode=", nil)
		rr := httptest.NewRecorder()

		h.List(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Fatalf("expected 400, got %d", rr.Code)
		}
	})
}

func TestCategoryPost(t *testing.T) {
	t.Run("JSONが不正なときバリデーションエラーで失敗する", func(t *testing.T) {

		catUC := &stubCategoryUC{}
		h := NewCategoryHandler(catUC)

		body := bytes.NewBufferString(`{invalid}`)
		req := httptest.NewRequest(http.MethodPost, "/api/category", body)
		rr := httptest.NewRecorder()

		h.Create(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Fatalf("expected 400, got %d", rr.Code)
		}
	})
	t.Run("有効な入力を渡したとき処理に成功する", func(t *testing.T) {

		catUC := &stubCategoryUC{}
		h := NewCategoryHandler(catUC)

		body := bytes.NewBufferString(`{"name":"a"}`)
		req := httptest.NewRequest(http.MethodPost, "/api/category", body)
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

func TestCategoryPut(t *testing.T) {
	t.Run("JSONが不正なときバリデーションエラーで失敗する", func(t *testing.T) {

		catUC := &stubCategoryUC{}
		h := NewCategoryHandler(catUC)

		body := bytes.NewBufferString(`{invalid}`)
		req := httptest.NewRequest(http.MethodPut, "/api/category/uuid", body)
		req = mux.SetURLVars(req, map[string]string{"category_uuid": testUUID})
		rr := httptest.NewRecorder()

		h.Update(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Fatalf("expected 400, got %d", rr.Code)
		}
	})
	t.Run("有効な入力を渡したとき処理に成功する", func(t *testing.T) {

		catUC := &stubCategoryUC{}
		h := NewCategoryHandler(catUC)

		body := bytes.NewBufferString(`{"name":"a"}`)
		req := httptest.NewRequest(http.MethodPut, "/api/category/uuid", body)
		req = mux.SetURLVars(req, map[string]string{"category_uuid": testUUID})
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

func TestCategoryDelete(t *testing.T) {
	t.Run("有効な入力を渡したとき処理に成功する", func(t *testing.T) {

		catUC := &stubCategoryUC{}
		h := NewCategoryHandler(catUC)

		req := httptest.NewRequest(http.MethodDelete, "/api/category/uuid", nil)
		req = mux.SetURLVars(req, map[string]string{"category_uuid": testUUID})
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

func mustCategory(uuidStr, name string) *domain.Category {
	category, err := domain.Rebuild(1, uuidStr, name)
	if err != nil {
		panic(err)
	}
	return category
}
