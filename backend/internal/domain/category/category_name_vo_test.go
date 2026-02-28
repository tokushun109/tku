package category

import "testing"

func TestNewCategoryName(t *testing.T) {
	t.Run("有効値を渡したとき有効な値の生成に成功する", func(t *testing.T) {

		name, err := NewCategoryName("accessory")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if name.Value() != "accessory" {
			t.Fatalf("expected trimmed value, got %q", name.Value())
		}
	})
	t.Run("日本語30文字の有効値を渡したとき有効な値の生成に成功する", func(t *testing.T) {

		name, err := NewCategoryName("あいうえおかきくけこさしすせそたちつてとなにぬねのはひふへほ")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if name.Value() != "あいうえおかきくけこさしすせそたちつてとなにぬねのはひふへほ" {
			t.Fatalf("expected value, got %q", name.Value())
		}
	})
	t.Run("値が短すぎるときバリデーションエラーで失敗する", func(t *testing.T) {

		_, err := NewCategoryName("")
		if err == nil {
			t.Fatalf("expected error")
		}
		if err != ErrInvalidName {
			t.Fatalf("expected ErrInvalidName, got %v", err)
		}
	})
	t.Run("値が長すぎるときバリデーションエラーで失敗する", func(t *testing.T) {

		_, err := NewCategoryName("1234567890123456789012345678901")
		if err == nil {
			t.Fatalf("expected error")
		}
		if err != ErrInvalidName {
			t.Fatalf("expected ErrInvalidName, got %v", err)
		}
	})
	t.Run("値が日本語31文字で不正なときバリデーションエラーで失敗する", func(t *testing.T) {

		_, err := NewCategoryName("あいうえおかきくけこさしすせそたちつてとなにぬねのはひふへほま")
		if err == nil {
			t.Fatalf("expected error")
		}
		if err != ErrInvalidName {
			t.Fatalf("expected ErrInvalidName, got %v", err)
		}
	})
}
