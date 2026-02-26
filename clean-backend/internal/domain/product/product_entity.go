package product

import (
	"github.com/tokushun109/tku/clean-backend/internal/domain/primitive"
	domainVO "github.com/tokushun109/tku/clean-backend/internal/domain/vo"
)

type Product struct {
	id          primitive.ID
	uuid        primitive.UUID
	name        ProductName
	description *ProductDescription
	price       ProductPrice
	isActive    bool
	isRecommend bool
	categoryID  *primitive.ID
	targetID    *primitive.ID
}

func New(
	rawUUID string,
	name string,
	description string,
	price int,
	isActive bool,
	isRecommend bool,
	categoryID *uint,
	targetID *uint,
) (*Product, error) {
	product, err := newWithValidatedValues(rawUUID, name, description, price, isActive, isRecommend, categoryID, targetID)
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
	categoryID *uint,
	targetID *uint,
) (*Product, error) {
	parsedID, err := primitive.NewID(id)
	if err != nil {
		return nil, ErrInvalidID
	}
	product, err := newWithValidatedValues(rawUUID, name, description, price, isActive, isRecommend, categoryID, targetID)
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
	categoryID *uint,
	targetID *uint,
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
	parsedCategoryID, err := domainVO.ParseOptionalValue(categoryID, primitive.NewID)
	if err != nil {
		return nil, ErrInvalidCategoryID
	}
	parsedTargetID, err := domainVO.ParseOptionalValue(targetID, primitive.NewID)
	if err != nil {
		return nil, ErrInvalidTargetID
	}

	return &Product{
		uuid:        uuid,
		name:        productName,
		description: productDescription,
		price:       productPrice,
		isActive:    isActive,
		isRecommend: isRecommend,
		categoryID:  parsedCategoryID,
		targetID:    parsedTargetID,
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

func (p *Product) CategoryID() *primitive.ID {
	return p.categoryID
}

func (p *Product) TargetID() *primitive.ID {
	return p.targetID
}

func (p *Product) ChangeProduct(
	name string,
	description string,
	price int,
	isActive bool,
	isRecommend bool,
	categoryID *uint,
	targetID *uint,
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
	parsedCategoryID, err := domainVO.ParseOptionalValue(categoryID, primitive.NewID)
	if err != nil {
		return ErrInvalidCategoryID
	}
	parsedTargetID, err := domainVO.ParseOptionalValue(targetID, primitive.NewID)
	if err != nil {
		return ErrInvalidTargetID
	}

	p.name = productName
	p.description = productDescription
	p.price = productPrice
	p.isActive = isActive
	p.isRecommend = isRecommend
	p.categoryID = parsedCategoryID
	p.targetID = parsedTargetID
	return nil
}
