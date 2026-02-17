package sns

import (
	"strings"
	"unicode/utf8"
)

const (
	nameMinLen = 1
	nameMaxLen = 30
)

type SnsName string

func NewSnsName(v string) (SnsName, error) {
	trimmed := strings.TrimSpace(v)
	length := utf8.RuneCountInString(trimmed)
	if length < nameMinLen || length > nameMaxLen {
		return "", ErrInvalidName
	}
	return SnsName(trimmed), nil
}

func (n SnsName) String() string {
	return string(n)
}
