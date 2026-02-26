package presenter

import (
	domain "github.com/tokushun109/tku/clean-backend/internal/domain/sns"
	"github.com/tokushun109/tku/clean-backend/internal/interface/http/response"
)

func ToSnsResponse(s *domain.Sns) *response.SnsResponse {
	return &response.SnsResponse{
		UUID: s.UUID().Value(),
		Name: s.Name().Value(),
		URL:  s.URL().Value(),
		Icon: s.Icon(),
	}
}

func ToSnsResponses(list []*domain.Sns) []*response.SnsResponse {
	res := make([]*response.SnsResponse, 0, len(list))
	for _, s := range list {
		res = append(res, ToSnsResponse(s))
	}
	return res
}
