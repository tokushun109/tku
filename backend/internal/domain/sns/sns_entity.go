package sns

import "github.com/tokushun109/tku/backend/internal/domain/primitive"

type Sns struct {
	id   primitive.ID
	uuid primitive.UUID
	name SnsName
	url  primitive.URL
}

func New(rawUUID string, name string, rawURL string) (*Sns, error) {
	sns, err := newWithValidatedValues(rawUUID, name, rawURL)
	if err != nil {
		return nil, err
	}
	return sns, nil
}

func Rebuild(id uint, rawUUID string, name string, rawURL string) (*Sns, error) {
	parsedID, err := primitive.NewID(id)
	if err != nil {
		return nil, ErrInvalidID
	}
	sns, err := newWithValidatedValues(rawUUID, name, rawURL)
	if err != nil {
		return nil, err
	}
	sns.id = parsedID
	return sns, nil
}

func newWithValidatedValues(rawUUID string, name string, rawURL string) (*Sns, error) {
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
	}, nil
}

func (s *Sns) ID() primitive.ID {
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

func (s *Sns) Change(name string, rawURL string) error {
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
	return nil
}
