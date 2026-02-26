package tag

import (
	"strings"
	"unicode/utf8"

	domainVO "github.com/tokushun109/tku/clean-backend/internal/domain/vo"
)

const (
	nameMinLen = 1
	nameMaxLen = 30
)

type TagName string

var _ domainVO.ValueObject[string] = TagName("")

func NewTagName(v string) (TagName, error) {
	trimmed := strings.TrimSpace(v)
	length := utf8.RuneCountInString(trimmed)
	if length < nameMinLen || length > nameMaxLen {
		return "", ErrInvalidName
	}
	return TagName(trimmed), nil
}

func (n TagName) Value() string {
	return string(n)
}

func (n TagName) String() string {
	return n.Value()
}
