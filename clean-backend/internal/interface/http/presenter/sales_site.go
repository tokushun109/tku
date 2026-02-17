package presenter

import (
	domain "github.com/tokushun109/tku/clean-backend/internal/domain/sales_site"
	"github.com/tokushun109/tku/clean-backend/internal/interface/http/response"
)

func ToSalesSiteResponse(s *domain.SalesSite) *response.SalesSiteResponse {
	return &response.SalesSiteResponse{
		UUID: s.UUID.String(),
		Name: s.Name.String(),
		URL:  s.URL.String(),
		Icon: s.Icon,
	}
}

func ToSalesSiteResponses(list []*domain.SalesSite) []*response.SalesSiteResponse {
	res := make([]*response.SalesSiteResponse, 0, len(list))
	for _, s := range list {
		res = append(res, ToSalesSiteResponse(s))
	}
	return res
}
