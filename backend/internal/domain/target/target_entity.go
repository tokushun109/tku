package target

import "github.com/tokushun109/tku/backend/internal/domain/primitive"

type Target struct {
	id   primitive.ID
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
	parsedID, err := primitive.NewID(id)
	if err != nil {
		return nil, ErrInvalidID
	}
	target, err := newWithValidatedValues(rawUUID, name)
	if err != nil {
		return nil, err
	}
	target.id = parsedID
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

func (t *Target) ID() primitive.ID {
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
