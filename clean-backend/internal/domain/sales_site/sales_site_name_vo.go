package sales_site

import (
	"strings"
	"unicode/utf8"
)

const (
	nameMinLen = 1
	nameMaxLen = 30
)

type SalesSiteName string

func NewSalesSiteName(v string) (SalesSiteName, error) {
	trimmed := strings.TrimSpace(v)
	length := utf8.RuneCountInString(trimmed)
	if length < nameMinLen || length > nameMaxLen {
		return "", ErrInvalidName
	}
	return SalesSiteName(trimmed), nil
}

func (n SalesSiteName) String() string {
	return string(n)
}
