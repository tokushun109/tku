package product

import (
	"strings"
	"unicode/utf8"

	domainVO "github.com/tokushun109/tku/backend/internal/domain/vo"
)

const (
	productImageNameMinLen = 1
	productImageNameMaxLen = 255
)

type ProductImageName string

var _ domainVO.ValueObject[string] = ProductImageName("")

func NewProductImageName(v string) (ProductImageName, error) {
	trimmed := strings.TrimSpace(v)
	length := utf8.RuneCountInString(trimmed)
	if length < productImageNameMinLen || length > productImageNameMaxLen {
		return "", ErrInvalidImageName
	}
	return ProductImageName(trimmed), nil
}

func (n ProductImageName) Value() string {
	return string(n)
}

func (n ProductImageName) String() string {
	return n.Value()
}
