package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	domain "github.com/tokushun109/tku/clean-backend/internal/domain/target"
	"github.com/tokushun109/tku/clean-backend/internal/interface/http/response"
	"github.com/tokushun109/tku/clean-backend/internal/shared/id"
)

var targetTestUUID = id.GenerateUUID()

type stubTargetUC struct {
	listRes   []*domain.Target
	listErr   error
	createErr error
	updateErr error
	deleteErr error
}

func (s *stubTargetUC) List(ctx context.Context, mode string) ([]*domain.Target, error) {
	return s.listRes, s.listErr
}

func (s *stubTargetUC) Create(ctx context.Context, name string) error {
	return s.createErr
}

func (s *stubTargetUC) Update(ctx context.Context, uuid string, name string) error {
	return s.updateErr
}

func (s *stubTargetUC) Delete(ctx context.Context, uuid string) error {
	return s.deleteErr
}

type targetResp struct {
	UUID string `json:"uuid"`
	Name string `json:"name"`
}

func TestTargetGet(t *testing.T) {
	t.Run("有効な入力を渡したとき処理に成功する", func(t *testing.T) {

		tg := mustTarget(targetTestUUID, "a")
		tgUC := &stubTargetUC{listRes: []*domain.Target{tg}}
		h := NewTargetHandler(tgUC)

		req := httptest.NewRequest(http.MethodGet, "/api/target?mode=all", nil)
		rr := httptest.NewRecorder()

		h.List(rr, req)

		if rr.Code != http.StatusOK {
			t.Fatalf("expected 200, got %d", rr.Code)
		}

		var resp []targetResp
		if err := json.NewDecoder(rr.Body).Decode(&resp); err != nil {
			t.Fatalf("decode error: %v", err)
		}
		if len(resp) != 1 || resp[0].Name != "a" {
			t.Fatalf("unexpected response: %+v", resp)
		}
	})
	t.Run("モードが不正なときバリデーションエラーで失敗する", func(t *testing.T) {

		tgUC := &stubTargetUC{}
		h := NewTargetHandler(tgUC)

		req := httptest.NewRequest(http.MethodGet, "/api/target?mode=", nil)
		rr := httptest.NewRecorder()

		h.List(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Fatalf("expected 400, got %d", rr.Code)
		}
	})
}

func TestTargetPost(t *testing.T) {
	t.Run("JSONが不正なときバリデーションエラーで失敗する", func(t *testing.T) {

		tgUC := &stubTargetUC{}
		h := NewTargetHandler(tgUC)

		body := bytes.NewBufferString(`{invalid}`)
		req := httptest.NewRequest(http.MethodPost, "/api/target", body)
		rr := httptest.NewRecorder()

		h.Create(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Fatalf("expected 400, got %d", rr.Code)
		}
	})
	t.Run("有効な入力を渡したとき処理に成功する", func(t *testing.T) {

		tgUC := &stubTargetUC{}
		h := NewTargetHandler(tgUC)

		body := bytes.NewBufferString(`{"name":"a"}`)
		req := httptest.NewRequest(http.MethodPost, "/api/target", body)
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

func TestTargetPut(t *testing.T) {
	t.Run("JSONが不正なときバリデーションエラーで失敗する", func(t *testing.T) {

		tgUC := &stubTargetUC{}
		h := NewTargetHandler(tgUC)

		body := bytes.NewBufferString(`{invalid}`)
		req := httptest.NewRequest(http.MethodPut, "/api/target/uuid", body)
		req = mux.SetURLVars(req, map[string]string{"target_uuid": targetTestUUID})
		rr := httptest.NewRecorder()

		h.Update(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Fatalf("expected 400, got %d", rr.Code)
		}
	})
	t.Run("有効な入力を渡したとき処理に成功する", func(t *testing.T) {

		tgUC := &stubTargetUC{}
		h := NewTargetHandler(tgUC)

		body := bytes.NewBufferString(`{"name":"a"}`)
		req := httptest.NewRequest(http.MethodPut, "/api/target/uuid", body)
		req = mux.SetURLVars(req, map[string]string{"target_uuid": targetTestUUID})
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

func TestTargetDelete(t *testing.T) {
	t.Run("有効な入力を渡したとき処理に成功する", func(t *testing.T) {

		tgUC := &stubTargetUC{}
		h := NewTargetHandler(tgUC)

		req := httptest.NewRequest(http.MethodDelete, "/api/target/uuid", nil)
		req = mux.SetURLVars(req, map[string]string{"target_uuid": targetTestUUID})
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

func mustTarget(uuidStr, name string) *domain.Target {
	target, err := domain.Rebuild(1, uuidStr, name)
	if err != nil {
		panic(err)
	}
	return target
}
