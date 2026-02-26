package presenter

import (
	"testing"

	"github.com/tokushun109/tku/clean-backend/internal/domain/primitive"
	domain "github.com/tokushun109/tku/clean-backend/internal/domain/sales_site"
	"github.com/tokushun109/tku/clean-backend/internal/shared/id"
)

func TestToSalesSiteResponse(t *testing.T) {
	t.Run("SalesSiteを渡したならSalesSiteResponseを返す", func(t *testing.T) {

		u, err := primitive.NewUUID(id.GenerateUUID())
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		n, err := domain.NewSalesSiteName("Creema")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		salesSiteURL, err := primitive.NewURL("https://www.creema.jp")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		salesSite, err := domain.Rebuild(1, u.Value(), n.Value(), salesSiteURL.Value(), "icon")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		res := ToSalesSiteResponse(salesSite)
		if res.UUID != u.Value() || res.Name != n.Value() || res.URL != salesSiteURL.Value() || res.Icon != "icon" {
			t.Fatalf("unexpected response: %+v", res)
		}
	})

}

func TestToSalesSiteResponses(t *testing.T) {
	t.Run("SalesSite配列を渡したならSalesSiteResponse配列を返す", func(t *testing.T) {

		u1, err := primitive.NewUUID(id.GenerateUUID())
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		u2, err := primitive.NewUUID(id.GenerateUUID())
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		n1, err := domain.NewSalesSiteName("a")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		n2, err := domain.NewSalesSiteName("b")
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

		salesSite1, err := domain.Rebuild(1, u1.Value(), n1.Value(), url1.Value(), "icon1")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		salesSite2, err := domain.Rebuild(2, u2.Value(), n2.Value(), url2.Value(), "icon2")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		res := ToSalesSiteResponses([]*domain.SalesSite{salesSite1, salesSite2})
		if len(res) != 2 {
			t.Fatalf("expected 2, got %d", len(res))
		}
		if res[0].Name != "a" || res[1].Name != "b" {
			t.Fatalf("unexpected response: %+v", res)
		}
	})

}
