package product

import "github.com/tokushun109/tku/clean-backend/internal/domain/primitive"

type ProductImage struct {
	id        primitive.ID
	uuid      primitive.UUID
	name      ProductImageName
	mimeType  ProductImageMimeType
	path      ProductImagePath
	order     ProductImageOrder
	productID primitive.ID
}

func NewProductImage(
	rawUUID string,
	name string,
	rawMimeType string,
	rawPath string,
	order int,
	productID uint,
) (*ProductImage, error) {
	image, err := newProductImageWithValidatedValues(rawUUID, name, rawMimeType, rawPath, order, productID)
	if err != nil {
		return nil, err
	}
	return image, nil
}

func RebuildProductImage(
	id uint,
	rawUUID string,
	name string,
	rawMimeType string,
	rawPath string,
	order int,
	productID uint,
) (*ProductImage, error) {
	parsedID, err := primitive.NewID(id)
	if err != nil {
		return nil, ErrInvalidImageID
	}
	image, err := newProductImageWithValidatedValues(rawUUID, name, rawMimeType, rawPath, order, productID)
	if err != nil {
		return nil, err
	}
	image.id = parsedID
	return image, nil
}

func newProductImageWithValidatedValues(
	rawUUID string,
	name string,
	rawMimeType string,
	rawPath string,
	order int,
	productID uint,
) (*ProductImage, error) {
	uuid, err := primitive.NewUUID(rawUUID)
	if err != nil {
		return nil, err
	}
	productImageName, err := NewProductImageName(name)
	if err != nil {
		return nil, err
	}
	mimeType, err := NewProductImageMimeType(rawMimeType)
	if err != nil {
		return nil, err
	}
	path, err := NewProductImagePath(rawPath)
	if err != nil {
		return nil, err
	}
	imageOrder, err := NewProductImageOrder(order)
	if err != nil {
		return nil, err
	}
	parsedProductID, err := primitive.NewID(productID)
	if err != nil {
		return nil, ErrInvalidImageProductID
	}

	return &ProductImage{
		uuid:      uuid,
		name:      productImageName,
		mimeType:  mimeType,
		path:      path,
		order:     imageOrder,
		productID: parsedProductID,
	}, nil
}

func (p *ProductImage) ID() primitive.ID {
	return p.id
}

func (p *ProductImage) UUID() primitive.UUID {
	return p.uuid
}

func (p *ProductImage) Name() ProductImageName {
	return p.name
}

func (p *ProductImage) MimeType() ProductImageMimeType {
	return p.mimeType
}

func (p *ProductImage) Path() ProductImagePath {
	return p.path
}

func (p *ProductImage) Order() ProductImageOrder {
	return p.order
}

func (p *ProductImage) ProductID() uint {
	return p.productID.Uint()
}

func (p *ProductImage) ChangeOrder(order int) error {
	imageOrder, err := NewProductImageOrder(order)
	if err != nil {
		return err
	}
	p.order = imageOrder
	return nil
}
