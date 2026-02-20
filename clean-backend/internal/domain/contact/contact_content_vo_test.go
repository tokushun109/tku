package contact

import "testing"

func TestNewContactContent(t *testing.T) {
	t.Run("有効な入力を渡したとき有効な値の生成に成功する", func(t *testing.T) {
		content, err := NewContactContent("お問い合わせ内容です")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if content.String() != "お問い合わせ内容です" {
			t.Fatalf("unexpected content: %s", content.String())
		}
	})

	t.Run("値が空のときバリデーションエラーで失敗する", func(t *testing.T) {
		_, err := NewContactContent("")
		if err != ErrInvalidContent {
			t.Fatalf("expected ErrInvalidContent, got %v", err)
		}
	})
}
