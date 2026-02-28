package skill_market

import (
	"context"

	"github.com/tokushun109/tku/clean-backend/internal/domain/primitive"
)

type Repository interface {
	Create(ctx context.Context, s *SkillMarket) (*SkillMarket, error)
	FindAll(ctx context.Context) ([]*SkillMarket, error)
	FindByUUID(ctx context.Context, uuid primitive.UUID) (*SkillMarket, error)
	Update(ctx context.Context, s *SkillMarket) (bool, error)
	Delete(ctx context.Context, uuid primitive.UUID) (bool, error)
}
