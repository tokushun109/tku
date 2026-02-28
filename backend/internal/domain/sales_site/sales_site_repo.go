package sales_site

import (
	"context"

	"github.com/tokushun109/tku/backend/internal/domain/primitive"
)

type Repository interface {
	Create(ctx context.Context, s *SalesSite) (*SalesSite, error)
	FindAll(ctx context.Context) ([]*SalesSite, error)
	FindByName(ctx context.Context, name SalesSiteName) (*SalesSite, error)
	FindByUUID(ctx context.Context, uuid primitive.UUID) (*SalesSite, error)
	Update(ctx context.Context, s *SalesSite) (bool, error)
	Delete(ctx context.Context, uuid primitive.UUID) (bool, error)
}
