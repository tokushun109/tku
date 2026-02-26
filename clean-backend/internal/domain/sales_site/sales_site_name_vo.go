package sales_site

import (
	"strings"
	"unicode/utf8"

	domainVO "github.com/tokushun109/tku/clean-backend/internal/domain/vo"
)

const (
	nameMinLen = 1
	nameMaxLen = 30
)

type SalesSiteName string

var _ domainVO.ValueObject[string] = SalesSiteName("")

func NewSalesSiteName(v string) (SalesSiteName, error) {
	trimmed := strings.TrimSpace(v)
	length := utf8.RuneCountInString(trimmed)
	if length < nameMinLen || length > nameMaxLen {
		return "", ErrInvalidName
	}
	return SalesSiteName(trimmed), nil
}

func (n SalesSiteName) Value() string {
	return string(n)
}

func (n SalesSiteName) String() string {
	return n.Value()
}
