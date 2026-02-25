package sales_site

import (
	"testing"

	"github.com/tokushun109/tku/clean-backend/internal/domain/primitive"
)

func TestNewSalesSite(t *testing.T) {
	t.Run("有効な入力を渡したとき処理に成功する", func(t *testing.T) {

		uuid := mustNewUUID("11111111-1111-4111-8111-111111111111")
		s, err := New(uuid.String(), "Creema", "https://www.creema.jp", "https://example.com/icon.png")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if s.Name().String() != "Creema" {
			t.Fatalf("expected name to be %q, got %q", "Creema", s.Name().String())
		}
		if s.URL().String() != "https://www.creema.jp" {
			t.Fatalf("expected url to be %q, got %q", "https://www.creema.jp", s.URL().String())
		}
		if s.Icon() != "https://example.com/icon.png" {
			t.Fatalf("expected icon to be %q, got %q", "https://example.com/icon.png", s.Icon())
		}
	})
	t.Run("名前が不正なときバリデーションエラーで失敗する", func(t *testing.T) {

		uuid := mustNewUUID("11111111-1111-4111-8111-111111111111")
		_, err := New(uuid.String(), "", "https://www.creema.jp", "")
		if err == nil {
			t.Fatalf("expected error")
		}
		if err != ErrInvalidName {
			t.Fatalf("expected ErrInvalidName, got %v", err)
		}
	})
	t.Run("URLが不正なときバリデーションエラーで失敗する", func(t *testing.T) {

		uuid := mustNewUUID("11111111-1111-4111-8111-111111111111")
		_, err := New(uuid.String(), "Creema", "not-url", "")
		if err == nil {
			t.Fatalf("expected error")
		}
		if err != primitive.ErrInvalidURL {
			t.Fatalf("expected primitive.ErrInvalidURL, got %v", err)
		}
	})
}

func mustNewUUID(s string) primitive.UUID {
	u, err := primitive.NewUUID(s)
	if err != nil {
		panic(err)
	}
	return u
}
