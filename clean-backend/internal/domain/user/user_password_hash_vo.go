package user

import "strings"

type UserPasswordHash string

func NewUserPasswordHash(v string) (UserPasswordHash, error) {
	if strings.TrimSpace(v) == "" {
		return "", ErrInvalidPasswordHash
	}
	return UserPasswordHash(v), nil
}

func (h UserPasswordHash) String() string {
	return string(h)
}
