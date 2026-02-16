package skill_market

import (
	"strings"
	"unicode/utf8"
)

const (
	nameMinLen = 1
	nameMaxLen = 30
)

type SkillMarketName string

func NewSkillMarketName(v string) (SkillMarketName, error) {
	trimmed := strings.TrimSpace(v)
	length := utf8.RuneCountInString(trimmed)
	if length < nameMinLen || length > nameMaxLen {
		return "", ErrInvalidName
	}
	return SkillMarketName(trimmed), nil
}

func (n SkillMarketName) String() string {
	return string(n)
}
