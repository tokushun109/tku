package creator

import (
	"strings"
	"unicode/utf8"
)

const (
	creatorNameMinLen = 1
	creatorNameMaxLen = 30
)

type CreatorName string

func NewCreatorName(v string) (CreatorName, error) {
	trimmed := strings.TrimSpace(v)
	length := utf8.RuneCountInString(trimmed)
	if length < creatorNameMinLen || length > creatorNameMaxLen {
		return "", ErrInvalidName
	}
	return CreatorName(trimmed), nil
}

func (n CreatorName) String() string {
	return string(n)
}
