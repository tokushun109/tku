package creator

type Creator struct {
	ID           uint
	Name         CreatorName
	Introduction CreatorIntroduction
	LogoMimeType *CreatorLogoMimeType
	LogoPath     *CreatorLogoPath
}

func New(name string, introduction string) (*Creator, error) {
	creatorName, err := NewCreatorName(name)
	if err != nil {
		return nil, err
	}
	creatorIntroduction, err := NewCreatorIntroduction(introduction)
	if err != nil {
		return nil, err
	}

	return &Creator{
		Name:         creatorName,
		Introduction: creatorIntroduction,
	}, nil
}
