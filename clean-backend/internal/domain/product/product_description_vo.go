package product

import "strings"

type ProductDescription string

func NewProductDescription(v string) (ProductDescription, error) {
	return ProductDescription(strings.TrimSpace(v)), nil
}

func (d ProductDescription) String() string {
	return string(d)
}
