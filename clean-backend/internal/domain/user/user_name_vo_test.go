package user

import "testing"

func TestNewUserName(t *testing.T) {
	t.Run("有効な入力を渡したとき有効な値の生成に成功する", func(t *testing.T) {
		name, err := NewUserName("山田太郎")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if name.Value() != "山田太郎" {
			t.Fatalf("unexpected name: %s", name.Value())
		}
	})

	t.Run("前後に空白を含む有効な入力を渡したときtrimされた値の生成に成功する", func(t *testing.T) {
		name, err := NewUserName("  山田太郎  ")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if name.Value() != "山田太郎" {
			t.Fatalf("unexpected name: %s", name.Value())
		}
	})

	t.Run("値が空のときバリデーションエラーで失敗する", func(t *testing.T) {
		_, err := NewUserName("")
		if err != ErrInvalidName {
			t.Fatalf("expected ErrInvalidName, got %v", err)
		}
	})

	t.Run("値が長すぎるときバリデーションエラーで失敗する", func(t *testing.T) {
		_, err := NewUserName("123456789012345678901")
		if err != ErrInvalidName {
			t.Fatalf("expected ErrInvalidName, got %v", err)
		}
	})
}
