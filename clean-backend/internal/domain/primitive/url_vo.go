package primitive

import (
	"net/url"
	"strings"
	"unicode/utf8"
)

const (
	urlMinLen = 1
	urlMaxLen = 255
)

type URL string

func NewURL(s string) (URL, error) {
	trimmed := strings.TrimSpace(s)
	length := utf8.RuneCountInString(trimmed)
	if length < urlMinLen || length > urlMaxLen {
		return "", ErrInvalidURL
	}

	u, err := url.Parse(trimmed)
	if err != nil || u.Scheme == "" || u.Host == "" {
		return "", ErrInvalidURL
	}

	return URL(trimmed), nil
}

func (u URL) String() string {
	return string(u)
}
