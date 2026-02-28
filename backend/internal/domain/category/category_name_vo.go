package category

import (
	"strings"
	"unicode/utf8"

	domainVO "github.com/tokushun109/tku/backend/internal/domain/vo"
)

const (
	nameMinLen = 1
	nameMaxLen = 30
)

type CategoryName string

var _ domainVO.ValueObject[string] = CategoryName("")

func NewCategoryName(v string) (CategoryName, error) {
	trimmed := strings.TrimSpace(v)
	length := utf8.RuneCountInString(trimmed)
	if length < nameMinLen || length > nameMaxLen {
		return "", ErrInvalidName
	}
	return CategoryName(trimmed), nil
}

func (n CategoryName) Value() string {
	return string(n)
}

func (n CategoryName) String() string {
	return n.Value()
}
