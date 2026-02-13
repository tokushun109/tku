package category

import "strings"

const (
	nameMinLen = 1
	nameMaxLen = 20
)

type CategoryName string

func NewCategoryName(v string) (CategoryName, error) {
	trimmed := strings.TrimSpace(v)
	if len(trimmed) < nameMinLen || len(trimmed) > nameMaxLen {
		return "", ErrInvalidName
	}
	return CategoryName(trimmed), nil
}

func (n CategoryName) String() string {
	return string(n)
}
