package creator

import "testing"

func TestNewCreatorIntroduction(t *testing.T) {
	t.Run("1文字以上1000文字以下の紹介文なら生成に成功する", func(t *testing.T) {
		introduction, err := NewCreatorIntroduction(" handmade ")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if introduction.String() != "handmade" {
			t.Fatalf("unexpected introduction: %s", introduction.String())
		}
	})

	t.Run("空文字の紹介文ならバリデーションエラーで失敗する", func(t *testing.T) {
		_, err := NewCreatorIntroduction("")
		if err != ErrInvalidIntroduction {
			t.Fatalf("expected ErrInvalidIntroduction, got %v", err)
		}
	})
}
