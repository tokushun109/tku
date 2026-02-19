package presenter

import (
	"testing"

	"github.com/tokushun109/tku/clean-backend/internal/domain/primitive"
	domain "github.com/tokushun109/tku/clean-backend/internal/domain/tag"
	"github.com/tokushun109/tku/clean-backend/internal/shared/id"
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

		res := ToTagResponse(&domain.Tag{UUID: u, Name: n})
		if res.UUID != u.String() || res.Name != n.String() {
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

		res := ToTagResponses([]*domain.Tag{
			{UUID: u1, Name: n1},
			{UUID: u2, Name: n2},
		})
		if len(res) != 2 {
			t.Fatalf("expected 2, got %d", len(res))
		}
		if res[0].Name != "a" || res[1].Name != "b" {
			t.Fatalf("unexpected response: %+v", res)
		}
	})

}
