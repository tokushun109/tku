package site_detail

import "github.com/tokushun109/tku/clean-backend/internal/domain/primitive"

type SiteDetail struct {
	id          primitive.ID
	uuid        primitive.UUID
	detailURL   SiteDetailDetailURL
	productID   primitive.ID
	salesSiteID primitive.ID
}

func New(rawUUID string, detailURL string, productID uint, salesSiteID uint) (*SiteDetail, error) {
	siteDetail, err := newWithValidatedValues(rawUUID, detailURL, productID, salesSiteID)
	if err != nil {
		return nil, err
	}
	return siteDetail, nil
}

func Rebuild(id uint, rawUUID string, detailURL string, productID uint, salesSiteID uint) (*SiteDetail, error) {
	parsedID, err := primitive.NewID(id)
	if err != nil {
		return nil, ErrInvalidID
	}
	siteDetail, err := newWithValidatedValues(rawUUID, detailURL, productID, salesSiteID)
	if err != nil {
		return nil, err
	}
	siteDetail.id = parsedID
	return siteDetail, nil
}

func newWithValidatedValues(rawUUID string, detailURL string, productID uint, salesSiteID uint) (*SiteDetail, error) {
	uuid, err := primitive.NewUUID(rawUUID)
	if err != nil {
		return nil, err
	}
	parsedURL, err := NewSiteDetailDetailURL(detailURL)
	if err != nil {
		return nil, err
	}
	parsedProductID, err := primitive.NewID(productID)
	if err != nil {
		return nil, ErrInvalidProductID
	}
	parsedSalesSiteID, err := primitive.NewID(salesSiteID)
	if err != nil {
		return nil, ErrInvalidSalesSiteID
	}

	return &SiteDetail{
		uuid:        uuid,
		detailURL:   parsedURL,
		productID:   parsedProductID,
		salesSiteID: parsedSalesSiteID,
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

func (s *SiteDetail) ProductID() uint {
	return s.productID.Value()
}

func (s *SiteDetail) SalesSiteID() uint {
	return s.salesSiteID.Value()
}
