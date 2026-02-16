package skill_market

import "github.com/tokushun109/tku/clean-backend/internal/domain/primitive"

type SkillMarket struct {
	UUID primitive.UUID
	Name SkillMarketName
	URL  primitive.URL
	Icon string
}

func New(name string, rawURL string, icon string) (*SkillMarket, error) {
	skillMarketName, err := NewSkillMarketName(name)
	if err != nil {
		return nil, err
	}
	skillMarketURL, err := primitive.NewURL(rawURL)
	if err != nil {
		return nil, err
	}
	return &SkillMarket{Name: skillMarketName, URL: skillMarketURL, Icon: icon}, nil
}
