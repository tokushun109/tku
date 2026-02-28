package target

import (
	"strings"
	"unicode/utf8"

	domainVO "github.com/tokushun109/tku/backend/internal/domain/vo"
)

const (
	nameMinLen = 1
	nameMaxLen = 30
)

type TargetName string

var _ domainVO.ValueObject[string] = TargetName("")

func NewTargetName(v string) (TargetName, error) {
	trimmed := strings.TrimSpace(v)
	length := utf8.RuneCountInString(trimmed)
	if length < nameMinLen || length > nameMaxLen {
		return "", ErrInvalidName
	}
	return TargetName(trimmed), nil
}

func (n TargetName) Value() string {
	return string(n)
}

func (n TargetName) String() string {
	return n.Value()
}
