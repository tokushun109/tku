package category

import "errors"

var (
	ErrInvalidID      = errors.New("invalid category id")
	ErrInvalidName    = errors.New("invalid category name")
	ErrNameDuplicated = errors.New("category name is duplicate")
)
