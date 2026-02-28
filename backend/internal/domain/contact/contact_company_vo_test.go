package contact

import "testing"

func TestNewContactCompany(t *testing.T) {
	t.Run("有効な入力を渡したとき有効な値の生成に成功する", func(t *testing.T) {
		company, err := NewContactCompany("株式会社サンプル")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if company.Value() != "株式会社サンプル" {
			t.Fatalf("unexpected company: %s", company.Value())
		}
	})

	t.Run("値が空のときバリデーションエラーで失敗する", func(t *testing.T) {
		_, err := NewContactCompany("")
		if err != ErrInvalidCompany {
			t.Fatalf("expected ErrInvalidCompany, got %v", err)
		}
	})

	t.Run("値が長すぎるときバリデーションエラーで失敗する", func(t *testing.T) {
		_, err := NewContactCompany("123456789012345678901")
		if err != ErrInvalidCompany {
			t.Fatalf("expected ErrInvalidCompany, got %v", err)
		}
	})
}
