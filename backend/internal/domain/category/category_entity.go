package category

import "github.com/tokushun109/tku/backend/internal/domain/primitive"

type Category struct {
	id   primitive.ID
	uuid primitive.UUID
	name CategoryName
}

func New(rawUUID string, name string) (*Category, error) {
	category, err := newWithValidatedValues(rawUUID, name)
	if err != nil {
		return nil, err
	}
	return category, nil
}

func Rebuild(id uint, rawUUID string, name string) (*Category, error) {
	parsedID, err := primitive.NewID(id)
	if err != nil {
		return nil, ErrInvalidID
	}
	category, err := newWithValidatedValues(rawUUID, name)
	if err != nil {
		return nil, err
	}
	category.id = parsedID
	return category, nil
}

func newWithValidatedValues(rawUUID string, name string) (*Category, error) {
	uuid, err := primitive.NewUUID(rawUUID)
	if err != nil {
		return nil, err
	}
	categoryName, err := NewCategoryName(name)
	if err != nil {
		return nil, err
	}
	return &Category{
		uuid: uuid,
		name: categoryName,
	}, nil
}

func (c *Category) ID() primitive.ID {
	return c.id
}

func (c *Category) UUID() primitive.UUID {
	return c.uuid
}

func (c *Category) Name() CategoryName {
	return c.name
}

func (c *Category) ChangeName(name string) error {
	categoryName, err := NewCategoryName(name)
	if err != nil {
		return err
	}
	c.name = categoryName
	return nil
}
