package product

import (
	"strings"
	"unicode/utf8"
)

const (
	productNameMinLen = 1
	productNameMaxLen = 255
)

type ProductName string

func NewProductName(v string) (ProductName, error) {
	trimmed := strings.TrimSpace(v)
	length := utf8.RuneCountInString(trimmed)
	if length < productNameMinLen || length > productNameMaxLen {
		return "", ErrInvalidName
	}
	return ProductName(trimmed), nil
}

func (n ProductName) String() string {
	return string(n)
}
