package user

import "errors"

var (
	ErrInvalidPasswordHash = errors.New("invalid user password hash")
)
