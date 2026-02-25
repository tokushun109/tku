package tag

import "errors"

var (
	ErrInvalidID      = errors.New("invalid tag id")
	ErrInvalidName    = errors.New("invalid tag name")
	ErrNameDuplicated = errors.New("tag name is duplicate")
)
