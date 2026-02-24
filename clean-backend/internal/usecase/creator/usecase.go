package creator

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"path"
	"strings"
	"time"

	domain "github.com/tokushun109/tku/clean-backend/internal/domain/creator"
	"github.com/tokushun109/tku/clean-backend/internal/domain/primitive"
	"github.com/tokushun109/tku/clean-backend/internal/usecase"
)

const (
	defaultLogoPresignTTL = 30 * time.Minute
)

type Usecase interface {
	Get(ctx context.Context) (*CreatorDetail, error)
	Update(ctx context.Context, name string, introduction string) error
	UpdateLogo(ctx context.Context, logoBytes []byte) error
	GetLogoBlob(ctx context.Context, requestLogoFile string) (*LogoBlob, error)
}

type CreatorDetail struct {
	Creator *domain.Creator
	APIPath string
}

type LogoBlob struct {
	ContentType string
	Body        io.ReadCloser
}

type Service struct {
	repo    domain.Repository
	storage usecase.Storage
	uuidGen usecase.UUIDGenerator
}

func New(
	repo domain.Repository,
	storage usecase.Storage,
	uuidGen usecase.UUIDGenerator,
) *Service {
	return &Service{
		repo:    repo,
		storage: storage,
		uuidGen: uuidGen,
	}
}

func (s *Service) Get(ctx context.Context) (*CreatorDetail, error) {
	current, err := s.repo.Find(ctx)
	if err != nil {
		return nil, usecase.NewAppErrorWithMessage(usecase.ErrInternal, err.Error())
	}
	if current == nil {
		return nil, usecase.NewAppErrorWithMessage(usecase.ErrNotFound, domain.ErrCreatorRecordMissing.Error())
	}

	logoAPIPath := ""
	if current.LogoPath != nil {
		url, err := s.storage.PresignGet(ctx, current.LogoPath.String(), defaultLogoPresignTTL)
		if err != nil {
			return nil, usecase.NewAppErrorWithMessage(usecase.ErrInternal, err.Error())
		}
		logoAPIPath = url
	}

	return &CreatorDetail{
		Creator: current,
		APIPath: logoAPIPath,
	}, nil
}

func (s *Service) Update(ctx context.Context, name string, introduction string) error {
	creatorName, err := domain.NewCreatorName(name)
	if err != nil {
		if errors.Is(err, domain.ErrInvalidName) {
			return usecase.NewAppErrorWithMessage(usecase.ErrInvalidInput, err.Error())
		}
		return usecase.NewAppErrorWithMessage(usecase.ErrInternal, err.Error())
	}
	creatorIntroduction, err := domain.NewCreatorIntroduction(introduction)
	if err != nil {
		if errors.Is(err, domain.ErrInvalidIntroduction) {
			return usecase.NewAppErrorWithMessage(usecase.ErrInvalidInput, err.Error())
		}
		return usecase.NewAppErrorWithMessage(usecase.ErrInternal, err.Error())
	}

	current, err := s.repo.Find(ctx)
	if err != nil {
		return usecase.NewAppErrorWithMessage(usecase.ErrInternal, err.Error())
	}
	if current == nil {
		return usecase.NewAppErrorWithMessage(usecase.ErrNotFound, domain.ErrCreatorRecordMissing.Error())
	}

	updated, err := s.repo.UpdateProfile(ctx, &domain.Creator{
		ID:           current.ID,
		Name:         creatorName,
		Introduction: creatorIntroduction,
		LogoMimeType: current.LogoMimeType,
		LogoPath:     current.LogoPath,
	})
	if err != nil {
		return usecase.NewAppErrorWithMessage(usecase.ErrInternal, err.Error())
	}
	if !updated {
		return usecase.NewAppErrorWithMessage(usecase.ErrNotFound, domain.ErrCreatorRecordMissing.Error())
	}

	return nil
}

