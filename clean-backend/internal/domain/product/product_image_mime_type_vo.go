package product

import (
	domainVO "github.com/tokushun109/tku/clean-backend/internal/domain/vo"
	"strings"
)

type ProductImageMimeType string

var _ domainVO.ValueObject[string] = ProductImageMimeType("")

var productImageExtensionMap = map[ProductImageMimeType]string{
	"image/jpeg": ".jpg",
	"image/jpg":  ".jpg",
	"image/png":  ".png",
	"image/gif":  ".gif",
	"image/webp": ".webp",
}

func NewProductImageMimeType(v string) (ProductImageMimeType, error) {
	normalized := strings.TrimSpace(strings.ToLower(v))
	mimeType := ProductImageMimeType(normalized)
	if _, ok := productImageExtensionMap[mimeType]; !ok {
		return "", ErrInvalidImageMimeType
	}
	return mimeType, nil
}

func (m ProductImageMimeType) Value() string {
	return string(m)
}

func (m ProductImageMimeType) String() string {
	return m.Value()
}

func (m ProductImageMimeType) Extension() string {
	return productImageExtensionMap[m]
}
