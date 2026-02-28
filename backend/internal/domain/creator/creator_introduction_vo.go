package creator

import (
	"strings"
	"unicode/utf8"

	domainVO "github.com/tokushun109/tku/backend/internal/domain/vo"
)

const (
	creatorIntroductionMinLen = 1
	creatorIntroductionMaxLen = 1000
)

type CreatorIntroduction string

var _ domainVO.ValueObject[string] = CreatorIntroduction("")

func NewCreatorIntroduction(v string) (CreatorIntroduction, error) {
	trimmed := strings.TrimSpace(v)
	length := utf8.RuneCountInString(trimmed)
	if length < creatorIntroductionMinLen || length > creatorIntroductionMaxLen {
		return "", ErrInvalidIntroduction
	}
	return CreatorIntroduction(trimmed), nil
}

func (i CreatorIntroduction) Value() string {
	return string(i)
}

func (i CreatorIntroduction) String() string {
	return i.Value()
}
