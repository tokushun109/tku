package contact

import (
	"strings"
	"testing"
)

func TestNewContactName(t *testing.T) {
	t.Run("有効な入力を渡したとき有効な値の生成に成功する", func(t *testing.T) {
		name, err := NewContactName("山田太郎")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if name.String() != "山田太郎" {
			t.Fatalf("unexpected name: %s", name.String())
		}
	})

	t.Run("値が空のときバリデーションエラーで失敗する", func(t *testing.T) {
		_, err := NewContactName("")
		if err != ErrInvalidName {
			t.Fatalf("expected ErrInvalidName, got %v", err)
		}
	})

	t.Run("値が最大長30文字のとき有効な値の生成に成功する", func(t *testing.T) {
		maxLenName := strings.Repeat("あ", 30)
		name, err := NewContactName(maxLenName)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if name.String() != maxLenName {
			t.Fatalf("unexpected name: %s", name.String())
		}
	})

	t.Run("値が長すぎるときバリデーションエラーで失敗する", func(t *testing.T) {
		tooLongName := strings.Repeat("あ", 31)
		_, err := NewContactName(tooLongName)
		if err != ErrInvalidName {
			t.Fatalf("expected ErrInvalidName, got %v", err)
		}
	})
}
