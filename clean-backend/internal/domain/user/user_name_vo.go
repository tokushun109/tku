package user

import (
	domainVO "github.com/tokushun109/tku/clean-backend/internal/domain/vo"
	"strings"
	"unicode/utf8"
)

const (
	nameMinLen = 1
	nameMaxLen = 20
)

type UserName string

var _ domainVO.ValueObject[string] = UserName("")

func NewUserName(v string) (UserName, error) {
	trimmed := strings.TrimSpace(v)
	length := utf8.RuneCountInString(trimmed)
	if length < nameMinLen || length > nameMaxLen {
		return "", ErrInvalidName
	}
	return UserName(trimmed), nil
}

func (n UserName) Value() string {
	return string(n)
}

func (n UserName) String() string {
	return n.Value()
}
