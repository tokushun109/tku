package skill_market

import "github.com/tokushun109/tku/backend/internal/domain/primitive"

type SkillMarket struct {
	id   primitive.ID
	uuid primitive.UUID
	name SkillMarketName
	url  primitive.URL
	icon string
}

func New(rawUUID string, name string, rawURL string, icon string) (*SkillMarket, error) {
	skillMarket, err := newWithValidatedValues(rawUUID, name, rawURL, icon)
	if err != nil {
		return nil, err
	}
	return skillMarket, nil
}

func Rebuild(id uint, rawUUID string, name string, rawURL string, icon string) (*SkillMarket, error) {
	parsedID, err := primitive.NewID(id)
	if err != nil {
		return nil, ErrInvalidID
	}
	skillMarket, err := newWithValidatedValues(rawUUID, name, rawURL, icon)
	if err != nil {
		return nil, err
	}
	skillMarket.id = parsedID
	return skillMarket, nil
}

func newWithValidatedValues(rawUUID string, name string, rawURL string, icon string) (*SkillMarket, error) {
	uuid, err := primitive.NewUUID(rawUUID)
	if err != nil {
		return nil, err
	}
	skillMarketName, err := NewSkillMarketName(name)
	if err != nil {
		return nil, err
	}
	skillMarketURL, err := primitive.NewURL(rawURL)
	if err != nil {
		return nil, err
	}
	return &SkillMarket{
		uuid: uuid,
		name: skillMarketName,
		url:  skillMarketURL,
		icon: icon,
	}, nil
}

func (s *SkillMarket) ID() primitive.ID {
	return s.id
}

func (s *SkillMarket) UUID() primitive.UUID {
	return s.uuid
}

func (s *SkillMarket) Name() SkillMarketName {
	return s.name
}

func (s *SkillMarket) URL() primitive.URL {
	return s.url
}

func (s *SkillMarket) Icon() string {
	return s.icon
}

func (s *SkillMarket) Change(name string, rawURL string, icon string) error {
	skillMarketName, err := NewSkillMarketName(name)
	if err != nil {
		return err
	}
	skillMarketURL, err := primitive.NewURL(rawURL)
	if err != nil {
		return err
	}
	s.name = skillMarketName
	s.url = skillMarketURL
	s.icon = icon
	return nil
}
