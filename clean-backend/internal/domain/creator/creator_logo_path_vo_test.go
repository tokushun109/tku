package creator

import "testing"

func TestNewCreatorLogoPath(t *testing.T) {
	t.Run("相対パス形式のロゴパスなら生成に成功する", func(t *testing.T) {
		logoPath, err := NewCreatorLogoPath("img/logo/a/b/test.png")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if logoPath.String() != "img/logo/a/b/test.png" {
			t.Fatalf("unexpected logo path: %s", logoPath.String())
		}
	})

	t.Run("絶対パスのロゴパスならバリデーションエラーで失敗する", func(t *testing.T) {
		_, err := NewCreatorLogoPath("/tmp/test.png")
		if err != ErrInvalidLogoPath {
			t.Fatalf("expected ErrInvalidLogoPath, got %v", err)
		}
	})
}
