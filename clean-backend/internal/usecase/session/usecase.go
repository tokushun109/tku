package session

import (
	"context"
	"time"

	"github.com/tokushun109/tku/clean-backend/internal/domain/primitive"
	domain "github.com/tokushun109/tku/clean-backend/internal/domain/session"
	"github.com/tokushun109/tku/clean-backend/internal/usecase"
)

type Usecase interface {
	Validate(ctx context.Context, token string) error
}

type Service struct {
	repo domain.Repository
	ttl  time.Duration
	clock usecase.Clock
}

func New(repo domain.Repository, ttl time.Duration, clock usecase.Clock) *Service {
	if clock == nil {
		clock = systemClock{}
	}
	return &Service{repo: repo, ttl: ttl, clock: clock}
}

func (s *Service) Validate(ctx context.Context, token string) error {
	if token == "" {
		return usecase.NewAppError(usecase.ErrUnauthorized)
	}
	uuid, err := primitive.NewUUID(token)
	if err != nil {
		return usecase.NewAppError(usecase.ErrUnauthorized)
	}
	sess, err := s.repo.FindByUUID(ctx, uuid)
	if err != nil {
		return usecase.NewAppErrorWithMessage(usecase.ErrInternal, err.Error())
	}
	if sess == nil {
		return usecase.NewAppError(usecase.ErrUnauthorized)
	}
	if sess.CreatedAt.IsZero() {
		return usecase.NewAppError(usecase.ErrUnauthorized)
	}
	if s.ttl > 0 && s.clock.Now().After(sess.CreatedAt.Add(s.ttl)) {
		if err := s.repo.DeleteByUUID(ctx, uuid); err != nil {
			return usecase.NewAppErrorWithMessage(usecase.ErrInternal, err.Error())
		}
		return usecase.NewAppError(usecase.ErrUnauthorized)
	}
	return nil
}

type systemClock struct{}

func (systemClock) Now() time.Time {
	return time.Now()
}
