package contact

import "errors"

var (
	ErrInvalidName    = errors.New("invalid contact name")
	ErrInvalidCompany = errors.New("invalid contact company")
	ErrInvalidContent = errors.New("invalid contact content")
)
