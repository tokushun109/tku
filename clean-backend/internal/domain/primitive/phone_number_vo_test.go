package primitive

import "testing"

func TestNewPhoneNumber(t *testing.T) {
	t.Run("有効な入力を渡したとき有効な値の生成に成功する", func(t *testing.T) {
		phoneNumber, err := NewPhoneNumber("09012345678")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if phoneNumber.String() != "09012345678" {
			t.Fatalf("unexpected phone number: %s", phoneNumber.String())
		}
	})

	t.Run("前後に空白を含む有効な入力を渡したときtrimされた値の生成に成功する", func(t *testing.T) {
		phoneNumber, err := NewPhoneNumber(" 09012345678 ")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if phoneNumber.String() != "09012345678" {
			t.Fatalf("unexpected phone number: %s", phoneNumber.String())
		}
	})

	t.Run("空文字のときバリデーションエラーで失敗する", func(t *testing.T) {
		_, err := NewPhoneNumber("")
		if err != ErrInvalidPhoneNumber {
			t.Fatalf("expected ErrInvalidPhoneNumber, got %v", err)
		}
	})

	t.Run("数字以外を含む有効な入力を渡したとき有効な値の生成に成功する", func(t *testing.T) {
		phoneNumber, err := NewPhoneNumber("090-1234-5678")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if phoneNumber.String() != "090-1234-5678" {
			t.Fatalf("unexpected phone number: %s", phoneNumber.String())
		}
	})

	t.Run("値が長すぎるときバリデーションエラーで失敗する", func(t *testing.T) {
		_, err := NewPhoneNumber("123456789012345678901")
		if err != ErrInvalidPhoneNumber {
			t.Fatalf("expected ErrInvalidPhoneNumber, got %v", err)
		}
	})
}
