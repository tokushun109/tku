package presenter

import (
	"testing"

	"github.com/tokushun109/tku/clean-backend/internal/domain/primitive"
	domain "github.com/tokushun109/tku/clean-backend/internal/domain/skill_market"
	"github.com/tokushun109/tku/clean-backend/internal/shared/id"
)

func TestToSkillMarketResponse(t *testing.T) {
	t.Run("SkillMarketを渡したならSkillMarketResponseを返す", func(t *testing.T) {

		u, err := primitive.NewUUID(id.GenerateUUID())
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		n, err := domain.NewSkillMarketName("minne")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		skillMarketURL, err := primitive.NewURL("https://minne.com")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		res := ToSkillMarketResponse(&domain.SkillMarket{UUID: u, Name: n, URL: skillMarketURL, Icon: "icon"})
		if res.UUID != u.String() || res.Name != n.String() || res.URL != skillMarketURL.String() || res.Icon != "icon" {
			t.Fatalf("unexpected response: %+v", res)
		}
	})

}

func TestToSkillMarketResponses(t *testing.T) {
	t.Run("SkillMarket配列を渡したならSkillMarketResponse配列を返す", func(t *testing.T) {

		u1, err := primitive.NewUUID(id.GenerateUUID())
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		u2, err := primitive.NewUUID(id.GenerateUUID())
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		n1, err := domain.NewSkillMarketName("a")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		n2, err := domain.NewSkillMarketName("b")
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

		res := ToSkillMarketResponses([]*domain.SkillMarket{
			{UUID: u1, Name: n1, URL: url1, Icon: "icon1"},
			{UUID: u2, Name: n2, URL: url2, Icon: "icon2"},
		})
		if len(res) != 2 {
			t.Fatalf("expected 2, got %d", len(res))
		}
		if res[0].Name != "a" || res[1].Name != "b" {
			t.Fatalf("unexpected response: %+v", res)
		}
	})

}
