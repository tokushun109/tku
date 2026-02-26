package presenter

import (
	"testing"

	"github.com/tokushun109/tku/clean-backend/internal/domain/primitive"
	domain "github.com/tokushun109/tku/clean-backend/internal/domain/target"
	"github.com/tokushun109/tku/clean-backend/internal/shared/id"
)

func TestToTargetResponse(t *testing.T) {
	t.Run("Targetを渡したならTargetResponseを返す", func(t *testing.T) {

		u, err := primitive.NewUUID(id.GenerateUUID())
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		n, err := domain.NewTargetName("name")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		target, err := domain.Rebuild(1, u.Value(), n.Value())
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		res := ToTargetResponse(target)
		if res.UUID != u.Value() || res.Name != n.Value() {
			t.Fatalf("unexpected response: %+v", res)
		}
	})

}

func TestToTargetResponses(t *testing.T) {
	t.Run("Target配列を渡したならTargetResponse配列を返す", func(t *testing.T) {

		u1, err := primitive.NewUUID(id.GenerateUUID())
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		u2, err := primitive.NewUUID(id.GenerateUUID())
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		n1, err := domain.NewTargetName("a")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		n2, err := domain.NewTargetName("b")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		target1, err := domain.Rebuild(1, u1.Value(), n1.Value())
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		target2, err := domain.Rebuild(2, u2.Value(), n2.Value())
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		res := ToTargetResponses([]*domain.Target{target1, target2})
		if len(res) != 2 {
			t.Fatalf("expected 2, got %d", len(res))
		}
		if res[0].Name != "a" || res[1].Name != "b" {
			t.Fatalf("unexpected response: %+v", res)
		}
	})

}
