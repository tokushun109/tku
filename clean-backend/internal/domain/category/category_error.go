package category

import "errors"

var (
	ErrInvalidName    = errors.New("invalid category name")
	ErrNameDuplicated = errors.New("category name is duplicate")
)
