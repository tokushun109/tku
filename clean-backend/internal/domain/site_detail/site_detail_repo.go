package site_detail

import (
	"context"

	"github.com/tokushun109/tku/clean-backend/internal/domain/primitive"
)

type Repository interface {
	ReplaceByProductID(ctx context.Context, productID primitive.ID, details []*SiteDetail) error
	DeleteByProductID(ctx context.Context, productID primitive.ID) error
}
