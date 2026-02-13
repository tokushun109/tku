package category

type Category struct {
	UUID CategoryUUID
	Name CategoryName
}

func New(name string) (*Category, error) {
	categoryName, err := NewCategoryName(name)
	if err != nil {
		return nil, err
	}
	return &Category{Name: categoryName}, nil
}
