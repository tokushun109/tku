package product

import "errors"

var (
	ErrInvalidID           = errors.New("invalid product id")
	ErrInvalidName         = errors.New("invalid product name")
	ErrInvalidPrice        = errors.New("invalid product price")
	ErrInvalidCategoryID   = errors.New("invalid product category id")
	ErrInvalidTargetID     = errors.New("invalid product target id")
	ErrInvalidCategoryUUID = errors.New("invalid product category uuid")
	ErrInvalidTargetUUID   = errors.New("invalid product target uuid")

	ErrInvalidImageID            = errors.New("invalid product image id")
	ErrInvalidImageName          = errors.New("invalid product image name")
	ErrInvalidImageMimeType      = errors.New("invalid product image mime type")
	ErrInvalidImagePath          = errors.New("invalid product image path")
	ErrInvalidImageDisplayOrder  = errors.New("invalid product image display order")
	ErrInvalidImageProductID     = errors.New("invalid product image product id")
	ErrInvalidImageProductUUID   = errors.New("invalid product image product uuid")
	ErrInvalidImageCategoryUUID  = errors.New("invalid product image category uuid")
	ErrInvalidImageReferenceUUID = errors.New("invalid product image reference uuid")
)
