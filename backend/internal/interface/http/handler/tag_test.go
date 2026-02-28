package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	domain "github.com/tokushun109/tku/backend/internal/domain/tag"
	"github.com/tokushun109/tku/backend/internal/interface/http/response"
	"github.com/tokushun109/tku/backend/internal/shared/id"
)

var tagTestUUID = id.GenerateUUID()

type stubTagUC struct {
	listRes   []*domain.Tag
	listErr   error
	createErr error
	updateErr error
	deleteErr error
}

func (s *stubTagUC) List(ctx context.Context) ([]*domain.Tag, error) {
	return s.listRes, s.listErr
}

func (s *stubTagUC) Create(ctx context.Context, name string) error {
	return s.createErr
}

func (s *stubTagUC) Update(ctx context.Context, uuid string, name string) error {
	return s.updateErr
}

func (s *stubTagUC) Delete(ctx context.Context, uuid string) error {
	return s.deleteErr
}

type tagResp struct {
	UUID string `json:"uuid"`
	Name string `json:"name"`
}

func TestTagGet(t *testing.T) {
	t.Run("有効な入力を渡したとき処理に成功する", func(t *testing.T) {

		tg := mustTag(tagTestUUID, "a")
		tgUC := &stubTagUC{listRes: []*domain.Tag{tg}}
		h := NewTagHandler(tgUC)

		req := httptest.NewRequest(http.MethodGet, "/api/tag", nil)
		rr := httptest.NewRecorder()

		h.List(rr, req)

		if rr.Code != http.StatusOK {
			t.Fatalf("expected 200, got %d", rr.Code)
		}

		var resp []tagResp
		if err := json.NewDecoder(rr.Body).Decode(&resp); err != nil {
			t.Fatalf("decode error: %v", err)
		}
		if len(resp) != 1 || resp[0].Name != "a" {
			t.Fatalf("unexpected response: %+v", resp)
		}
	})

}

func TestTagPost(t *testing.T) {
	t.Run("JSONが不正なときバリデーションエラーで失敗する", func(t *testing.T) {

		tgUC := &stubTagUC{}
		h := NewTagHandler(tgUC)

		body := bytes.NewBufferString(`{invalid}`)
		req := httptest.NewRequest(http.MethodPost, "/api/tag", body)
		rr := httptest.NewRecorder()

		h.Create(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Fatalf("expected 400, got %d", rr.Code)
		}
	})
	t.Run("有効な入力を渡したとき処理に成功する", func(t *testing.T) {

		tgUC := &stubTagUC{}
		h := NewTagHandler(tgUC)

		body := bytes.NewBufferString(`{"name":"a"}`)
		req := httptest.NewRequest(http.MethodPost, "/api/tag", body)
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

func TestTagPut(t *testing.T) {
	t.Run("JSONが不正なときバリデーションエラーで失敗する", func(t *testing.T) {

		tgUC := &stubTagUC{}
		h := NewTagHandler(tgUC)

		body := bytes.NewBufferString(`{invalid}`)
		req := httptest.NewRequest(http.MethodPut, "/api/tag/uuid", body)
		req = mux.SetURLVars(req, map[string]string{"tag_uuid": tagTestUUID})
		rr := httptest.NewRecorder()

		h.Update(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Fatalf("expected 400, got %d", rr.Code)
		}
	})
	t.Run("有効な入力を渡したとき処理に成功する", func(t *testing.T) {

		tgUC := &stubTagUC{}
		h := NewTagHandler(tgUC)

		body := bytes.NewBufferString(`{"name":"a"}`)
		req := httptest.NewRequest(http.MethodPut, "/api/tag/uuid", body)
		req = mux.SetURLVars(req, map[string]string{"tag_uuid": tagTestUUID})
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

func TestTagDelete(t *testing.T) {
	t.Run("有効な入力を渡したとき処理に成功する", func(t *testing.T) {

		tgUC := &stubTagUC{}
		h := NewTagHandler(tgUC)

		req := httptest.NewRequest(http.MethodDelete, "/api/tag/uuid", nil)
		req = mux.SetURLVars(req, map[string]string{"tag_uuid": tagTestUUID})
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

func mustTag(uuidStr, name string) *domain.Tag {
	tag, err := domain.Rebuild(1, uuidStr, name)
	if err != nil {
		panic(err)
	}
	return tag
}
