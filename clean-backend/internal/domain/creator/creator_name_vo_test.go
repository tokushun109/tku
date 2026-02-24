package creator

import (
	"strings"
	"testing"
)

func TestNewCreatorName(t *testing.T) {
	t.Run("前後の空白を含む名前を渡したときトリムして生成に成功する", func(t *testing.T) {
		name, err := NewCreatorName(" とこりり ")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if name.String() != "とこりり" {
			t.Fatalf("unexpected name: %s", name.String())
		}
	})

	t.Run("30文字の名前なら生成に成功する", func(t *testing.T) {
		raw := strings.Repeat("a", 30)
		name, err := NewCreatorName(raw)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if name.String() != raw {
			t.Fatalf("unexpected name: %s", name.String())
		}
	})

	t.Run("31文字の名前ならバリデーションエラーで失敗する", func(t *testing.T) {
		_, err := NewCreatorName(strings.Repeat("a", 31))
		if err != ErrInvalidName {
			t.Fatalf("expected ErrInvalidName, got %v", err)
		}
	})

	t.Run("空文字の名前ならバリデーションエラーで失敗する", func(t *testing.T) {
		_, err := NewCreatorName("   ")
		if err != ErrInvalidName {
			t.Fatalf("expected ErrInvalidName, got %v", err)
		}
	})
}
