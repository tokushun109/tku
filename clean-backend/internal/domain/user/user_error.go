package user

import "errors"

var (
	ErrInvalidName         = errors.New("invalid user name")
	ErrInvalidPasswordHash = errors.New("invalid user password hash")
)
