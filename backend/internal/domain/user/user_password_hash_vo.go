package user

import (
	"strings"

	domainVO "github.com/tokushun109/tku/backend/internal/domain/vo"
)

type UserPasswordHash string

var _ domainVO.ValueObject[string] = UserPasswordHash("")

func NewUserPasswordHash(v string) (UserPasswordHash, error) {
	if strings.TrimSpace(v) == "" {
		return "", ErrInvalidPasswordHash
	}
	return UserPasswordHash(v), nil
}

func (h UserPasswordHash) Value() string {
	return string(h)
}

func (h UserPasswordHash) String() string {
	return h.Value()
}
