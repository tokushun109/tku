package creator

import (
	"strings"
	"unicode/utf8"
)

const (
	creatorIntroductionMinLen = 1
	creatorIntroductionMaxLen = 1000
)

type CreatorIntroduction string

func NewCreatorIntroduction(v string) (CreatorIntroduction, error) {
	trimmed := strings.TrimSpace(v)
	length := utf8.RuneCountInString(trimmed)
	if length < creatorIntroductionMinLen || length > creatorIntroductionMaxLen {
		return "", ErrInvalidIntroduction
	}
	return CreatorIntroduction(trimmed), nil
}

func NewCreatorIntroductionForRead(v string) (CreatorIntroduction, error) {
	trimmed := strings.TrimSpace(v)
	length := utf8.RuneCountInString(trimmed)
	if length > creatorIntroductionMaxLen {
		return "", ErrInvalidIntroduction
	}
	return CreatorIntroduction(trimmed), nil
}

func (i CreatorIntroduction) String() string {
	return string(i)
}
