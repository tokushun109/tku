package creator

import "github.com/tokushun109/tku/clean-backend/internal/shared/optional"

type Creator struct {
	id           uint
	name         CreatorName
	introduction *CreatorIntroduction
	logoMimeType *CreatorLogoMimeType
	logoPath     *CreatorLogoPath
}

func New(name string, introduction string) (*Creator, error) {
	creator, err := newWithValidatedValues(name, introduction, "", "")
	if err != nil {
		return nil, err
	}
	return creator, nil
}

func Rebuild(
	id uint,
	name string,
	introduction string,
	logoMimeType string,
	logoPath string,
) (*Creator, error) {
	if id == 0 {
		return nil, ErrInvalidID
	}
	creator, err := newWithValidatedValues(name, introduction, logoMimeType, logoPath)
	if err != nil {
		return nil, err
	}
	creator.id = id
	return creator, nil
}

func newWithValidatedValues(name string, introduction string, logoMimeType string, logoPath string) (*Creator, error) {
	creatorName, err := NewCreatorName(name)
	if err != nil {
		return nil, err
	}
	creatorIntroduction, err := optional.ParseOptionalString(introduction, NewCreatorIntroduction)
	if err != nil {
		return nil, err
	}
	parsedLogoMimeType, err := optional.ParseOptionalString(logoMimeType, NewCreatorLogoMimeType)
	if err != nil {
		return nil, err
	}
	parsedLogoPath, err := optional.ParseOptionalString(logoPath, NewCreatorLogoPath)
	if err != nil {
		return nil, err
	}

	return &Creator{
		name:         creatorName,
		introduction: creatorIntroduction,
		logoMimeType: parsedLogoMimeType,
		logoPath:     parsedLogoPath,
	}, nil
}

func (c *Creator) ID() uint {
	return c.id
}

func (c *Creator) Name() CreatorName {
	return c.name
}

func (c *Creator) Introduction() *CreatorIntroduction {
	return c.introduction
}

func (c *Creator) LogoMimeType() *CreatorLogoMimeType {
	return c.logoMimeType
}

func (c *Creator) LogoPath() *CreatorLogoPath {
	return c.logoPath
}

func (c *Creator) ChangeProfile(name string, introduction string) error {
	creatorName, err := NewCreatorName(name)
	if err != nil {
		return err
	}
	creatorIntroduction, err := optional.ParseOptionalString(introduction, NewCreatorIntroduction)
	if err != nil {
		return err
	}

	c.name = creatorName
	c.introduction = creatorIntroduction
	return nil
}
