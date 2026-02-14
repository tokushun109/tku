package category

import "github.com/tokushun109/tku/clean-backend/internal/domain/primitive"

type Category struct {
	UUID primitive.UUID
	Name CategoryName
}

func New(name string) (*Category, error) {
	categoryName, err := NewCategoryName(name)
	if err != nil {
		return nil, err
	}
	return &Category{Name: categoryName}, nil
}
