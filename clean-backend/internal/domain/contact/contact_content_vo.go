package contact

import (
	"strings"
	"unicode/utf8"
)

const contentMinLen = 1

type ContactContent string

func NewContactContent(v string) (ContactContent, error) {
	trimmed := strings.TrimSpace(v)
	if utf8.RuneCountInString(trimmed) < contentMinLen {
		return "", ErrInvalidContent
	}
	return ContactContent(trimmed), nil
}

func (c ContactContent) String() string {
	return string(c)
}
