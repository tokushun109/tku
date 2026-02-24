package creator

import "context"

type Repository interface {
	Find(ctx context.Context) (*Creator, error)
	UpdateProfile(ctx context.Context, c *Creator) (bool, error)
	UpdateLogo(ctx context.Context, creatorID uint, mimeType CreatorLogoMimeType, logoPath CreatorLogoPath) (bool, error)
}
