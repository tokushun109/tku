package product

import "strings"

const productImagePathMaxLen = 255

type ProductImagePath string

func NewProductImagePath(v string) (ProductImagePath, error) {
	trimmed := strings.TrimSpace(v)
	if len(trimmed) == 0 || len(trimmed) > productImagePathMaxLen {
		return "", ErrInvalidImagePath
	}
	return ProductImagePath(trimmed), nil
}

func (p ProductImagePath) String() string {
	return string(p)
}
