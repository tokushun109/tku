package category

import (
	"context"
	"errors"

	domain "github.com/tokushun109/tku/clean-backend/internal/domain/category"
	"github.com/tokushun109/tku/clean-backend/internal/shared/id"
	"github.com/tokushun109/tku/clean-backend/internal/usecase"
)

type Usecase interface {
	List(ctx context.Context, mode string) ([]*domain.Category, error)
	Create(ctx context.Context, name string) error
}

type Service struct {
	repo domain.Repository
}

func New(repo domain.Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) List(ctx context.Context, mode string) ([]*domain.Category, error) {
	switch mode {
	case "all":
		categories, err := s.repo.FindAll(ctx)
		if err != nil {
			return nil, usecase.NewAppErrorWithMessage(usecase.ErrInternal, err.Error())
		}
		return categories, nil
	case "used":
		categories, err := s.repo.FindUsed(ctx)
		if err != nil {
			return nil, usecase.NewAppErrorWithMessage(usecase.ErrInternal, err.Error())
		}
		return categories, nil
	default:
		return nil, usecase.NewAppError(usecase.ErrInvalidInput)
	}
}

func (s *Service) Create(ctx context.Context, name string) error {
	c, err := domain.New(name)
	if err != nil {
		if errors.Is(err, domain.ErrInvalidName) {
			return usecase.NewAppErrorWithMessage(usecase.ErrInvalidInput, err.Error())
		}
		return usecase.NewAppError(usecase.ErrInternal)
	}

	exists, err := s.repo.ExistsByName(ctx, c.Name)
	if err != nil {
		return usecase.NewAppErrorWithMessage(usecase.ErrInternal, err.Error())
	}
	if exists {
		return usecase.NewAppErrorWithMessage(usecase.ErrConflict, domain.ErrNameDuplicated.Error())
	}

	c.ID = id.NewUUID()
	if err := s.repo.Create(ctx, c); err != nil {
		return usecase.NewAppErrorWithMessage(usecase.ErrInternal, err.Error())
	}
	return nil
}
