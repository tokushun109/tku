package creator

import (
	"bytes"
	"context"
	"errors"
	"io"
	"testing"
	"time"

	domain "github.com/tokushun109/tku/backend/internal/domain/creator"
	"github.com/tokushun109/tku/backend/internal/domain/primitive"
	"github.com/tokushun109/tku/backend/internal/usecase"
)

type stubCreatorRepository struct {
	findRes *domain.Creator
	findErr error

	updateProfileRes bool
	updateProfileErr error
	updatedProfile   *domain.Creator

	updateLogoRes         bool
	updateLogoErr         error
	updateLogoCreatorUUID primitive.UUID
	updateLogoMimeType    domain.CreatorLogoMimeType
	updateLogoPath        domain.CreatorLogoPath
}

func (s *stubCreatorRepository) Find(ctx context.Context) (*domain.Creator, error) {
	return s.findRes, s.findErr
}

func (s *stubCreatorRepository) UpdateProfile(ctx context.Context, c *domain.Creator) (bool, error) {
	s.updatedProfile = c
	return s.updateProfileRes, s.updateProfileErr
}

func (s *stubCreatorRepository) UpdateLogo(ctx context.Context, creatorUUID primitive.UUID, mimeType domain.CreatorLogoMimeType, logoPath domain.CreatorLogoPath) (bool, error) {
	s.updateLogoCreatorUUID = creatorUUID
	s.updateLogoMimeType = mimeType
	s.updateLogoPath = logoPath
	return s.updateLogoRes, s.updateLogoErr
}

type stubLogoStorage struct {
	putErr error
	getErr error
	getRes []byte

	presignErr     error
	presignURL     string
	presignExpires time.Duration

	putKey        string
	putMimeType   string
	putBinarySize int
	deletedKeys   []string
}

func (s *stubLogoStorage) Put(ctx context.Context, key string, contentType string, data []byte) error {
	s.putKey = key
	s.putMimeType = contentType
	s.putBinarySize = len(data)
	return s.putErr
}

func (s *stubLogoStorage) Get(ctx context.Context, key string) (io.ReadCloser, error) {
	if s.getErr != nil {
		return nil, s.getErr
	}
	return io.NopCloser(bytes.NewReader(s.getRes)), nil
}

func (s *stubLogoStorage) Delete(ctx context.Context, key string) error {
	s.deletedKeys = append(s.deletedKeys, key)
	return nil
}

func (s *stubLogoStorage) PresignGet(ctx context.Context, key string, expires time.Duration) (string, error) {
	s.presignExpires = expires
	if s.presignErr != nil {
		return "", s.presignErr
	}
	return s.presignURL, nil
}

type stubUUIDGenerator struct {
	uuid string
}

func (s *stubUUIDGenerator) New() string {
	return s.uuid
}

func TestServiceGet(t *testing.T) {
	t.Run("ロゴがあるならpresignedURLをapiPathに返す", func(t *testing.T) {
		creator := mustCreator(1, "作家", "紹介", "image/png", "img/logo/a/b/test.png")
		repo := &stubCreatorRepository{findRes: creator}
		storage := &stubLogoStorage{presignURL: "https://example.com/presigned"}
		uuidGen := &stubUUIDGenerator{}
		svc := New(repo, storage, uuidGen)

		detail, err := svc.Get(context.Background())
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if detail.APIPath != "https://example.com/presigned" {
			t.Fatalf("unexpected apiPath: %s", detail.APIPath)
		}
		if storage.presignExpires != defaultLogoPresignTTL {
			t.Fatalf("unexpected presign ttl: %s", storage.presignExpires)
		}
	})

	t.Run("presignedURLの生成に失敗したらエラーを返す", func(t *testing.T) {
		creator := mustCreator(1, "作家", "紹介", "image/png", "img/logo/a/b/test.png")
		repo := &stubCreatorRepository{findRes: creator}
		storage := &stubLogoStorage{presignErr: errors.New("presign failed")}
		uuidGen := &stubUUIDGenerator{}
		svc := New(repo, storage, uuidGen)

		_, err := svc.Get(context.Background())
		if err == nil || !errors.Is(err, usecase.ErrInternal) {
			t.Fatalf("expected ErrInternal, got %v", err)
		}
	})
}

func TestServiceUpdate(t *testing.T) {
	t.Run("紹介文が空白のみなら紹介文をクリアして更新に成功する", func(t *testing.T) {
		creator := mustCreator(1, "作家", "既存の紹介文", "", "")
		repo := &stubCreatorRepository{
			findRes:          creator,
			updateProfileRes: true,
			updateProfileErr: nil,
			updatedProfile:   nil,
			updateLogoRes:    false,
			updateLogoErr:    nil,
		}
		svc := New(repo, &stubLogoStorage{}, &stubUUIDGenerator{})

		err := svc.Update(context.Background(), "作家", "   ")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if repo.updatedProfile == nil {
			t.Fatalf("expected updated profile to be set")
		}
		if repo.updatedProfile.Introduction() != nil {
			t.Fatalf("expected introduction to be nil, got %#v", repo.updatedProfile.Introduction())
		}
	})
}

