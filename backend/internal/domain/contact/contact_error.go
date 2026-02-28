package contact

import "errors"

var (
	ErrInvalidID      = errors.New("invalid contact id")
	ErrInvalidName    = errors.New("invalid contact name")
	ErrInvalidCompany = errors.New("invalid contact company")
	ErrInvalidContent = errors.New("invalid contact content")
)
