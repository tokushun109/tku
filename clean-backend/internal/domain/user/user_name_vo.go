package user

import (
	"strings"
	"unicode/utf8"
)

const (
	nameMinLen = 1
	nameMaxLen = 20
)

type UserName string

func NewUserName(v string) (UserName, error) {
	trimmed := strings.TrimSpace(v)
	length := utf8.RuneCountInString(trimmed)
	if length < nameMinLen || length > nameMaxLen {
		return "", ErrInvalidName
	}
	return UserName(trimmed), nil
}

func (n UserName) String() string {
	return string(n)
}
