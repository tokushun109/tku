package creator

import (
	"strings"
	"unicode/utf8"

	domainVO "github.com/tokushun109/tku/backend/internal/domain/vo"
)

const (
	creatorNameMinLen = 1
	creatorNameMaxLen = 30
)

type CreatorName string

var _ domainVO.ValueObject[string] = CreatorName("")

func NewCreatorName(v string) (CreatorName, error) {
	trimmed := strings.TrimSpace(v)
	length := utf8.RuneCountInString(trimmed)
	if length < creatorNameMinLen || length > creatorNameMaxLen {
		return "", ErrInvalidName
	}
	return CreatorName(trimmed), nil
}

func (n CreatorName) Value() string {
	return string(n)
}

func (n CreatorName) String() string {
	return n.Value()
}
