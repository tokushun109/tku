package handler

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/tokushun109/tku/backend/internal/interface/http/response"
)

type stubHealthUC struct {
	err error
}

func (s *stubHealthUC) Check(ctx context.Context) error {
	return s.err
}

func TestHealthCheck(t *testing.T) {
	t.Run("有効な入力を渡したとき処理に成功する", func(t *testing.T) {

		uc := &stubHealthUC{}
		h := NewHealthHandler(uc)

		req := httptest.NewRequest(http.MethodGet, "/api/health_check", nil)
		rr := httptest.NewRecorder()

		h.Check(rr, req)

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
	t.Run("依存処理が失敗したならエラーを返す", func(t *testing.T) {

		uc := &stubHealthUC{err: errors.New("db down")}
		h := NewHealthHandler(uc)

		req := httptest.NewRequest(http.MethodGet, "/api/health_check", nil)
		rr := httptest.NewRecorder()

		h.Check(rr, req)

		if rr.Code != http.StatusInternalServerError {
			t.Fatalf("expected 500, got %d", rr.Code)
		}
		var resp response.ErrorResponse
		if err := json.NewDecoder(rr.Body).Decode(&resp); err != nil {
			t.Fatalf("decode error: %v", err)
		}
		if resp.Message == "" {
			t.Fatalf("expected error message")
		}
	})
}
