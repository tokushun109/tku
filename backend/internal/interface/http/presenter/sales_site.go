package presenter

import (
	domain "github.com/tokushun109/tku/backend/internal/domain/sales_site"
	"github.com/tokushun109/tku/backend/internal/interface/http/response"
)

func ToSalesSiteResponse(s *domain.SalesSite) *response.SalesSiteResponse {
	return &response.SalesSiteResponse{
		UUID: s.UUID().Value(),
		Name: s.Name().Value(),
		URL:  s.URL().Value(),
		Icon: s.Icon(),
	}
}

func ToSalesSiteResponses(list []*domain.SalesSite) []*response.SalesSiteResponse {
	res := make([]*response.SalesSiteResponse, 0, len(list))
	for _, s := range list {
		res = append(res, ToSalesSiteResponse(s))
	}
	return res
}
