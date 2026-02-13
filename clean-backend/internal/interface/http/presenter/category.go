package presenter

import (
	domain "github.com/tokushun109/tku/clean-backend/internal/domain/category"
	"github.com/tokushun109/tku/clean-backend/internal/interface/http/response"
)

func ToCategoryResponse(c *domain.Category) *response.CategoryResponse {
	return &response.CategoryResponse{
		UUID: c.ID,
		Name: c.Name,
	}
}

func ToCategoryResponses(list []*domain.Category) []*response.CategoryResponse {
	res := make([]*response.CategoryResponse, 0, len(list))
	for _, c := range list {
		res = append(res, ToCategoryResponse(c))
	}
	return res
}
