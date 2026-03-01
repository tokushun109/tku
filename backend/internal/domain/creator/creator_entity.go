package creator

import (
	"github.com/tokushun109/tku/backend/internal/domain/primitive"
	domainVO "github.com/tokushun109/tku/backend/internal/domain/vo"
)

type Creator struct {
	uuid         primitive.UUID
	id           primitive.ID
	name         CreatorName
	introduction *CreatorIntroduction
	logoMimeType *CreatorLogoMimeType
	logoPath     *CreatorLogoPath
}

func New(rawUUID string, name string, introduction string) (*Creator, error) {
	creator, err := newWithValidatedValues(rawUUID, name, introduction, "", "")
	if err != nil {
		return nil, err
	}
	return creator, nil
}

func Rebuild(
	id uint,
	rawUUID string,
	name string,
	introduction string,
	logoMimeType string,
	logoPath string,
) (*Creator, error) {
	parsedID, err := primitive.NewID(id)
	if err != nil {
		return nil, ErrInvalidID
	}
	creator, err := newWithValidatedValues(rawUUID, name, introduction, logoMimeType, logoPath)
	if err != nil {
		return nil, err
	}
	creator.id = parsedID
	return creator, nil
}

func newWithValidatedValues(rawUUID string, name string, introduction string, logoMimeType string, logoPath string) (*Creator, error) {
	creatorUUID, err := primitive.NewUUID(rawUUID)
	if err != nil {
		return nil, ErrInvalidUUID
	}
	creatorName, err := NewCreatorName(name)
	if err != nil {
		return nil, err
	}
	creatorIntroduction, err := domainVO.ParseOptionalValue(&introduction, NewCreatorIntroduction)
	if err != nil {
		return nil, err
	}
	parsedLogoMimeType, err := domainVO.ParseOptionalValue(&logoMimeType, NewCreatorLogoMimeType)
	if err != nil {
		return nil, err
	}
	parsedLogoPath, err := domainVO.ParseOptionalValue(&logoPath, NewCreatorLogoPath)
	if err != nil {
		return nil, err
	}

	return &Creator{
		uuid:         creatorUUID,
		name:         creatorName,
		introduction: creatorIntroduction,
		logoMimeType: parsedLogoMimeType,
		logoPath:     parsedLogoPath,
	}, nil
}

func (c *Creator) UUID() primitive.UUID {
	return c.uuid
}

func (c *Creator) ID() primitive.ID {
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
	creatorIntroduction, err := domainVO.ParseOptionalValue(&introduction, NewCreatorIntroduction)
	if err != nil {
		return err
	}

	c.name = creatorName
	c.introduction = creatorIntroduction
	return nil
}
