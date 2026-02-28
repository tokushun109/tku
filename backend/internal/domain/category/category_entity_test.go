package category

import (
	"testing"
)

func TestNewCategory(t *testing.T) {
	t.Run("有効な入力を渡したとき処理に成功する", func(t *testing.T) {

		c, err := New("11111111-1111-4111-8111-111111111111", "accessory")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if c.Name().Value() != "accessory" {
			t.Fatalf("expected name to be trimmed value, got %q", c.Name().Value())
		}
	})

}
