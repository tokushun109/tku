package creator

import (
	"context"
	"errors"
	"fmt"
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
	Binary      []byte
}

type Service struct {
	repo       domain.Repository
	storage    usecase.Storage
	uuidGen    usecase.UUIDGenerator
	env        string
	apiBaseURL string
	logoURLTTL time.Duration
}

func New(
	repo domain.Repository,
	storage usecase.Storage,
	uuidGen usecase.UUIDGenerator,
	env string,
	apiBaseURL string,
	logoURLTTL time.Duration,
) *Service {
	if logoURLTTL <= 0 {
		logoURLTTL = defaultLogoPresignTTL
	}
	return &Service{
		repo:       repo,
		storage:    storage,
		uuidGen:    uuidGen,
		env:        env,
		apiBaseURL: apiBaseURL,
		logoURLTTL: logoURLTTL,
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
		if isLocalEnv(s.env) {
			logoAPIPath = buildLocalLogoAPIPath(s.apiBaseURL, current.LogoPath.String())
		} else {
			url, err := s.storage.PresignGet(ctx, current.LogoPath.String(), s.logoURLTTL)
			if err != nil {
				return nil, usecase.NewAppErrorWithMessage(usecase.ErrInternal, err.Error())
			}
			logoAPIPath = url
		}
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
		_ = s.storage.Delete(ctx, newLogoPath.String())
		return usecase.NewAppErrorWithMessage(usecase.ErrInternal, err.Error())
	}
	if !updated {
		_ = s.storage.Delete(ctx, newLogoPath.String())
		return usecase.NewAppErrorWithMessage(usecase.ErrNotFound, domain.ErrCreatorRecordMissing.Error())
	}

	// DB の更新後に旧ファイルを削除する。削除失敗は orphan を許容して成功扱いにする。
	if current.LogoPath != nil && current.LogoPath.String() != newLogoPath.String() {
		_ = s.storage.Delete(ctx, current.LogoPath.String())
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

	logoBinary, err := s.storage.Get(ctx, current.LogoPath.String())
	if err != nil {
		if errors.Is(err, usecase.ErrStorageNotFound) {
			return nil, usecase.NewAppError(usecase.ErrNotFound)
		}
		return nil, usecase.NewAppErrorWithMessage(usecase.ErrInternal, err.Error())
	}

	return &LogoBlob{
		ContentType: current.LogoMimeType.String(),
		Binary:      logoBinary,
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

func buildLocalLogoAPIPath(apiBaseURL string, logoPath string) string {
	logoFile := path.Base(logoPath)
	base := strings.TrimRight(strings.TrimSpace(apiBaseURL), "/")
	if base == "" {
		return "/api/creator/logo/" + logoFile + "/blob"
	}
	return base + "/creator/logo/" + logoFile + "/blob"
}

func isLocalEnv(env string) bool {
	normalized := strings.ToLower(strings.TrimSpace(env))
	return normalized == "" || normalized == "local"
}
