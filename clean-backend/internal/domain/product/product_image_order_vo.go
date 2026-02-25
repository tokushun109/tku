package product

const (
	productImageOrderMin = 0
	productImageOrderMax = 100000
)

// NOTE: 値が大きいほど、表示優先度が高い
type ProductImageOrder int

func NewProductImageOrder(v int) (ProductImageOrder, error) {
	if v < productImageOrderMin || v > productImageOrderMax {
		return 0, ErrInvalidImageOrder
	}
	return ProductImageOrder(v), nil
}

func (o ProductImageOrder) Int() int {
	return int(o)
}
