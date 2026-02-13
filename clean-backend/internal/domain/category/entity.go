package category

import "strings"

const (
	NameMinLen = 1
	NameMaxLen = 20
)

type Category struct {
	ID   string
	Name string
}

func New(name string) (*Category, error) {
	trimmed := strings.TrimSpace(name)
	if len(trimmed) < NameMinLen || len(trimmed) > NameMaxLen {
		return nil, ErrInvalidName
	}
	return &Category{Name: trimmed}, nil
}
