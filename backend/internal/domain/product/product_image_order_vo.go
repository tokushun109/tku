package product

import (
	"strconv"

	domainVO "github.com/tokushun109/tku/backend/internal/domain/vo"
)

const (
	productImageDisplayOrderMin = 0
	productImageDisplayOrderMax = 100000
)

// NOTE: 値が大きいほど、表示優先度が高い
type ProductImageDisplayOrder int

var _ domainVO.ValueObject[int] = ProductImageDisplayOrder(0)

func NewProductImageDisplayOrder(v int) (ProductImageDisplayOrder, error) {
	if v < productImageDisplayOrderMin || v > productImageDisplayOrderMax {
		return 0, ErrInvalidImageDisplayOrder
	}
	return ProductImageDisplayOrder(v), nil
}

func (o ProductImageDisplayOrder) Value() int {
	return int(o)
}

func (o ProductImageDisplayOrder) String() string {
	return strconv.Itoa(o.Value())
}
