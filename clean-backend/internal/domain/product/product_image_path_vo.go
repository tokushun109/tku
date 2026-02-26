package product

import (
	"strings"

	domainVO "github.com/tokushun109/tku/clean-backend/internal/domain/vo"
)

const productImagePathMaxLen = 255

type ProductImagePath string

var _ domainVO.ValueObject[string] = ProductImagePath("")

func NewProductImagePath(v string) (ProductImagePath, error) {
	trimmed := strings.TrimSpace(v)
	if len(trimmed) == 0 || len(trimmed) > productImagePathMaxLen {
		return "", ErrInvalidImagePath
	}
	return ProductImagePath(trimmed), nil
}

func (p ProductImagePath) Value() string {
	return string(p)
}

func (p ProductImagePath) String() string {
	return p.Value()
}
