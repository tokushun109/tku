package sns

import (
	"context"
	"errors"

	"github.com/tokushun109/tku/clean-backend/internal/domain/primitive"
	domain "github.com/tokushun109/tku/clean-backend/internal/domain/sns"
	"github.com/tokushun109/tku/clean-backend/internal/usecase"
)

type Usecase interface {
	List(ctx context.Context) ([]*domain.Sns, error)
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

func (s *Service) List(ctx context.Context) ([]*domain.Sns, error) {
	snsList, err := s.repo.FindAll(ctx)
	if err != nil {
		return nil, usecase.NewAppErrorWithMessage(usecase.ErrInternal, err.Error())
	}
	return snsList, nil
}

func (s *Service) Create(ctx context.Context, name string, rawURL string, icon string) error {
	snsObj, err := domain.New(name, rawURL, icon)
	if err != nil {
		if errors.Is(err, domain.ErrInvalidName) || errors.Is(err, primitive.ErrInvalidURL) {
			return usecase.NewAppErrorWithMessage(usecase.ErrInvalidInput, err.Error())
		}
		return usecase.NewAppErrorWithMessage(usecase.ErrInternal, err.Error())
	}

	newUUID, err := s.uuidGen.New()
	if err != nil {
		return usecase.NewAppErrorWithMessage(usecase.ErrInternal, err.Error())
	}
	snsObj.UUID = newUUID
	if err := s.repo.Create(ctx, snsObj); err != nil {
		return usecase.NewAppErrorWithMessage(usecase.ErrInternal, err.Error())
	}
	return nil
}

func (s *Service) Update(ctx context.Context, uuidStr string, name string, rawURL string, icon string) error {
	uuid, err := primitive.NewUUID(uuidStr)
	if err != nil {
		return usecase.NewAppError(usecase.ErrInvalidInput)
	}
	newName, err := domain.NewSnsName(name)
	if err != nil {
		if errors.Is(err, domain.ErrInvalidName) {
			return usecase.NewAppErrorWithMessage(usecase.ErrInvalidInput, err.Error())
		}
		return usecase.NewAppErrorWithMessage(usecase.ErrInternal, err.Error())
	}
	newURL, err := primitive.NewURL(rawURL)
	if err != nil {
		if errors.Is(err, primitive.ErrInvalidURL) {
			return usecase.NewAppErrorWithMessage(usecase.ErrInvalidInput, err.Error())
		}
		return usecase.NewAppErrorWithMessage(usecase.ErrInternal, err.Error())
	}

	current, err := s.repo.FindByUUID(ctx, uuid)
	if err != nil {
		return usecase.NewAppErrorWithMessage(usecase.ErrInternal, err.Error())
	}
	if current == nil {
		return usecase.NewAppError(usecase.ErrNotFound)
	}

	updated, err := s.repo.Update(ctx, &domain.Sns{UUID: uuid, Name: newName, URL: newURL, Icon: icon})
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
