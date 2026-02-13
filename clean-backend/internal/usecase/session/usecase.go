package session

import (
	"context"

	domain "github.com/tokushun109/tku/clean-backend/internal/domain/session"
	"github.com/tokushun109/tku/clean-backend/internal/usecase"
)

type Usecase interface {
	Validate(ctx context.Context, token string) error
}

type Service struct {
	repo domain.Repository
}

func New(repo domain.Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Validate(ctx context.Context, token string) error {
	if token == "" {
		return usecase.NewAppError(usecase.ErrUnauthorized)
	}
	uuid, err := domain.ParseSessionUUID(token)
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
	return nil
}
