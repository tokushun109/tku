package session

import (
	"time"

	"github.com/tokushun109/tku/clean-backend/internal/domain/primitive"
)

type Session struct {
	id        primitive.ID
	uuid      primitive.UUID
	userID    primitive.ID
	createdAt time.Time
}

func New(rawUUID string, userID uint, createdAt time.Time) (*Session, error) {
	session, err := newWithValidatedValues(rawUUID, userID, createdAt)
	if err != nil {
		return nil, err
	}
	return session, nil
}

func Rebuild(id uint, rawUUID string, userID uint, createdAt time.Time) (*Session, error) {
	parsedID, err := primitive.NewID(id)
	if err != nil {
		return nil, ErrInvalidID
	}
	session, err := newWithValidatedValues(rawUUID, userID, createdAt)
	if err != nil {
		return nil, err
	}
	session.id = parsedID
	return session, nil
}

func newWithValidatedValues(rawUUID string, userID uint, createdAt time.Time) (*Session, error) {
	uuid, err := primitive.NewUUID(rawUUID)
	if err != nil {
		return nil, err
	}
	parsedUserID, err := primitive.NewID(userID)
	if err != nil {
		return nil, ErrInvalidUserID
	}
	if createdAt.IsZero() {
		return nil, ErrInvalidCreatedAt
	}
	return &Session{
		uuid:      uuid,
		userID:    parsedUserID,
		createdAt: createdAt,
	}, nil
}

func (s *Session) ID() primitive.ID {
	return s.id
}

func (s *Session) UUID() primitive.UUID {
	return s.uuid
}

func (s *Session) UserID() uint {
	return s.userID.Uint()
}

func (s *Session) CreatedAt() time.Time {
	return s.createdAt
}
