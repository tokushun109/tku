package tag

import "errors"

var (
	ErrInvalidName    = errors.New("invalid tag name")
	ErrNameDuplicated = errors.New("tag name is duplicate")
)
