package target

import "errors"

var (
	ErrInvalidID      = errors.New("invalid target id")
	ErrInvalidName    = errors.New("invalid target name")
	ErrNameDuplicated = errors.New("target name is duplicate")
)
