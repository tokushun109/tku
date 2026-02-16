package primitive

import "regexp"

var uuidPattern = regexp.MustCompile(`(?i)^[0-9a-f]{8}-[0-9a-f]{4}-[1-5][0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$`)

type UUID string

func NewUUID(s string) (UUID, error) {
	if !uuidPattern.MatchString(s) {
		return "", ErrInvalidUUID
	}
	return UUID(s), nil
}

func (u UUID) String() string {
	return string(u)
}
