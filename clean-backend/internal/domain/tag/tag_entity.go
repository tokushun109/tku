package tag

import "github.com/tokushun109/tku/clean-backend/internal/domain/primitive"

type Tag struct {
	UUID primitive.UUID
	Name TagName
}

func New(name string) (*Tag, error) {
	tagName, err := NewTagName(name)
	if err != nil {
		return nil, err
	}
	return &Tag{Name: tagName}, nil
}
