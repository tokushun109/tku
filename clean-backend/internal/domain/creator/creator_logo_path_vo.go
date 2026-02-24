package creator

import (
	"path"
	"strings"
	"unicode/utf8"
)

const creatorLogoPathMaxLen = 255

type CreatorLogoPath string

func NewCreatorLogoPath(v string) (CreatorLogoPath, error) {
	trimmed := strings.TrimSpace(v)
	cleaned := path.Clean(trimmed)
	if cleaned == "." || cleaned == "/" || strings.HasPrefix(cleaned, "../") || strings.HasPrefix(cleaned, "/") {
		return "", ErrInvalidLogoPath
	}
	if utf8.RuneCountInString(cleaned) > creatorLogoPathMaxLen {
		return "", ErrInvalidLogoPath
	}
	return CreatorLogoPath(cleaned), nil
}

func (p CreatorLogoPath) String() string {
	return string(p)
}
