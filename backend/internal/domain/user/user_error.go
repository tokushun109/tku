package user

import "errors"

var (
	ErrInvalidID           = errors.New("invalid user id")
	ErrInvalidName         = errors.New("invalid user name")
	ErrInvalidPasswordHash = errors.New("invalid user password hash")
)
