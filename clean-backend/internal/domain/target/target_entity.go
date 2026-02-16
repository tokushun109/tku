package target

import "github.com/tokushun109/tku/clean-backend/internal/domain/primitive"

type Target struct {
	UUID primitive.UUID
	Name TargetName
}

func New(name string) (*Target, error) {
	targetName, err := NewTargetName(name)
	if err != nil {
		return nil, err
	}
	return &Target{Name: targetName}, nil
}
