package sns

import "github.com/tokushun109/tku/clean-backend/internal/domain/primitive"

type Sns struct {
	id   uint
	uuid primitive.UUID
	name SnsName
	url  primitive.URL
	icon string
}

func New(rawUUID string, name string, rawURL string, icon string) (*Sns, error) {
	sns, err := newWithValidatedValues(rawUUID, name, rawURL, icon)
	if err != nil {
		return nil, err
	}
	return sns, nil
}

func Rebuild(id uint, rawUUID string, name string, rawURL string, icon string) (*Sns, error) {
	if id == 0 {
		return nil, ErrInvalidID
	}
	sns, err := newWithValidatedValues(rawUUID, name, rawURL, icon)
	if err != nil {
		return nil, err
	}
	sns.id = id
	return sns, nil
}

func newWithValidatedValues(rawUUID string, name string, rawURL string, icon string) (*Sns, error) {
	uuid, err := primitive.NewUUID(rawUUID)
	if err != nil {
		return nil, err
	}
	snsName, err := NewSnsName(name)
	if err != nil {
		return nil, err
	}
	snsURL, err := primitive.NewURL(rawURL)
	if err != nil {
		return nil, err
	}
	return &Sns{
		uuid: uuid,
		name: snsName,
		url:  snsURL,
		icon: icon,
	}, nil
}

func (s *Sns) ID() uint {
	return s.id
}

func (s *Sns) UUID() primitive.UUID {
	return s.uuid
}

func (s *Sns) Name() SnsName {
	return s.name
}

func (s *Sns) URL() primitive.URL {
	return s.url
}

func (s *Sns) Icon() string {
	return s.icon
}

func (s *Sns) Change(name string, rawURL string, icon string) error {
	snsName, err := NewSnsName(name)
	if err != nil {
		return err
	}
	snsURL, err := primitive.NewURL(rawURL)
	if err != nil {
		return err
	}
	s.name = snsName
	s.url = snsURL
	s.icon = icon
	return nil
}
