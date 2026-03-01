package site_detail

import (
	"context"

	"github.com/tokushun109/tku/backend/internal/domain/primitive"
)

type Repository interface {
	ReplaceByProductUUID(ctx context.Context, productUUID primitive.UUID, details []*SiteDetail) error
	DeleteByProductUUID(ctx context.Context, productUUID primitive.UUID) error
}
