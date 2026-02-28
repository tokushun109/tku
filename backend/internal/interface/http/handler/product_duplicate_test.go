package handler

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/tokushun109/tku/backend/internal/usecase"
)

func TestProductDuplicate(t *testing.T) {
	t.Run("不正なリクエストボディならバリデーションエラーで失敗する", func(t *testing.T) {
		uc := &stubProductUC{}
		h := NewProductHandler(uc)

		req := httptest.NewRequest(http.MethodPost, "/api/product/duplicate", strings.NewReader("{"))
		rr := httptest.NewRecorder()

		h.Duplicate(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Fatalf("expected 400, got %d", rr.Code)
		}
		if uc.duplicateCalled {
			t.Fatalf("usecase should not be called")
		}
	})

	t.Run("ユースケースがエラーを返したとき適切なステータスで失敗する", func(t *testing.T) {
		uc := &stubProductUC{duplicateErr: usecase.NewAppErrorWithMessage(usecase.ErrInternal, context.DeadlineExceeded.Error())}
		h := NewProductHandler(uc)

		req := httptest.NewRequest(http.MethodPost, "/api/product/duplicate", strings.NewReader(`{"url":"https://www.creema.jp/items/123"}`))
		rr := httptest.NewRecorder()

		h.Duplicate(rr, req)

		if rr.Code != http.StatusInternalServerError {
			t.Fatalf("expected 500, got %d", rr.Code)
		}
		if !uc.duplicateCalled {
			t.Fatalf("usecase should be called")
		}
	})

	t.Run("有効な入力を渡したとき複製に成功する", func(t *testing.T) {
		uc := &stubProductUC{}
		h := NewProductHandler(uc)

		req := httptest.NewRequest(http.MethodPost, "/api/product/duplicate", strings.NewReader(`{"url":"https://www.creema.jp/items/123"}`))
		rr := httptest.NewRecorder()

		h.Duplicate(rr, req)

		if rr.Code != http.StatusOK {
			t.Fatalf("expected 200, got %d", rr.Code)
		}
		if !uc.duplicateCalled {
			t.Fatalf("usecase should be called")
		}
		if uc.duplicateReq != "https://www.creema.jp/items/123" {
			t.Fatalf("unexpected duplicate url: %s", uc.duplicateReq)
		}
		if !strings.Contains(rr.Body.String(), `"success":true`) {
			t.Fatalf("unexpected response body: %s", rr.Body.String())
		}
	})
}
