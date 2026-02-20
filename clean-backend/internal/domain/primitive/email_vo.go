package primitive

import (
	"net/mail"
	"strings"
	"unicode/utf8"
)

const (
	emailMinLen = 1
	emailMaxLen = 50
)

type Email string

func NewEmail(v string) (Email, error) {
	trimmed := strings.TrimSpace(v)
	length := utf8.RuneCountInString(trimmed)
	if length < emailMinLen || length > emailMaxLen {
		return "", ErrInvalidEmail
	}

	parsed, err := mail.ParseAddress(trimmed)
	if err != nil || parsed.Address != trimmed {
		return "", ErrInvalidEmail
	}

	return Email(trimmed), nil
}

func (e Email) String() string {
	return string(e)
}
