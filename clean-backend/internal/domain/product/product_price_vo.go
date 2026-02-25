package product

const (
	productPriceMin = 1
	productPriceMax = 1000000
)

type ProductPrice int

func NewProductPrice(v int) (ProductPrice, error) {
	if v < productPriceMin || v > productPriceMax {
		return 0, ErrInvalidPrice
	}
	return ProductPrice(v), nil
}

func (p ProductPrice) Int() int {
	return int(p)
}
