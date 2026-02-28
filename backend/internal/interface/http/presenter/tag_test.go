package presenter

import (
	"testing"

	"github.com/tokushun109/tku/backend/internal/domain/primitive"
	domain "github.com/tokushun109/tku/backend/internal/domain/tag"
	"github.com/tokushun109/tku/backend/internal/shared/id"
)

func TestToTagResponse(t *testing.T) {
	t.Run("Tagを渡したならTagResponseを返す", func(t *testing.T) {

		u, err := primitive.NewUUID(id.GenerateUUID())
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		n, err := domain.NewTagName("name")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		tag, err := domain.Rebuild(1, u.Value(), n.Value())
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		res := ToTagResponse(tag)
		if res.UUID != u.Value() || res.Name != n.Value() {
			t.Fatalf("unexpected response: %+v", res)
		}
	})

}

func TestToTagResponses(t *testing.T) {
	t.Run("Tag配列を渡したならTagResponse配列を返す", func(t *testing.T) {

		u1, err := primitive.NewUUID(id.GenerateUUID())
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		u2, err := primitive.NewUUID(id.GenerateUUID())
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		n1, err := domain.NewTagName("a")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		n2, err := domain.NewTagName("b")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		tag1, err := domain.Rebuild(1, u1.Value(), n1.Value())
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		tag2, err := domain.Rebuild(2, u2.Value(), n2.Value())
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		res := ToTagResponses([]*domain.Tag{tag1, tag2})
		if len(res) != 2 {
			t.Fatalf("expected 2, got %d", len(res))
		}
		if res[0].Name != "a" || res[1].Name != "b" {
			t.Fatalf("unexpected response: %+v", res)
		}
	})

}
