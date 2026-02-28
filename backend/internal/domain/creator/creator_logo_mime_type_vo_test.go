package creator

import "testing"

func TestNewCreatorLogoMimeType(t *testing.T) {
	t.Run("許可されたMIMEタイプなら生成に成功する", func(t *testing.T) {
		mimeType, err := NewCreatorLogoMimeType("image/png")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if mimeType.Extension() != ".png" {
			t.Fatalf("unexpected extension: %s", mimeType.Extension())
		}
	})

	t.Run("許可されていないMIMEタイプならバリデーションエラーで失敗する", func(t *testing.T) {
		_, err := NewCreatorLogoMimeType("application/pdf")
		if err != ErrInvalidLogoMimeType {
			t.Fatalf("expected ErrInvalidLogoMimeType, got %v", err)
		}
	})
}
