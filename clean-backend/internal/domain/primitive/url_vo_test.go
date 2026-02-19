package primitive

import "testing"

func TestNewURL(t *testing.T) {
	t.Run("有効値を渡したとき有効な値の生成に成功する", func(t *testing.T) {

		u, err := NewURL("https://www.creema.jp/items/123")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if u.String() != "https://www.creema.jp/items/123" {
			t.Fatalf("expected value, got %q", u.String())
		}
	})
	t.Run("値が短すぎるときバリデーションエラーで失敗する", func(t *testing.T) {

		_, err := NewURL("")
		if err == nil {
			t.Fatalf("expected error")
		}
		if err != ErrInvalidURL {
			t.Fatalf("expected ErrInvalidURL, got %v", err)
		}
	})
	t.Run("URLスキームがないときバリデーションエラーで失敗する", func(t *testing.T) {

		_, err := NewURL("example.com/shop")
		if err == nil {
			t.Fatalf("expected error")
		}
		if err != ErrInvalidURL {
			t.Fatalf("expected ErrInvalidURL, got %v", err)
		}
	})
	t.Run("hostとopaqueとfragmentがないときバリデーションエラーで失敗する", func(t *testing.T) {

		_, err := NewURL("https://")
		if err == nil {
			t.Fatalf("expected error")
		}
		if err != ErrInvalidURL {
			t.Fatalf("expected ErrInvalidURL, got %v", err)
		}
	})
	t.Run("mailtoスキームの有効値を渡したとき有効なURLの生成に成功する", func(t *testing.T) {

		u, err := NewURL("mailto:test@example.com")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if u.String() != "mailto:test@example.com" {
			t.Fatalf("expected value, got %q", u.String())
		}
	})
	t.Run("値の前後に空白を含むときバリデーションエラーで失敗する", func(t *testing.T) {

		_, err := NewURL(" https://example.com ")
		if err == nil {
			t.Fatalf("expected error")
		}
		if err != ErrInvalidURL {
			t.Fatalf("expected ErrInvalidURL, got %v", err)
		}
	})
	t.Run("fileスキームでパスを含む有効値を渡したとき有効なURLの生成に成功する", func(t *testing.T) {

		u, err := NewURL("file:///tmp/file.txt")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if u.String() != "file:///tmp/file.txt" {
			t.Fatalf("expected value, got %q", u.String())
		}
	})
	t.Run("fileスキームでルートのみのときバリデーションエラーで失敗する", func(t *testing.T) {

		_, err := NewURL("file:///")
		if err == nil {
			t.Fatalf("expected error")
		}
		if err != ErrInvalidURL {
			t.Fatalf("expected ErrInvalidURL, got %v", err)
		}
	})
}
