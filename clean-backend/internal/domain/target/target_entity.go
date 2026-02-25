package target

import "github.com/tokushun109/tku/clean-backend/internal/domain/primitive"

type Target struct {
	id   uint
	uuid primitive.UUID
	name TargetName
}

func New(rawUUID string, name string) (*Target, error) {
	target, err := newWithValidatedValues(rawUUID, name)
	if err != nil {
		return nil, err
	}
	return target, nil
}

func Rebuild(id uint, rawUUID string, name string) (*Target, error) {
	if id == 0 {
		return nil, ErrInvalidID
	}
	target, err := newWithValidatedValues(rawUUID, name)
	if err != nil {
		return nil, err
	}
	target.id = id
	return target, nil
}

func newWithValidatedValues(rawUUID string, name string) (*Target, error) {
	uuid, err := primitive.NewUUID(rawUUID)
	if err != nil {
		return nil, err
	}
	targetName, err := NewTargetName(name)
	if err != nil {
		return nil, err
	}
	return &Target{
		uuid: uuid,
		name: targetName,
	}, nil
}

func (t *Target) ID() uint {
	return t.id
}

func (t *Target) UUID() primitive.UUID {
	return t.uuid
}

func (t *Target) Name() TargetName {
	return t.name
}

func (t *Target) ChangeName(name string) error {
	targetName, err := NewTargetName(name)
	if err != nil {
		return err
	}
	t.name = targetName
	return nil
}
