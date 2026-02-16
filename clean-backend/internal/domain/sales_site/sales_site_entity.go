package sales_site

import "github.com/tokushun109/tku/clean-backend/internal/domain/primitive"

type SalesSite struct {
	UUID primitive.UUID
	Name SalesSiteName
	URL  primitive.URL
	Icon string
}

func New(name string, rawURL string, icon string) (*SalesSite, error) {
	salesSiteName, err := NewSalesSiteName(name)
	if err != nil {
		return nil, err
	}
	salesSiteURL, err := primitive.NewURL(rawURL)
	if err != nil {
		return nil, err
	}
	return &SalesSite{Name: salesSiteName, URL: salesSiteURL, Icon: icon}, nil
}
