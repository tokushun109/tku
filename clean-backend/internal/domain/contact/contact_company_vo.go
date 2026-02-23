package contact

import (
	"strings"
	"unicode/utf8"
)

const companyMaxLen = 20

type ContactCompany string

func NewContactCompany(v string) (ContactCompany, error) {
	trimmed := strings.TrimSpace(v)
	length := utf8.RuneCountInString(trimmed)
	if length == 0 || length > companyMaxLen {
		return "", ErrInvalidCompany
	}
	return ContactCompany(trimmed), nil
}

func (c ContactCompany) String() string {
	return string(c)
}
