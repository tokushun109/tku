package product

import (
	"strings"
	"unicode/utf8"
)

const (
	productImageNameMinLen = 1
	productImageNameMaxLen = 255
)

type ProductImageName string

func NewProductImageName(v string) (ProductImageName, error) {
	trimmed := strings.TrimSpace(v)
	length := utf8.RuneCountInString(trimmed)
	if length < productImageNameMinLen || length > productImageNameMaxLen {
		return "", ErrInvalidImageName
	}
	return ProductImageName(trimmed), nil
}

func (n ProductImageName) String() string {
	return string(n)
}
