package contact

import (
	"strings"
	"unicode/utf8"

	domainVO "github.com/tokushun109/tku/backend/internal/domain/vo"
)

const (
	nameMinLen = 1
	nameMaxLen = 30
)

type ContactName string

var _ domainVO.ValueObject[string] = ContactName("")

func NewContactName(v string) (ContactName, error) {
	trimmed := strings.TrimSpace(v)
	length := utf8.RuneCountInString(trimmed)
	if length < nameMinLen || length > nameMaxLen {
		return "", ErrInvalidName
	}
	return ContactName(trimmed), nil
}

func (n ContactName) Value() string {
	return string(n)
}

func (n ContactName) String() string {
	return n.Value()
}
