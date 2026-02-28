package creator

import (
	"context"

	"github.com/tokushun109/tku/backend/internal/domain/primitive"
)

type Repository interface {
	Find(ctx context.Context) (*Creator, error)
	UpdateProfile(ctx context.Context, c *Creator) (bool, error)
	UpdateLogo(ctx context.Context, creatorID primitive.ID, mimeType CreatorLogoMimeType, logoPath CreatorLogoPath) (bool, error)
}
