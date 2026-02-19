package target

import "testing"

func TestNewTarget(t *testing.T) {
	t.Run("有効な入力を渡したとき処理に成功する", func(t *testing.T) {

		tg, err := New("accessory")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if tg.Name.String() != "accessory" {
			t.Fatalf("expected name to be trimmed value, got %q", tg.Name.String())
		}
	})

}
