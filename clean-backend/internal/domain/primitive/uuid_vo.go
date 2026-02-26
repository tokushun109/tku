package primitive

import (
	domainVO "github.com/tokushun109/tku/clean-backend/internal/domain/vo"
	"regexp"
)

var uuidPattern = regexp.MustCompile(`(?i)^[0-9a-f]{8}-[0-9a-f]{4}-[1-5][0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$`)

type UUID string

var _ domainVO.ValueObject[string] = UUID("")

func NewUUID(s string) (UUID, error) {
	if !uuidPattern.MatchString(s) {
		return "", ErrInvalidUUID
	}
	return UUID(s), nil
}

func (u UUID) Value() string {
	return string(u)
}

func (u UUID) String() string {
	return u.Value()
}
