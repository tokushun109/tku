package session

import (
	"time"

	"github.com/tokushun109/tku/clean-backend/internal/domain/primitive"
)

type Session struct {
	id        uint
	uuid      primitive.UUID
	userID    uint
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
	if id == 0 {
		return nil, ErrInvalidID
	}
	session, err := newWithValidatedValues(rawUUID, userID, createdAt)
	if err != nil {
		return nil, err
	}
	session.id = id
	return session, nil
}

func newWithValidatedValues(rawUUID string, userID uint, createdAt time.Time) (*Session, error) {
	uuid, err := primitive.NewUUID(rawUUID)
	if err != nil {
		return nil, err
	}
	if userID == 0 {
		return nil, ErrInvalidUserID
	}
	if createdAt.IsZero() {
		return nil, ErrInvalidCreatedAt
	}
	return &Session{
		uuid:      uuid,
		userID:    userID,
		createdAt: createdAt,
	}, nil
}

func (s *Session) ID() uint {
	return s.id
}

func (s *Session) UUID() primitive.UUID {
	return s.uuid
}

func (s *Session) UserID() uint {
	return s.userID
}

func (s *Session) CreatedAt() time.Time {
	return s.createdAt
}
