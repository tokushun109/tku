package creator

import "github.com/tokushun109/tku/clean-backend/internal/shared/optional"

type Creator struct {
	ID           uint
	Name         CreatorName
	Introduction *CreatorIntroduction
	LogoMimeType *CreatorLogoMimeType
	LogoPath     *CreatorLogoPath
}

func New(name string, introduction string) (*Creator, error) {
	creatorName, err := NewCreatorName(name)
	if err != nil {
		return nil, err
	}
	creatorIntroduction, err := optional.ParseOptionalString(introduction, NewCreatorIntroduction)
	if err != nil {
		return nil, err
	}

	return &Creator{
		Name:         creatorName,
		Introduction: creatorIntroduction,
	}, nil
}
