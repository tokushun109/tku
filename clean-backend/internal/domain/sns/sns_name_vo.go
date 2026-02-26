package sns

import (
	"strings"
	"unicode/utf8"

	domainVO "github.com/tokushun109/tku/clean-backend/internal/domain/vo"
)

const (
	nameMinLen = 1
	nameMaxLen = 30
)

type SnsName string

var _ domainVO.ValueObject[string] = SnsName("")

func NewSnsName(v string) (SnsName, error) {
	trimmed := strings.TrimSpace(v)
	length := utf8.RuneCountInString(trimmed)
	if length < nameMinLen || length > nameMaxLen {
		return "", ErrInvalidName
	}
	return SnsName(trimmed), nil
}

func (n SnsName) Value() string {
	return string(n)
}

func (n SnsName) String() string {
	return n.Value()
}
