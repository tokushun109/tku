package contact

import (
	"testing"

	"github.com/tokushun109/tku/backend/internal/domain/primitive"
)

func TestNewContact(t *testing.T) {
	t.Run("必須項目のみ有効な入力を渡したときエンティティ生成に成功する", func(t *testing.T) {
		contact, err := New("山田太郎", "", "", "test@example.com", "お問い合わせです")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if contact.Name().Value() != "山田太郎" {
			t.Fatalf("unexpected name: %s", contact.Name().Value())
		}
		if contact.Company() != nil {
			t.Fatalf("expected company nil, got %v", contact.Company().Value())
		}
		if contact.PhoneNumber() != nil {
			t.Fatalf("expected phone number nil, got %v", contact.PhoneNumber().Value())
		}
	})

	t.Run("メールアドレスが不正なときメール形式エラーで失敗する", func(t *testing.T) {
		_, err := New("山田太郎", "", "", "invalid-mail", "お問い合わせです")
		if err != primitive.ErrInvalidEmail {
			t.Fatalf("expected ErrInvalidEmail, got %v", err)
		}
	})

	t.Run("会社名が20文字を超えるとき会社名バリデーションエラーで失敗する", func(t *testing.T) {
		_, err := New("山田太郎", "123456789012345678901", "", "test@example.com", "お問い合わせです")
		if err != ErrInvalidCompany {
			t.Fatalf("expected ErrInvalidCompany, got %v", err)
		}
	})

	t.Run("電話番号が20文字を超えるとき電話番号バリデーションエラーで失敗する", func(t *testing.T) {
		_, err := New("山田太郎", "", "123456789012345678901", "test@example.com", "お問い合わせです")
		if err != primitive.ErrInvalidPhoneNumber {
			t.Fatalf("expected ErrInvalidPhoneNumber, got %v", err)
		}
	})

	t.Run("お問い合わせ内容が空のとき内容バリデーションエラーで失敗する", func(t *testing.T) {
		_, err := New("山田太郎", "", "", "test@example.com", "")
		if err != ErrInvalidContent {
			t.Fatalf("expected ErrInvalidContent, got %v", err)
		}
	})
}
