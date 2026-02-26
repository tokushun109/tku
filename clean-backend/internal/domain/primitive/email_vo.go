package primitive

import (
	"net/mail"
	"strings"
	"unicode/utf8"

	domainVO "github.com/tokushun109/tku/clean-backend/internal/domain/vo"
)

const (
	emailMinLen = 1
	emailMaxLen = 50
)

type Email string

var _ domainVO.ValueObject[string] = Email("")

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

func (e Email) Value() string {
	return string(e)
}

func (e Email) String() string {
	return e.Value()
}
