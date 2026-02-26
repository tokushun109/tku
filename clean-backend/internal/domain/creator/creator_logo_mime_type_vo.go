package creator

import (
	"strings"

	domainVO "github.com/tokushun109/tku/clean-backend/internal/domain/vo"
)

var creatorLogoMimeTypeToExtension = map[string]string{
	"image/jpeg": ".jpg",
	"image/png":  ".png",
	"image/gif":  ".gif",
	"image/webp": ".webp",
}

type CreatorLogoMimeType string

var _ domainVO.ValueObject[string] = CreatorLogoMimeType("")

func NewCreatorLogoMimeType(v string) (CreatorLogoMimeType, error) {
	normalized := strings.ToLower(strings.TrimSpace(v))
	if _, ok := creatorLogoMimeTypeToExtension[normalized]; !ok {
		return "", ErrInvalidLogoMimeType
	}
	return CreatorLogoMimeType(normalized), nil
}

func (m CreatorLogoMimeType) Value() string {
	return string(m)
}

func (m CreatorLogoMimeType) String() string {
	return m.Value()
}

func (m CreatorLogoMimeType) Extension() string {
	return creatorLogoMimeTypeToExtension[m.Value()]
}
