package primitive

import (
	"net/url"
	"strings"
)

type URL string

func NewURL(s string) (URL, error) {
	normalized := strings.ToLower(s)
	if len(normalized) == 0 {
		return "", ErrInvalidURL
	}

	u, err := url.Parse(normalized)
	if err != nil || u.Scheme == "" {
		return "", ErrInvalidURL
	}

	isFileScheme := u.Scheme == "file"
	if (isFileScheme && (len(u.Path) == 0 || u.Path == "/")) ||
		(!isFileScheme && len(u.Host) == 0 && len(u.Fragment) == 0 && len(u.Opaque) == 0) {
		return "", ErrInvalidURL
	}

	return URL(s), nil
}

func (u URL) String() string {
	return string(u)
}
