package sales_site

import "github.com/tokushun109/tku/clean-backend/internal/domain/primitive"

type SalesSite struct {
	id   primitive.ID
	uuid primitive.UUID
	name SalesSiteName
	url  primitive.URL
	icon string
}

func New(rawUUID string, name string, rawURL string, icon string) (*SalesSite, error) {
	salesSite, err := newWithValidatedValues(rawUUID, name, rawURL, icon)
	if err != nil {
		return nil, err
	}
	return salesSite, nil
}

func Rebuild(id uint, rawUUID string, name string, rawURL string, icon string) (*SalesSite, error) {
	parsedID, err := primitive.NewID(id)
	if err != nil {
		return nil, ErrInvalidID
	}
	salesSite, err := newWithValidatedValues(rawUUID, name, rawURL, icon)
	if err != nil {
		return nil, err
	}
	salesSite.id = parsedID
	return salesSite, nil
}

func newWithValidatedValues(rawUUID string, name string, rawURL string, icon string) (*SalesSite, error) {
	uuid, err := primitive.NewUUID(rawUUID)
	if err != nil {
		return nil, err
	}
	salesSiteName, err := NewSalesSiteName(name)
	if err != nil {
		return nil, err
	}
	salesSiteURL, err := primitive.NewURL(rawURL)
	if err != nil {
		return nil, err
	}
	return &SalesSite{
		uuid: uuid,
		name: salesSiteName,
		url:  salesSiteURL,
		icon: icon,
	}, nil
}

func (s *SalesSite) ID() primitive.ID {
	return s.id
}

func (s *SalesSite) UUID() primitive.UUID {
	return s.uuid
}

func (s *SalesSite) Name() SalesSiteName {
	return s.name
}

func (s *SalesSite) URL() primitive.URL {
	return s.url
}

func (s *SalesSite) Icon() string {
	return s.icon
}

func (s *SalesSite) Change(name string, rawURL string, icon string) error {
	salesSiteName, err := NewSalesSiteName(name)
	if err != nil {
		return err
	}
	salesSiteURL, err := primitive.NewURL(rawURL)
	if err != nil {
		return err
	}
	s.name = salesSiteName
	s.url = salesSiteURL
	s.icon = icon
	return nil
}