func (s *Service) UpdateLogo(ctx context.Context, logoBytes []byte) error {
	if len(logoBytes) == 0 {
		return usecase.NewAppError(usecase.ErrInvalidInput)
	}

	current, err := s.repo.Find(ctx)
	if err != nil {
		return usecase.NewAppErrorWithMessage(usecase.ErrInternal, err.Error())
	}
	if current == nil {
		return usecase.NewAppErrorWithMessage(usecase.ErrNotFound, domain.ErrCreatorRecordMissing.Error())
	}

	detectedMimeType := http.DetectContentType(logoBytes)
	logoMimeType, err := domain.NewCreatorLogoMimeType(detectedMimeType)
	if err != nil {
		if errors.Is(err, domain.ErrInvalidLogoMimeType) {
			return usecase.NewAppErrorWithMessage(usecase.ErrInvalidInput, err.Error())
		}
		return usecase.NewAppErrorWithMessage(usecase.ErrInternal, err.Error())
	}

	newUUID, err := s.uuidGen.New()
	if err != nil {
		return usecase.NewAppErrorWithMessage(usecase.ErrInternal, err.Error())
	}

	newLogoPath, err := buildCreatorLogoPath(newUUID, logoMimeType)
	if err != nil {
		return usecase.NewAppErrorWithMessage(usecase.ErrInternal, err.Error())
	}

	if err := s.storage.Put(ctx, newLogoPath.String(), logoMimeType.String(), logoBytes); err != nil {
		return usecase.NewAppErrorWithMessage(usecase.ErrInternal, err.Error())
	}

	updated, err := s.repo.UpdateLogo(ctx, current.ID, logoMimeType, newLogoPath)
	if err != nil {
		if delErr := s.storage.Delete(ctx, newLogoPath.String()); delErr != nil {
			log.Printf("[WARN] creator update logo rollback delete failed: path=%s err=%v", newLogoPath.String(), delErr)
		}
		return usecase.NewAppErrorWithMessage(usecase.ErrInternal, err.Error())
	}
	if !updated {
		if delErr := s.storage.Delete(ctx, newLogoPath.String()); delErr != nil {
			log.Printf("[WARN] creator update logo rollback delete failed: path=%s err=%v", newLogoPath.String(), delErr)
		}
		return usecase.NewAppErrorWithMessage(usecase.ErrNotFound, domain.ErrCreatorRecordMissing.Error())
	}

	// DB の更新後に旧ファイルを削除する。削除失敗は orphan を許容して成功扱いにする。
	if current.LogoPath != nil && current.LogoPath.String() != newLogoPath.String() {
		if delErr := s.storage.Delete(ctx, current.LogoPath.String()); delErr != nil {
			log.Printf("[WARN] creator update logo old file delete failed: path=%s err=%v", current.LogoPath.String(), delErr)
		}
	}

	return nil
}

func (s *Service) GetLogoBlob(ctx context.Context, requestLogoFile string) (*LogoBlob, error) {
	logoFile := strings.TrimSpace(requestLogoFile)
	if logoFile == "" {
		return nil, usecase.NewAppErrorWithMessage(usecase.ErrInvalidInput, domain.ErrInvalidLogoFileName.Error())
	}

	current, err := s.repo.Find(ctx)
	if err != nil {
		return nil, usecase.NewAppErrorWithMessage(usecase.ErrInternal, err.Error())
	}
	if current == nil {
		return nil, usecase.NewAppErrorWithMessage(usecase.ErrNotFound, domain.ErrCreatorRecordMissing.Error())
	}
	if current.LogoPath == nil || current.LogoMimeType == nil {
		return nil, usecase.NewAppError(usecase.ErrNotFound)
	}

	savedLogoFile := path.Base(current.LogoPath.String())
	if savedLogoFile != logoFile {
		return nil, usecase.NewAppErrorWithMessage(usecase.ErrInvalidInput, domain.ErrInvalidLogoFileName.Error())
	}

	logoBody, err := s.storage.Get(ctx, current.LogoPath.String())
	if err != nil {
		if errors.Is(err, usecase.ErrStorageNotFound) {
			return nil, usecase.NewAppError(usecase.ErrNotFound)
		}
		return nil, usecase.NewAppErrorWithMessage(usecase.ErrInternal, err.Error())
	}

	return &LogoBlob{
		ContentType: current.LogoMimeType.String(),
		Body:        logoBody,
	}, nil
}

func buildCreatorLogoPath(uuid primitive.UUID, mimeType domain.CreatorLogoMimeType) (domain.CreatorLogoPath, error) {
	uuidStr := uuid.String()
	if len(uuidStr) < 2 {
		return "", fmt.Errorf("invalid uuid length: %d", len(uuidStr))
	}

	rawPath := fmt.Sprintf(
		"img/logo/%s/%s/%s%s",
		uuidStr[0:1],
		uuidStr[1:2],
		uuidStr,
		mimeType.Extension(),
	)

	return domain.NewCreatorLogoPath(rawPath)
}
