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
	if current.LogoPath() != nil {
		url, err := s.storage.PresignGet(ctx, current.LogoPath().Value(), defaultLogoPresignTTL)
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
	current, err := s.repo.Find(ctx)
	if err != nil {
		return usecase.NewAppErrorWithMessage(usecase.ErrInternal, err.Error())
	}
	if current == nil {
		return usecase.NewAppErrorWithMessage(usecase.ErrNotFound, domain.ErrCreatorRecordMissing.Error())
	}

	if err := current.ChangeProfile(name, introduction); err != nil {
		if errors.Is(err, domain.ErrInvalidName) || errors.Is(err, domain.ErrInvalidIntroduction) {
			return usecase.NewAppErrorWithMessage(usecase.ErrInvalidInput, err.Error())
		}
		return usecase.NewAppErrorWithMessage(usecase.ErrInternal, err.Error())
	}

	updated, err := s.repo.UpdateProfile(ctx, current)
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

	newUUID := s.uuidGen.New()

	newLogoPath, err := buildCreatorLogoPath(newUUID, logoMimeType)
	if err != nil {
		return usecase.NewAppErrorWithMessage(usecase.ErrInternal, err.Error())
	}

	if err := s.storage.Put(ctx, newLogoPath.Value(), logoMimeType.Value(), logoBytes); err != nil {
		return usecase.NewAppErrorWithMessage(usecase.ErrInternal, err.Error())
	}

	updated, err := s.repo.UpdateLogo(ctx, current.ID(), logoMimeType, newLogoPath)
	if err != nil {
		if delErr := s.storage.Delete(ctx, newLogoPath.Value()); delErr != nil {
			log.Printf("[WARN] creator update logo rollback delete failed: path=%s err=%v", newLogoPath.Value(), delErr)
		}
		return usecase.NewAppErrorWithMessage(usecase.ErrInternal, err.Error())
	}
	if !updated {
		if delErr := s.storage.Delete(ctx, newLogoPath.Value()); delErr != nil {
			log.Printf("[WARN] creator update logo rollback delete failed: path=%s err=%v", newLogoPath.Value(), delErr)
		}
		return usecase.NewAppErrorWithMessage(usecase.ErrNotFound, domain.ErrCreatorRecordMissing.Error())
	}

	// DB の更新後に旧ファイルを削除する。削除失敗は orphan を許容して成功扱いにする。
	if current.LogoPath() != nil && *current.LogoPath() != newLogoPath {
		if delErr := s.storage.Delete(ctx, current.LogoPath().Value()); delErr != nil {
			log.Printf("[WARN] creator update logo old file delete failed: path=%s err=%v", current.LogoPath().Value(), delErr)
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
	if current.LogoPath() == nil || current.LogoMimeType() == nil {
		return nil, usecase.NewAppError(usecase.ErrNotFound)
	}

	savedLogoFile := path.Base(current.LogoPath().Value())
	if savedLogoFile != logoFile {
		return nil, usecase.NewAppErrorWithMessage(usecase.ErrInvalidInput, domain.ErrInvalidLogoFileName.Error())
	}

	logoBody, err := s.storage.Get(ctx, current.LogoPath().Value())
	if err != nil {
		if errors.Is(err, usecase.ErrStorageNotFound) {
			return nil, usecase.NewAppError(usecase.ErrNotFound)
		}
		return nil, usecase.NewAppErrorWithMessage(usecase.ErrInternal, err.Error())
	}

	return &LogoBlob{
		ContentType: current.LogoMimeType().Value(),
		Body:        logoBody,
	}, nil
}

func buildCreatorLogoPath(uuidStr string, mimeType domain.CreatorLogoMimeType) (domain.CreatorLogoPath, error) {
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
