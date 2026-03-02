package sales_site

import "errors"

var (
	ErrInvalidID      = errors.New("invalid sales site id")
	ErrInvalidName    = errors.New("invalid sales site name")
	ErrNameDuplicated = errors.New("sales site name is duplicate")
)
