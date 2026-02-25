package target

import (
	"testing"

	"github.com/tokushun109/tku/clean-backend/internal/domain/primitive"
)

func TestNewTarget(t *testing.T) {
	t.Run("有効な入力を渡したとき処理に成功する", func(t *testing.T) {

		uuid := mustNewUUID("11111111-1111-4111-8111-111111111111")
		tg, err := New(uuid.String(), "accessory")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if tg.Name().String() != "accessory" {
			t.Fatalf("expected name to be trimmed value, got %q", tg.Name().String())
		}
	})

}

func mustNewUUID(s string) primitive.UUID {
	u, err := primitive.NewUUID(s)
	if err != nil {
		panic(err)
	}
	return u
}
