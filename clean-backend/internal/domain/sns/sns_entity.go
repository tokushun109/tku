package sns

import "github.com/tokushun109/tku/clean-backend/internal/domain/primitive"

type Sns struct {
	UUID primitive.UUID
	Name SnsName
	URL  primitive.URL
	Icon string
}

func New(name string, rawURL string, icon string) (*Sns, error) {
	snsName, err := NewSnsName(name)
	if err != nil {
		return nil, err
	}
	snsURL, err := primitive.NewURL(rawURL)
	if err != nil {
		return nil, err
	}
	return &Sns{Name: snsName, URL: snsURL, Icon: icon}, nil
}
