package sns

import (
	"testing"

	"github.com/tokushun109/tku/backend/internal/domain/primitive"
)

func TestNewSns(t *testing.T) {
	t.Run("有効な入力を渡したとき処理に成功する", func(t *testing.T) {

		s, err := New("11111111-1111-4111-8111-111111111111", "Instagram", "https://www.instagram.com", "https://example.com/icon.png")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if s.Name().Value() != "Instagram" {
			t.Fatalf("expected name to be %q, got %q", "Instagram", s.Name().Value())
		}
		if s.URL().Value() != "https://www.instagram.com" {
			t.Fatalf("expected url to be %q, got %q", "https://www.instagram.com", s.URL().Value())
		}
		if s.Icon() != "https://example.com/icon.png" {
			t.Fatalf("expected icon to be %q, got %q", "https://example.com/icon.png", s.Icon())
		}
	})
	t.Run("名前が不正なときバリデーションエラーで失敗する", func(t *testing.T) {

		_, err := New("11111111-1111-4111-8111-111111111111", "", "https://www.instagram.com", "")
		if err == nil {
			t.Fatalf("expected error")
		}
		if err != ErrInvalidName {
			t.Fatalf("expected ErrInvalidName, got %v", err)
		}
	})
	t.Run("URLが不正なときバリデーションエラーで失敗する", func(t *testing.T) {

		_, err := New("11111111-1111-4111-8111-111111111111", "Instagram", "not-url", "")
		if err == nil {
			t.Fatalf("expected error")
		}
		if err != primitive.ErrInvalidURL {
			t.Fatalf("expected primitive.ErrInvalidURL, got %v", err)
		}
	})
}
