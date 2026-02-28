package skill_market

import (
	"strings"
	"unicode/utf8"

	domainVO "github.com/tokushun109/tku/backend/internal/domain/vo"
)

const (
	nameMinLen = 1
	nameMaxLen = 30
)

type SkillMarketName string

var _ domainVO.ValueObject[string] = SkillMarketName("")

func NewSkillMarketName(v string) (SkillMarketName, error) {
	trimmed := strings.TrimSpace(v)
	length := utf8.RuneCountInString(trimmed)
	if length < nameMinLen || length > nameMaxLen {
		return "", ErrInvalidName
	}
	return SkillMarketName(trimmed), nil
}

func (n SkillMarketName) Value() string {
	return string(n)
}

func (n SkillMarketName) String() string {
	return n.Value()
}
