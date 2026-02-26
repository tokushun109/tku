package product

import (
	domainVO "github.com/tokushun109/tku/clean-backend/internal/domain/vo"
	"strings"
	"unicode/utf8"
)

const (
	productNameMinLen = 1
	productNameMaxLen = 255
)

type ProductName string

var _ domainVO.ValueObject[string] = ProductName("")

func NewProductName(v string) (ProductName, error) {
	trimmed := strings.TrimSpace(v)
	length := utf8.RuneCountInString(trimmed)
	if length < productNameMinLen || length > productNameMaxLen {
		return "", ErrInvalidName
	}
	return ProductName(trimmed), nil
}

func (n ProductName) Value() string {
	return string(n)
}

func (n ProductName) String() string {
	return n.Value()
}
