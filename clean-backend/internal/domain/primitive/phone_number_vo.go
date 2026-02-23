package primitive

import (
	"strings"
)

const phoneNumberMaxLen = 20

type PhoneNumber string

func NewPhoneNumber(v string) (PhoneNumber, error) {
	trimmed := strings.TrimSpace(v)
	if trimmed == "" || len(trimmed) > phoneNumberMaxLen {
		return "", ErrInvalidPhoneNumber
	}

	for i := 0; i < len(trimmed); i++ {
		if trimmed[i] < '0' || trimmed[i] > '9' {
			return "", ErrInvalidPhoneNumber
		}
	}

	return PhoneNumber(trimmed), nil
}

func (p PhoneNumber) String() string {
	return string(p)
}
