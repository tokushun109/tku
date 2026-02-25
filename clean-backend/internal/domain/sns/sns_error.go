package sns

import "errors"

var (
	ErrInvalidID   = errors.New("invalid sns id")
	ErrInvalidName = errors.New("invalid sns name")
)
