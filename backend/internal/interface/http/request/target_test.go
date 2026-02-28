package request

import (
	"net/http"
	"net/http/httptest"
	"testing"

	usecaseTarget "github.com/tokushun109/tku/backend/internal/usecase/target"
)

func TestParseListTargetQuery(t *testing.T) {
	t.Run("allモードを指定したならallモードを返す", func(t *testing.T) {

		req := httptest.NewRequest(http.MethodGet, "/api/target?mode=all", nil)

		q, err := ParseListTargetQuery(req)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if q.Mode != usecaseTarget.ListModeAll {
			t.Fatalf("expected %q, got %q", usecaseTarget.ListModeAll, q.Mode)
		}
	})
	t.Run("usedモードを指定したならusedモードを返す", func(t *testing.T) {

		req := httptest.NewRequest(http.MethodGet, "/api/target?mode=used", nil)

		q, err := ParseListTargetQuery(req)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if q.Mode != usecaseTarget.ListModeUsed {
			t.Fatalf("expected %q, got %q", usecaseTarget.ListModeUsed, q.Mode)
		}
	})
	t.Run("値が不正なときバリデーションエラーで失敗する", func(t *testing.T) {

		req := httptest.NewRequest(http.MethodGet, "/api/target?mode=bad", nil)

		_, err := ParseListTargetQuery(req)
		if err == nil {
			t.Fatalf("expected error")
		}
	})
}
