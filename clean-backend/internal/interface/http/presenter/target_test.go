package presenter

import (
	"testing"

	domain "github.com/tokushun109/tku/clean-backend/internal/domain/target"
	"github.com/tokushun109/tku/clean-backend/internal/domain/primitive"
	"github.com/tokushun109/tku/clean-backend/internal/shared/id"
)

func TestToTargetResponse(t *testing.T) {
	u, err := primitive.NewUUID(id.GenerateUUID())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	n, err := domain.NewTargetName("name")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	res := ToTargetResponse(&domain.Target{UUID: u, Name: n})
	if res.UUID != u.String() || res.Name != n.String() {
		t.Fatalf("unexpected response: %+v", res)
	}
}

func TestToTargetResponses(t *testing.T) {
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

	res := ToTargetResponses([]*domain.Target{
		{UUID: u1, Name: n1},
		{UUID: u2, Name: n2},
	})
	if len(res) != 2 {
		t.Fatalf("expected 2, got %d", len(res))
	}
	if res[0].Name != "a" || res[1].Name != "b" {
		t.Fatalf("unexpected response: %+v", res)
	}
}
