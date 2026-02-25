package tag

import "github.com/tokushun109/tku/clean-backend/internal/domain/primitive"

type Tag struct {
	id   uint
	uuid primitive.UUID
	name TagName
}

func New(rawUUID string, name string) (*Tag, error) {
	tag, err := newWithValidatedValues(rawUUID, name)
	if err != nil {
		return nil, err
	}
	return tag, nil
}

func Rebuild(id uint, rawUUID string, name string) (*Tag, error) {
	if id == 0 {
		return nil, ErrInvalidID
	}
	tag, err := newWithValidatedValues(rawUUID, name)
	if err != nil {
		return nil, err
	}
	tag.id = id
	return tag, nil
}

func newWithValidatedValues(rawUUID string, name string) (*Tag, error) {
	uuid, err := primitive.NewUUID(rawUUID)
	if err != nil {
		return nil, err
	}
	tagName, err := NewTagName(name)
	if err != nil {
		return nil, err
	}
	return &Tag{
		uuid: uuid,
		name: tagName,
	}, nil
}

func (t *Tag) ID() uint {
	return t.id
}

func (t *Tag) UUID() primitive.UUID {
	return t.uuid
}

func (t *Tag) Name() TagName {
	return t.name
}

func (t *Tag) ChangeName(name string) error {
	tagName, err := NewTagName(name)
	if err != nil {
		return err
	}
	t.name = tagName
	return nil
}
