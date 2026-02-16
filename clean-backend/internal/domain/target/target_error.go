package target

import "errors"

var (
	ErrInvalidName    = errors.New("invalid target name")
	ErrNameDuplicated = errors.New("target name is duplicate")
)
