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
	Resolve(ctx context.Context, token string) (*domain.Session, error)
	Delete(ctx context.Context, token string) error
}

type Service struct {
	repo  domain.Repository
	ttl   time.Duration
	clock usecase.Clock
}

func New(repo domain.Repository, ttl time.Duration, clock usecase.Clock) *Service {
	return &Service{repo: repo, ttl: ttl, clock: clock}
}

func (s *Service) Validate(ctx context.Context, token string) error {
	_, err := s.Resolve(ctx, token)
	return err
}

func (s *Service) Resolve(ctx context.Context, token string) (*domain.Session, error) {
	uuid, err := parseToken(token)
	if err != nil {
		return nil, err
	}
	sess, err := s.repo.FindByUUID(ctx, uuid)
	if err != nil {
		return nil, usecase.NewAppErrorWithMessage(usecase.ErrInternal, err.Error())
	}
	if sess == nil {
		return nil, usecase.NewAppError(usecase.ErrUnauthorized)
	}
	if sess.CreatedAt.IsZero() {
		return nil, usecase.NewAppError(usecase.ErrUnauthorized)
	}
	if s.ttl > 0 && s.clock.Now().After(sess.CreatedAt.Add(s.ttl)) {
		if err := s.repo.DeleteByUUID(ctx, uuid); err != nil {
			return nil, usecase.NewAppErrorWithMessage(usecase.ErrInternal, err.Error())
		}
		return nil, usecase.NewAppError(usecase.ErrUnauthorized)
	}
	return sess, nil
}

func (s *Service) Delete(ctx context.Context, token string) error {
	sess, err := s.Resolve(ctx, token)
	if err != nil {
		return err
	}
	if err := s.repo.DeleteByUUID(ctx, sess.UUID); err != nil {
		return usecase.NewAppErrorWithMessage(usecase.ErrInternal, err.Error())
	}
	return nil
}

func parseToken(token string) (primitive.UUID, error) {
	if token == "" {
		return "", usecase.NewAppError(usecase.ErrUnauthorized)
	}
	uuid, err := primitive.NewUUID(token)
	if err != nil {
		return "", usecase.NewAppError(usecase.ErrUnauthorized)
	}
	return uuid, nil
}
