package contact

import (
	"strings"
	"unicode/utf8"
)

const companyMaxLen = 20

type ContactCompany string

func NewContactCompany(v string) (ContactCompany, error) {
	trimmed := strings.TrimSpace(v)
	if trimmed == "" {
		return "", ErrInvalidCompany
	}
	if utf8.RuneCountInString(trimmed) > companyMaxLen {
		return "", ErrInvalidCompany
	}
	return ContactCompany(trimmed), nil
}

func (c ContactCompany) String() string {
	return string(c)
}
