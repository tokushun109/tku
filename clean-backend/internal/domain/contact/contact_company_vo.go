package contact

import (
	"strings"
	"unicode/utf8"

	domainVO "github.com/tokushun109/tku/clean-backend/internal/domain/vo"
)

const companyMaxLen = 20

type ContactCompany string

var _ domainVO.ValueObject[string] = ContactCompany("")

func NewContactCompany(v string) (ContactCompany, error) {
	trimmed := strings.TrimSpace(v)
	length := utf8.RuneCountInString(trimmed)
	if length == 0 || length > companyMaxLen {
		return "", ErrInvalidCompany
	}
	return ContactCompany(trimmed), nil
}

func (c ContactCompany) Value() string {
	return string(c)
}

func (c ContactCompany) String() string {
	return c.Value()
}
