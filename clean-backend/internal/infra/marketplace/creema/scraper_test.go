package creema

import (
	"errors"
	"testing"

	"github.com/tokushun109/tku/clean-backend/internal/domain/primitive"
	"github.com/tokushun109/tku/clean-backend/internal/usecase"
)

func TestValidateProductPageURL(t *testing.T) {
	t.Run("有効なCreema商品URLを受け入れる", func(t *testing.T) {
		u, err := validateProductPageURL("https://www.creema.jp/items/123")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if u.String() != "https://www.creema.jp/items/123" {
			t.Fatalf("unexpected url: %s", u.String())
		}
	})

	t.Run("creemaを含むだけの不正なホストを拒否する", func(t *testing.T) {
		_, err := validateProductPageURL("https://creema.attacker.com/items/123")
		if err == nil || !errors.Is(err, usecase.ErrInvalidInput) {
			t.Fatalf("expected ErrInvalidInput, got %v", err)
		}
	})

	t.Run("https以外のスキームを拒否する", func(t *testing.T) {
		_, err := validateProductPageURL("http://www.creema.jp/items/123")
		if err == nil || !errors.Is(err, usecase.ErrInvalidInput) {
			t.Fatalf("expected ErrInvalidInput, got %v", err)
		}
	})

	t.Run("URL形式が不正なときはバリデーションエラーを返す", func(t *testing.T) {
		_, err := validateProductPageURL("www.creema.jp/items/123")
		if err == nil || !errors.Is(err, primitive.ErrInvalidURL) {
			t.Fatalf("expected ErrInvalidURL, got %v", err)
		}
	})
}
