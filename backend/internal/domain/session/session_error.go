package session

import "errors"

var (
	ErrInvalidID        = errors.New("invalid session id")
	ErrInvalidUserID    = errors.New("invalid session user id")
	ErrInvalidUserUUID  = errors.New("invalid session user uuid")
	ErrInvalidCreatedAt = errors.New("invalid session created at")
)
