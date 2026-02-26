package product

import (
	domainVO "github.com/tokushun109/tku/clean-backend/internal/domain/vo"
	"strconv"
)

const (
	productImageOrderMin = 0
	productImageOrderMax = 100000
)

// NOTE: 値が大きいほど、表示優先度が高い
type ProductImageOrder int

var _ domainVO.ValueObject[int] = ProductImageOrder(0)

func NewProductImageOrder(v int) (ProductImageOrder, error) {
	if v < productImageOrderMin || v > productImageOrderMax {
		return 0, ErrInvalidImageOrder
	}
	return ProductImageOrder(v), nil
}

func (o ProductImageOrder) Value() int {
	return int(o)
}

func (o ProductImageOrder) String() string {
	return strconv.Itoa(o.Value())
}
