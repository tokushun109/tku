package category

import (
	"strings"
	"unicode/utf8"
)

const (
	nameMinLen = 1
	nameMaxLen = 30
)

type CategoryName string

func NewCategoryName(v string) (CategoryName, error) {
	trimmed := strings.TrimSpace(v)
	length := utf8.RuneCountInString(trimmed)
	if length < nameMinLen || length > nameMaxLen {
		return "", ErrInvalidName
	}
	return CategoryName(trimmed), nil
}

func (n CategoryName) String() string {
	return string(n)
}
