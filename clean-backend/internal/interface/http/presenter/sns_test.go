package presenter

import (
	"testing"

	"github.com/tokushun109/tku/clean-backend/internal/domain/primitive"
	domain "github.com/tokushun109/tku/clean-backend/internal/domain/sns"
	"github.com/tokushun109/tku/clean-backend/internal/shared/id"
)

func TestToSnsResponse(t *testing.T) {
	t.Run("Snsを渡したならSnsResponseを返す", func(t *testing.T) {

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

		sns, err := domain.Rebuild(1, u.String(), n.String(), snsURL.String(), "icon")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		res := ToSnsResponse(sns)
		if res.UUID != u.String() || res.Name != n.String() || res.URL != snsURL.String() || res.Icon != "icon" {
			t.Fatalf("unexpected response: %+v", res)
		}
	})

}

func TestToSnsResponses(t *testing.T) {
	t.Run("Sns配列を渡したならSnsResponse配列を返す", func(t *testing.T) {

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

		sns1, err := domain.Rebuild(1, u1.String(), n1.String(), url1.String(), "icon1")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		sns2, err := domain.Rebuild(2, u2.String(), n2.String(), url2.String(), "icon2")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		res := ToSnsResponses([]*domain.Sns{sns1, sns2})
		if len(res) != 2 {
			t.Fatalf("expected 2, got %d", len(res))
		}
		if res[0].Name != "a" || res[1].Name != "b" {
			t.Fatalf("unexpected response: %+v", res)
		}
	})

}
