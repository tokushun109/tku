package session

import (
	"time"

	"github.com/tokushun109/tku/backend/internal/domain/primitive"
)

type Session struct {
	id        primitive.ID
	uuid      primitive.UUID
	userUUID  primitive.UUID
	createdAt time.Time
}

func New(rawUUID string, userUUID string, createdAt time.Time) (*Session, error) {
	session, err := newWithValidatedValues(rawUUID, userUUID, createdAt)
	if err != nil {
		return nil, err
	}
	return session, nil
}

func Rebuild(id uint, rawUUID string, userUUID string, createdAt time.Time) (*Session, error) {
	parsedID, err := primitive.NewID(id)
	if err != nil {
		return nil, ErrInvalidID
	}
	session, err := newWithValidatedValues(rawUUID, userUUID, createdAt)
	if err != nil {
		return nil, err
	}
	session.id = parsedID
	return session, nil
}

func newWithValidatedValues(rawUUID string, userUUID string, createdAt time.Time) (*Session, error) {
	uuid, err := primitive.NewUUID(rawUUID)
	if err != nil {
		return nil, err
	}
	parsedUserUUID, err := primitive.NewUUID(userUUID)
	if err != nil {
		return nil, ErrInvalidUserUUID
	}
	if createdAt.IsZero() {
		return nil, ErrInvalidCreatedAt
	}
	return &Session{
		uuid:      uuid,
		userUUID:  parsedUserUUID,
		createdAt: createdAt,
	}, nil
}

func (s *Session) ID() primitive.ID {
	return s.id
}

func (s *Session) UUID() primitive.UUID {
	return s.uuid
}

func (s *Session) UserUUID() primitive.UUID {
	return s.userUUID
}

func (s *Session) CreatedAt() time.Time {
	return s.createdAt
}
