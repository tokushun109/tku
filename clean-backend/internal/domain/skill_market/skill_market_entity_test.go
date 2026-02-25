package skill_market

import (
	"testing"

	"github.com/tokushun109/tku/clean-backend/internal/domain/primitive"
)

func TestNewSkillMarket(t *testing.T) {
	t.Run("有効な入力を渡したとき処理に成功する", func(t *testing.T) {

		uuid := mustNewUUID("11111111-1111-4111-8111-111111111111")
		s, err := New(uuid.String(), "minne", "https://minne.com", "https://example.com/icon.png")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if s.Name().String() != "minne" {
			t.Fatalf("expected name to be %q, got %q", "minne", s.Name().String())
		}
		if s.URL().String() != "https://minne.com" {
			t.Fatalf("expected url to be %q, got %q", "https://minne.com", s.URL().String())
		}
		if s.Icon() != "https://example.com/icon.png" {
			t.Fatalf("expected icon to be %q, got %q", "https://example.com/icon.png", s.Icon())
		}
	})
	t.Run("名前が不正なときバリデーションエラーで失敗する", func(t *testing.T) {

		uuid := mustNewUUID("11111111-1111-4111-8111-111111111111")
		_, err := New(uuid.String(), "", "https://minne.com", "")
		if err == nil {
			t.Fatalf("expected error")
		}
		if err != ErrInvalidName {
			t.Fatalf("expected ErrInvalidName, got %v", err)
		}
	})
	t.Run("URLが不正なときバリデーションエラーで失敗する", func(t *testing.T) {

		uuid := mustNewUUID("11111111-1111-4111-8111-111111111111")
		_, err := New(uuid.String(), "minne", "not-url", "")
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
