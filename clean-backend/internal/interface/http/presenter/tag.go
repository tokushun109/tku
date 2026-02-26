package presenter

import (
	domain "github.com/tokushun109/tku/clean-backend/internal/domain/tag"
	"github.com/tokushun109/tku/clean-backend/internal/interface/http/response"
)

func ToTagResponse(t *domain.Tag) *response.TagResponse {
	return &response.TagResponse{
		UUID: t.UUID().Value(),
		Name: t.Name().Value(),
	}
}

func ToTagResponses(list []*domain.Tag) []*response.TagResponse {
	res := make([]*response.TagResponse, 0, len(list))
	for _, t := range list {
		res = append(res, ToTagResponse(t))
	}
	return res
}
