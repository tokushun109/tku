package contact

import (
	"strings"
	"unicode/utf8"

	domainVO "github.com/tokushun109/tku/clean-backend/internal/domain/vo"
)

const contentMinLen = 1
const contentMaxLen = 1000

type ContactContent string

var _ domainVO.ValueObject[string] = ContactContent("")

func NewContactContent(v string) (ContactContent, error) {
	trimmed := strings.TrimSpace(v)
	length := utf8.RuneCountInString(trimmed)
	if length < contentMinLen || length > contentMaxLen {
		return "", ErrInvalidContent
	}
	return ContactContent(trimmed), nil
}

func (c ContactContent) Value() string {
	return string(c)
}

func (c ContactContent) String() string {
	return c.Value()
}
