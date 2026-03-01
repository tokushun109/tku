package site_detail

import "github.com/tokushun109/tku/backend/internal/domain/primitive"

type SiteDetail struct {
	id            primitive.ID
	uuid          primitive.UUID
	detailURL     SiteDetailDetailURL
	productUUID   primitive.UUID
	salesSiteUUID primitive.UUID
}

func New(rawUUID string, detailURL string, productUUID string, salesSiteUUID string) (*SiteDetail, error) {
	siteDetail, err := newWithValidatedValues(rawUUID, detailURL, productUUID, salesSiteUUID)
	if err != nil {
		return nil, err
	}
	return siteDetail, nil
}

func Rebuild(id uint, rawUUID string, detailURL string, productUUID string, salesSiteUUID string) (*SiteDetail, error) {
	parsedID, err := primitive.NewID(id)
	if err != nil {
		return nil, ErrInvalidID
	}
	siteDetail, err := newWithValidatedValues(rawUUID, detailURL, productUUID, salesSiteUUID)
	if err != nil {
		return nil, err
	}
	siteDetail.id = parsedID
	return siteDetail, nil
}

func newWithValidatedValues(rawUUID string, detailURL string, productUUID string, salesSiteUUID string) (*SiteDetail, error) {
	uuid, err := primitive.NewUUID(rawUUID)
	if err != nil {
		return nil, err
	}
	parsedURL, err := NewSiteDetailDetailURL(detailURL)
	if err != nil {
		return nil, err
	}
	parsedProductUUID, err := primitive.NewUUID(productUUID)
	if err != nil {
		return nil, ErrInvalidProductUUID
	}
	parsedSalesSiteUUID, err := primitive.NewUUID(salesSiteUUID)
	if err != nil {
		return nil, ErrInvalidSalesSiteUUID
	}

	return &SiteDetail{
		uuid:          uuid,
		detailURL:     parsedURL,
		productUUID:   parsedProductUUID,
		salesSiteUUID: parsedSalesSiteUUID,
	}, nil
}

func (s *SiteDetail) ID() primitive.ID {
	return s.id
}

func (s *SiteDetail) UUID() primitive.UUID {
	return s.uuid
}

func (s *SiteDetail) DetailURL() SiteDetailDetailURL {
	return s.detailURL
}

func (s *SiteDetail) ProductUUID() primitive.UUID {
	return s.productUUID
}

func (s *SiteDetail) SalesSiteUUID() primitive.UUID {
	return s.salesSiteUUID
}