func TestServiceUpdateLogo(t *testing.T) {
	t.Run("有効な画像を渡したとき新規保存とDB更新に成功する", func(t *testing.T) {
		creator := mustCreator(10, "作家", "紹介", "image/jpeg", "img/logo/o/l/old.jpg")
		repo := &stubCreatorRepository{
			findRes:       creator,
			updateLogoRes: true,
		}
		storage := &stubLogoStorage{}

		uuid, err := primitive.NewUUID("123e4567-e89b-12d3-a456-426614174000")
		if err != nil {
			t.Fatalf("unexpected uuid error: %v", err)
		}
		uuidGen := &stubUUIDGenerator{uuid: uuid.Value()}
		svc := New(repo, storage, uuidGen)

		pngBinary := []byte{0x89, 'P', 'N', 'G', '\r', '\n', 0x1a, '\n', 0, 0, 0, 0}
		err = svc.UpdateLogo(context.Background(), pngBinary)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		expectedPath := "img/logo/1/2/123e4567-e89b-12d3-a456-426614174000.png"
		if storage.putKey != expectedPath {
			t.Fatalf("unexpected save path: %s", storage.putKey)
		}
		if repo.updateLogoPath.Value() != expectedPath {
			t.Fatalf("unexpected updated logo path: %s", repo.updateLogoPath.Value())
		}
		if repo.updateLogoCreatorUUID.Value() != creator.UUID().Value() {
			t.Fatalf("unexpected creator uuid: %s", repo.updateLogoCreatorUUID.Value())
		}
		if len(storage.deletedKeys) != 1 || storage.deletedKeys[0] != "img/logo/o/l/old.jpg" {
			t.Fatalf("unexpected deleted keys: %#v", storage.deletedKeys)
		}
	})

	t.Run("対応外のMIMEタイプならバリデーションエラーで失敗する", func(t *testing.T) {
		creator := mustCreator(10, "作家", "紹介", "", "")
		repo := &stubCreatorRepository{findRes: creator}
		storage := &stubLogoStorage{}
		uuidGen := &stubUUIDGenerator{}
		svc := New(repo, storage, uuidGen)

		err := svc.UpdateLogo(context.Background(), []byte("plain-text"))
		if err == nil || !errors.Is(err, usecase.ErrInvalidInput) {
			t.Fatalf("expected ErrInvalidInput, got %v", err)
		}
	})
}

func TestServiceGetLogoBlob(t *testing.T) {
	t.Run("リクエストのファイル名が保存済みロゴと一致するなら画像取得に成功する", func(t *testing.T) {
		creator := mustCreator(1, "作家", "紹介", "image/png", "img/logo/a/b/logo.png")
		repo := &stubCreatorRepository{findRes: creator}
		storage := &stubLogoStorage{getRes: []byte("binary")}
		uuidGen := &stubUUIDGenerator{}
		svc := New(repo, storage, uuidGen)

		blob, err := svc.GetLogoBlob(context.Background(), "logo.png")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		defer func() {
			_ = blob.Body.Close()
		}()
		if blob.ContentType != "image/png" {
			t.Fatalf("unexpected content type: %s", blob.ContentType)
		}
		binary, err := io.ReadAll(blob.Body)
		if err != nil {
			t.Fatalf("unexpected read error: %v", err)
		}
		if string(binary) != "binary" {
			t.Fatalf("unexpected binary: %s", string(binary))
		}
	})

	t.Run("リクエストのファイル名が不一致ならバリデーションエラーで失敗する", func(t *testing.T) {
		creator := mustCreator(1, "作家", "紹介", "image/png", "img/logo/a/b/logo.png")
		repo := &stubCreatorRepository{findRes: creator}
		storage := &stubLogoStorage{}
		uuidGen := &stubUUIDGenerator{}
		svc := New(repo, storage, uuidGen)

		_, err := svc.GetLogoBlob(context.Background(), "other.png")
		if err == nil || !errors.Is(err, usecase.ErrInvalidInput) {
			t.Fatalf("expected ErrInvalidInput, got %v", err)
		}
	})

	t.Run("ストレージ上にロゴが存在しないならNotFoundで失敗する", func(t *testing.T) {
		creator := mustCreator(1, "作家", "紹介", "image/png", "img/logo/a/b/logo.png")
		repo := &stubCreatorRepository{findRes: creator}
		storage := &stubLogoStorage{getErr: usecase.ErrStorageNotFound}
		uuidGen := &stubUUIDGenerator{}
		svc := New(repo, storage, uuidGen)

		_, err := svc.GetLogoBlob(context.Background(), "logo.png")
		if err == nil || !errors.Is(err, usecase.ErrNotFound) {
			t.Fatalf("expected ErrNotFound, got %v", err)
		}
	})
}

func mustCreator(id uint, name, introduction, mimeType, logoPath string) *domain.Creator {
	creator, err := domain.Rebuild(id, "11111111-1111-4111-8111-111111111111", name, introduction, mimeType, logoPath)
	if err != nil {
		panic(err)
	}
	return creator
}
