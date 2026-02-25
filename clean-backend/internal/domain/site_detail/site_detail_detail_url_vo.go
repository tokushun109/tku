package site_detail

import "github.com/tokushun109/tku/clean-backend/internal/domain/primitive"

type SiteDetailDetailURL string

func NewSiteDetailDetailURL(v string) (SiteDetailDetailURL, error) {
	url, err := primitive.NewURL(v)
	if err != nil {
		return "", ErrInvalidDetailURL
	}
	return SiteDetailDetailURL(url.String()), nil
}

func (u SiteDetailDetailURL) String() string {
	return string(u)
}
