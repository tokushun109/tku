package product

import (
	"github.com/tokushun109/tku/backend/internal/domain/primitive"
	domainVO "github.com/tokushun109/tku/backend/internal/domain/vo"
)

type Product struct {
	id           primitive.ID
	uuid         primitive.UUID
	name         ProductName
	description  *ProductDescription
	price        ProductPrice
	isActive     bool
	isRecommend  bool
	categoryUUID *primitive.UUID
	targetUUID   *primitive.UUID
}

func New(
	rawUUID string,
	name string,
	description string,
	price int,
	isActive bool,
	isRecommend bool,
	categoryUUID *string,
	targetUUID *string,
) (*Product, error) {
	product, err := newWithValidatedValues(rawUUID, name, description, price, isActive, isRecommend, categoryUUID, targetUUID)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func Rebuild(
	id uint,
	rawUUID string,
	name string,
	description string,
	price int,
	isActive bool,
	isRecommend bool,
	categoryUUID *string,
	targetUUID *string,
) (*Product, error) {
	parsedID, err := primitive.NewID(id)
	if err != nil {
		return nil, ErrInvalidID
	}
	product, err := newWithValidatedValues(rawUUID, name, description, price, isActive, isRecommend, categoryUUID, targetUUID)
	if err != nil {
		return nil, err
	}
	product.id = parsedID
	return product, nil
}

func newWithValidatedValues(
	rawUUID string,
	name string,
	description string,
	price int,
	isActive bool,
	isRecommend bool,
	categoryUUID *string,
	targetUUID *string,
) (*Product, error) {
	uuid, err := primitive.NewUUID(rawUUID)
	if err != nil {
		return nil, err
	}
	productName, err := NewProductName(name)
	if err != nil {
		return nil, err
	}
	productDescription, err := domainVO.ParseOptionalValue(&description, NewProductDescription)
	if err != nil {
		return nil, err
	}
	productPrice, err := NewProductPrice(price)
	if err != nil {
		return nil, err
	}
	parsedCategoryUUID, err := domainVO.ParseOptionalValue(categoryUUID, primitive.NewUUID)
	if err != nil {
		return nil, ErrInvalidCategoryUUID
	}
	parsedTargetUUID, err := domainVO.ParseOptionalValue(targetUUID, primitive.NewUUID)
	if err != nil {
		return nil, ErrInvalidTargetUUID
	}

	return &Product{
		uuid:         uuid,
		name:         productName,
		description:  productDescription,
		price:        productPrice,
		isActive:     isActive,
		isRecommend:  isRecommend,
		categoryUUID: parsedCategoryUUID,
		targetUUID:   parsedTargetUUID,
	}, nil
}

func (p *Product) ID() primitive.ID {
	return p.id
}

func (p *Product) UUID() primitive.UUID {
	return p.uuid
}

func (p *Product) Name() ProductName {
	return p.name
}

func (p *Product) Description() *ProductDescription {
	return p.description
}

func (p *Product) Price() ProductPrice {
	return p.price
}

func (p *Product) IsActive() bool {
	return p.isActive
}

func (p *Product) IsRecommend() bool {
	return p.isRecommend
}

func (p *Product) CategoryUUID() *primitive.UUID {
	return p.categoryUUID
}

func (p *Product) TargetUUID() *primitive.UUID {
	return p.targetUUID
}

func (p *Product) ChangeProduct(
	name string,
	description string,
	price int,
	isActive bool,
	isRecommend bool,
	categoryUUID *string,
	targetUUID *string,
) error {
	productName, err := NewProductName(name)
	if err != nil {
		return err
	}
	productDescription, err := domainVO.ParseOptionalValue(&description, NewProductDescription)
	if err != nil {
		return err
	}
	productPrice, err := NewProductPrice(price)
	if err != nil {
		return err
	}
	parsedCategoryUUID, err := domainVO.ParseOptionalValue(categoryUUID, primitive.NewUUID)
	if err != nil {
		return ErrInvalidCategoryUUID
	}
	parsedTargetUUID, err := domainVO.ParseOptionalValue(targetUUID, primitive.NewUUID)
	if err != nil {
		return ErrInvalidTargetUUID
	}

	p.name = productName
	p.description = productDescription
	p.price = productPrice
	p.isActive = isActive
	p.isRecommend = isRecommend
	p.categoryUUID = parsedCategoryUUID
	p.targetUUID = parsedTargetUUID
	return nil
}
