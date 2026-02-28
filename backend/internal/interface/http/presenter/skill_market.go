package presenter

import (
	domain "github.com/tokushun109/tku/backend/internal/domain/skill_market"
	"github.com/tokushun109/tku/backend/internal/interface/http/response"
)

func ToSkillMarketResponse(s *domain.SkillMarket) *response.SkillMarketResponse {
	return &response.SkillMarketResponse{
		UUID: s.UUID().Value(),
		Name: s.Name().Value(),
		URL:  s.URL().Value(),
		Icon: s.Icon(),
	}
}

func ToSkillMarketResponses(list []*domain.SkillMarket) []*response.SkillMarketResponse {
	res := make([]*response.SkillMarketResponse, 0, len(list))
	for _, s := range list {
		res = append(res, ToSkillMarketResponse(s))
	}
	return res
}
