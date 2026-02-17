package presenter

import (
	"testing"

	"github.com/tokushun109/tku/clean-backend/internal/domain/primitive"
	domain "github.com/tokushun109/tku/clean-backend/internal/domain/sns"
	"github.com/tokushun109/tku/clean-backend/internal/shared/id"
)

func TestToSnsResponse(t *testing.T) {
	u, err := primitive.NewUUID(id.GenerateUUID())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	n, err := domain.NewSnsName("Instagram")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	snsURL, err := primitive.NewURL("https://www.instagram.com")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	res := ToSnsResponse(&domain.Sns{UUID: u, Name: n, URL: snsURL, Icon: "icon"})
	if res.UUID != u.String() || res.Name != n.String() || res.URL != snsURL.String() || res.Icon != "icon" {
		t.Fatalf("unexpected response: %+v", res)
	}
}

func TestToSnsResponses(t *testing.T) {
	u1, err := primitive.NewUUID(id.GenerateUUID())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	u2, err := primitive.NewUUID(id.GenerateUUID())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	n1, err := domain.NewSnsName("a")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	n2, err := domain.NewSnsName("b")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	url1, err := primitive.NewURL("https://example.com/a")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	url2, err := primitive.NewURL("https://example.com/b")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	res := ToSnsResponses([]*domain.Sns{
		{UUID: u1, Name: n1, URL: url1, Icon: "icon1"},
		{UUID: u2, Name: n2, URL: url2, Icon: "icon2"},
	})
	if len(res) != 2 {
		t.Fatalf("expected 2, got %d", len(res))
	}
	if res[0].Name != "a" || res[1].Name != "b" {
		t.Fatalf("unexpected response: %+v", res)
	}
}
