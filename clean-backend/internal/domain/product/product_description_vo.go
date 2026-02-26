package product

import (
	"strings"

	domainVO "github.com/tokushun109/tku/clean-backend/internal/domain/vo"
)

type ProductDescription string

var _ domainVO.ValueObject[string] = ProductDescription("")

func NewProductDescription(v string) (ProductDescription, error) {
	return ProductDescription(strings.TrimSpace(v)), nil
}

func (d ProductDescription) Value() string {
	return string(d)
}

func (d ProductDescription) String() string {
	return d.Value()
}
