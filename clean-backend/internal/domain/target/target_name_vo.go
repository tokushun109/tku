package target

import (
	"strings"
	"unicode/utf8"
)

const (
	nameMinLen = 1
	nameMaxLen = 30
)

type TargetName string

func NewTargetName(v string) (TargetName, error) {
	trimmed := strings.TrimSpace(v)
	length := utf8.RuneCountInString(trimmed)
	if length < nameMinLen || length > nameMaxLen {
		return "", ErrInvalidName
	}
	return TargetName(trimmed), nil
}

func (n TargetName) String() string {
	return string(n)
}
