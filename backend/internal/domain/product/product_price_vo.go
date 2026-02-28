package product

import (
	"strconv"

	domainVO "github.com/tokushun109/tku/backend/internal/domain/vo"
)

const (
	productPriceMin = 1
	productPriceMax = 1000000
)

type ProductPrice int

var _ domainVO.ValueObject[int] = ProductPrice(0)

func NewProductPrice(v int) (ProductPrice, error) {
	if v < productPriceMin || v > productPriceMax {
		return 0, ErrInvalidPrice
	}
	return ProductPrice(v), nil
}

func (p ProductPrice) Value() int {
	return int(p)
}

func (p ProductPrice) String() string {
	return strconv.Itoa(p.Value())
}
