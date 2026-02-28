package sales_site

import (
	"context"
	"errors"

	"github.com/tokushun109/tku/clean-backend/internal/domain/primitive"
	domain "github.com/tokushun109/tku/clean-backend/internal/domain/sales_site"
	"github.com/tokushun109/tku/clean-backend/internal/usecase"
)

type Usecase interface {
	List(ctx context.Context) ([]*domain.SalesSite, error)
	Create(ctx context.Context, name string, rawURL string, icon string) error
	Update(ctx context.Context, uuid string, name string, rawURL string, icon string) error
	Delete(ctx context.Context, uuid string) error
}

type Service struct {
	repo    domain.Repository
	uuidGen usecase.UUIDGenerator
}

func New(repo domain.Repository, uuidGen usecase.UUIDGenerator) *Service {
	return &Service{repo: repo, uuidGen: uuidGen}
}

func (s *Service) List(ctx context.Context) ([]*domain.SalesSite, error) {
	salesSites, err := s.repo.FindAll(ctx)
	if err != nil {
		return nil, usecase.NewAppErrorWithMessage(usecase.ErrInternal, err.Error())
	}
	return salesSites, nil
}

func (s *Service) Create(ctx context.Context, name string, rawURL string, icon string) error {
	newUUID := s.uuidGen.New()

	salesSite, err := domain.New(newUUID, name, rawURL, icon)
	if err != nil {
		if errors.Is(err, domain.ErrInvalidName) || errors.Is(err, primitive.ErrInvalidURL) {
			return usecase.NewAppErrorWithMessage(usecase.ErrInvalidInput, err.Error())
		}
		return usecase.NewAppErrorWithMessage(usecase.ErrInternal, err.Error())
	}

	if _, err := s.repo.Create(ctx, salesSite); err != nil {
		return usecase.NewAppErrorWithMessage(usecase.ErrInternal, err.Error())
	}
	return nil
}

func (s *Service) Update(ctx context.Context, uuidStr string, name string, rawURL string, icon string) error {
	uuid, err := primitive.NewUUID(uuidStr)
	if err != nil {
		return usecase.NewAppError(usecase.ErrInvalidInput)
	}
	current, err := s.repo.FindByUUID(ctx, uuid)
	if err != nil {
		return usecase.NewAppErrorWithMessage(usecase.ErrInternal, err.Error())
	}
	if current == nil {
		return usecase.NewAppError(usecase.ErrNotFound)
	}

	if err := current.Change(name, rawURL, icon); err != nil {
		if errors.Is(err, domain.ErrInvalidName) || errors.Is(err, primitive.ErrInvalidURL) {
			return usecase.NewAppErrorWithMessage(usecase.ErrInvalidInput, err.Error())
		}
		return usecase.NewAppErrorWithMessage(usecase.ErrInternal, err.Error())
	}

	updated, err := s.repo.Update(ctx, current)
	if err != nil {
		return usecase.NewAppErrorWithMessage(usecase.ErrInternal, err.Error())
	}
	if !updated {
		return usecase.NewAppError(usecase.ErrNotFound)
	}
	return nil
}

func (s *Service) Delete(ctx context.Context, uuidStr string) error {
	uuid, err := primitive.NewUUID(uuidStr)
	if err != nil {
		return usecase.NewAppError(usecase.ErrInvalidInput)
	}

	deleted, err := s.repo.Delete(ctx, uuid)
	if err != nil {
		return usecase.NewAppErrorWithMessage(usecase.ErrInternal, err.Error())
	}
	if !deleted {
		return usecase.NewAppError(usecase.ErrNotFound)
	}
	return nil
}
