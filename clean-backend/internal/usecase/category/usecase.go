package category

import (
	"context"
	"errors"

	domain "github.com/tokushun109/tku/clean-backend/internal/domain/category"
	"github.com/tokushun109/tku/clean-backend/internal/domain/primitive"
	"github.com/tokushun109/tku/clean-backend/internal/usecase"
)

type Usecase interface {
	List(ctx context.Context, mode string) ([]*domain.Category, error)
	Create(ctx context.Context, name string) error
	Update(ctx context.Context, uuid string, name string) error
}

type Service struct {
	repo    domain.Repository
	uuidGen usecase.UUIDGenerator
}

func New(repo domain.Repository, uuidGen usecase.UUIDGenerator) *Service {
	return &Service{repo: repo, uuidGen: uuidGen}
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

	newUUID, err := s.uuidGen.New()
	if err != nil {
		return usecase.NewAppErrorWithMessage(usecase.ErrInternal, err.Error())
	}
	c.UUID = newUUID
	if err := s.repo.Create(ctx, c); err != nil {
		return usecase.NewAppErrorWithMessage(usecase.ErrInternal, err.Error())
	}
	return nil
}

func (s *Service) Update(ctx context.Context, uuidStr string, name string) error {
	uuid, err := primitive.NewUUID(uuidStr)
	if err != nil {
		return usecase.NewAppError(usecase.ErrInvalidInput)
	}
	newName, err := domain.NewCategoryName(name)
	if err != nil {
		if errors.Is(err, domain.ErrInvalidName) {
			return usecase.NewAppErrorWithMessage(usecase.ErrInvalidInput, err.Error())
		}
		return usecase.NewAppError(usecase.ErrInternal)
	}

	current, err := s.repo.FindByUUID(ctx, uuid)
	if err != nil {
		return usecase.NewAppErrorWithMessage(usecase.ErrInternal, err.Error())
	}
	if current == nil {
		return usecase.NewAppError(usecase.ErrNotFound)
	}

	if current.Name.String() != newName.String() {
		exists, err := s.repo.ExistsByName(ctx, newName)
		if err != nil {
			return usecase.NewAppErrorWithMessage(usecase.ErrInternal, err.Error())
		}
		if exists {
			return usecase.NewAppErrorWithMessage(usecase.ErrConflict, domain.ErrNameDuplicated.Error())
		}
	}

	updated, err := s.repo.Update(ctx, &domain.Category{UUID: uuid, Name: newName})
	if err != nil {
		return usecase.NewAppErrorWithMessage(usecase.ErrInternal, err.Error())
	}
	if !updated {
		return usecase.NewAppError(usecase.ErrNotFound)
	}
	return nil
}
