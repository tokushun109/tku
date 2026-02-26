package user

import (
	domainVO "github.com/tokushun109/tku/clean-backend/internal/domain/vo"
	"strings"
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
