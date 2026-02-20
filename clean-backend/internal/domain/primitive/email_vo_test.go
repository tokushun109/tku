package primitive

import "testing"

func TestNewEmail(t *testing.T) {
	t.Run("有効な入力を渡したとき有効な値の生成に成功する", func(t *testing.T) {
		email, err := NewEmail("test@example.com")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if email.String() != "test@example.com" {
			t.Fatalf("unexpected email: %s", email.String())
		}
	})

	t.Run("前後に空白を含む有効な入力を渡したときtrimされた値の生成に成功する", func(t *testing.T) {
		email, err := NewEmail("  test@example.com  ")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if email.String() != "test@example.com" {
			t.Fatalf("unexpected email: %s", email.String())
		}
	})

	t.Run("メール形式が不正なときバリデーションエラーで失敗する", func(t *testing.T) {
		_, err := NewEmail("invalid-email")
		if err != ErrInvalidEmail {
			t.Fatalf("expected ErrInvalidEmail, got %v", err)
		}
	})

	t.Run("値が長すぎるときバリデーションエラーで失敗する", func(t *testing.T) {
		_, err := NewEmail("123456789012345678901234567890123456789012345678901@example.com")
		if err != ErrInvalidEmail {
			t.Fatalf("expected ErrInvalidEmail, got %v", err)
		}
	})
}
