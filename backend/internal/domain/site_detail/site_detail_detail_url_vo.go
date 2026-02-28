package site_detail

import (
	"github.com/tokushun109/tku/backend/internal/domain/primitive"
	domainVO "github.com/tokushun109/tku/backend/internal/domain/vo"
)

type SiteDetailDetailURL string

var _ domainVO.ValueObject[string] = SiteDetailDetailURL("")

func NewSiteDetailDetailURL(v string) (SiteDetailDetailURL, error) {
	url, err := primitive.NewURL(v)
	if err != nil {
		return "", ErrInvalidDetailURL
	}
	return SiteDetailDetailURL(url.Value()), nil
}

func (u SiteDetailDetailURL) Value() string {
	return string(u)
}

func (u SiteDetailDetailURL) String() string {
	return u.Value()
}
