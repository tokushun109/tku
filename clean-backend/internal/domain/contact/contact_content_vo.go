package contact

import (
	"strings"
	"unicode/utf8"
)

const contentMinLen = 1
const contentMaxLen = 1000

type ContactContent string

func NewContactContent(v string) (ContactContent, error) {
	trimmed := strings.TrimSpace(v)
	length := utf8.RuneCountInString(trimmed)
	if length < contentMinLen || length > contentMaxLen {
		return "", ErrInvalidContent
	}
	return ContactContent(trimmed), nil
}

func (c ContactContent) String() string {
	return string(c)
}
