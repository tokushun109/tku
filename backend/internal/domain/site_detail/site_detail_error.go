package site_detail

import "errors"

var (
	ErrInvalidID            = errors.New("invalid site detail id")
	ErrInvalidDetailURL     = errors.New("invalid site detail url")
	ErrInvalidProductID     = errors.New("invalid site detail product id")
	ErrInvalidSalesSiteID   = errors.New("invalid site detail sales site id")
	ErrInvalidProductUUID   = errors.New("invalid site detail product uuid")
	ErrInvalidSalesSiteUUID = errors.New("invalid site detail sales site uuid")
)
