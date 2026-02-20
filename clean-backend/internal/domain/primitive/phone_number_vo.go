package primitive

import (
	"strings"
	"unicode/utf8"
)

const phoneNumberMaxLen = 20

type PhoneNumber string

func NewPhoneNumber(v string) (PhoneNumber, error) {
	trimmed := strings.TrimSpace(v)
	if trimmed == "" || utf8.RuneCountInString(trimmed) > phoneNumberMaxLen {
		return "", ErrInvalidPhoneNumber
	}

	return PhoneNumber(trimmed), nil
}

func (p PhoneNumber) String() string {
	return string(p)
}
