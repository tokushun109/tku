package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	domain "github.com/tokushun109/tku/backend/internal/domain/sns"
	"github.com/tokushun109/tku/backend/internal/interface/http/response"
	"github.com/tokushun109/tku/backend/internal/shared/id"
)

var snsTestUUID = id.GenerateUUID()

type stubSnsUC struct {
	listRes   []*domain.Sns
	listErr   error
	createErr error
	updateErr error
	deleteErr error
}

func (s *stubSnsUC) List(ctx context.Context) ([]*domain.Sns, error) {
	return s.listRes, s.listErr
}

func (s *stubSnsUC) Create(ctx context.Context, name string, rawURL string) error {
	return s.createErr
}

func (s *stubSnsUC) Update(ctx context.Context, uuid string, name string, rawURL string) error {
	return s.updateErr
}

func (s *stubSnsUC) Delete(ctx context.Context, uuid string) error {
	return s.deleteErr
}

type snsResp struct {
	UUID string `json:"uuid"`
	Name string `json:"name"`
	URL  string `json:"url"`
}

func TestSnsGet(t *testing.T) {
	t.Run("有効な入力を渡したとき処理に成功する", func(t *testing.T) {

		s := mustSns(snsTestUUID, "Instagram", "https://www.instagram.com")
		snsUC := &stubSnsUC{listRes: []*domain.Sns{s}}
		h := NewSnsHandler(snsUC)

		req := httptest.NewRequest(http.MethodGet, "/api/sns", nil)
		rr := httptest.NewRecorder()

		h.List(rr, req)

		if rr.Code != http.StatusOK {
			t.Fatalf("expected 200, got %d", rr.Code)
		}

		var resp []snsResp
		if err := json.NewDecoder(rr.Body).Decode(&resp); err != nil {
			t.Fatalf("decode error: %v", err)
		}
		if len(resp) != 1 || resp[0].Name != "Instagram" || resp[0].URL != "https://www.instagram.com" {
			t.Fatalf("unexpected response: %+v", resp)
		}
	})

}

func TestSnsPost(t *testing.T) {
	t.Run("JSONが不正なときバリデーションエラーで失敗する", func(t *testing.T) {

		snsUC := &stubSnsUC{}
		h := NewSnsHandler(snsUC)

		body := bytes.NewBufferString(`{invalid}`)
		req := httptest.NewRequest(http.MethodPost, "/api/sns", body)
		rr := httptest.NewRecorder()

		h.Create(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Fatalf("expected 400, got %d", rr.Code)
		}
	})
	t.Run("有効な入力を渡したとき処理に成功する", func(t *testing.T) {

		snsUC := &stubSnsUC{}
		h := NewSnsHandler(snsUC)

		body := bytes.NewBufferString(`{"name":"Instagram","url":"https://www.instagram.com"}`)
		req := httptest.NewRequest(http.MethodPost, "/api/sns", body)
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

func TestSnsPut(t *testing.T) {
	t.Run("JSONが不正なときバリデーションエラーで失敗する", func(t *testing.T) {

		snsUC := &stubSnsUC{}
		h := NewSnsHandler(snsUC)

		body := bytes.NewBufferString(`{invalid}`)
		req := httptest.NewRequest(http.MethodPut, "/api/sns/uuid", body)
		req = mux.SetURLVars(req, map[string]string{"sns_uuid": snsTestUUID})
		rr := httptest.NewRecorder()

		h.Update(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Fatalf("expected 400, got %d", rr.Code)
		}
	})
	t.Run("有効な入力を渡したとき処理に成功する", func(t *testing.T) {

		snsUC := &stubSnsUC{}
		h := NewSnsHandler(snsUC)

		body := bytes.NewBufferString(`{"name":"Instagram","url":"https://www.instagram.com"}`)
		req := httptest.NewRequest(http.MethodPut, "/api/sns/uuid", body)
		req = mux.SetURLVars(req, map[string]string{"sns_uuid": snsTestUUID})
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

func TestSnsDelete(t *testing.T) {
	t.Run("有効な入力を渡したとき処理に成功する", func(t *testing.T) {

		snsUC := &stubSnsUC{}
		h := NewSnsHandler(snsUC)

		req := httptest.NewRequest(http.MethodDelete, "/api/sns/uuid", nil)
		req = mux.SetURLVars(req, map[string]string{"sns_uuid": snsTestUUID})
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

func mustSns(uuidStr, name, rawURL string) *domain.Sns {
	sns, err := domain.Rebuild(1, uuidStr, name, rawURL)
	if err != nil {
		panic(err)
	}
	return sns
}
