package contact

import (
	"strings"
	"unicode/utf8"
)

const (
	nameMinLen = 1
	nameMaxLen = 30
)

type ContactName string

func NewContactName(v string) (ContactName, error) {
	trimmed := strings.TrimSpace(v)
	length := utf8.RuneCountInString(trimmed)
	if length < nameMinLen || length > nameMaxLen {
		return "", ErrInvalidName
	}
	return ContactName(trimmed), nil
}

func (n ContactName) String() string {
	return string(n)
}
