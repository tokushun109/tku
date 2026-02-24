package creator

import "strings"

var creatorLogoMimeTypeToExtension = map[string]string{
	"image/jpeg": ".jpg",
	"image/png":  ".png",
	"image/gif":  ".gif",
	"image/webp": ".webp",
}

type CreatorLogoMimeType string

func NewCreatorLogoMimeType(v string) (CreatorLogoMimeType, error) {
	normalized := strings.ToLower(strings.TrimSpace(v))
	if _, ok := creatorLogoMimeTypeToExtension[normalized]; !ok {
		return "", ErrInvalidLogoMimeType
	}
	return CreatorLogoMimeType(normalized), nil
}

func (m CreatorLogoMimeType) String() string {
	return string(m)
}

func (m CreatorLogoMimeType) Extension() string {
	return creatorLogoMimeTypeToExtension[m.String()]
}
