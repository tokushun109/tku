package presenter

import (
	domain "github.com/tokushun109/tku/clean-backend/internal/domain/target"
	"github.com/tokushun109/tku/clean-backend/internal/interface/http/response"
)

func ToTargetResponse(t *domain.Target) *response.TargetResponse {
	return &response.TargetResponse{
		UUID: t.UUID().String(),
		Name: t.Name().String(),
	}
}

func ToTargetResponses(list []*domain.Target) []*response.TargetResponse {
	res := make([]*response.TargetResponse, 0, len(list))
	for _, t := range list {
		res = append(res, ToTargetResponse(t))
	}
	return res
}
