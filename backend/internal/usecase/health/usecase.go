package health

import (
	"context"

	domain "github.com/tokushun109/tku/backend/internal/domain/health"
	"github.com/tokushun109/tku/backend/internal/usecase"
)

type Usecase interface {
	Check(ctx context.Context) error
}

type Service struct {
	repo domain.Repository
}

func New(repo domain.Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Check(ctx context.Context) error {
	if err := s.repo.Ping(ctx); err != nil {
		return usecase.NewAppErrorWithMessage(usecase.ErrInternal, err.Error())
	}
	return nil
}
