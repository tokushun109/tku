package local

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/tokushun109/tku/clean-backend/internal/usecase"
)

type Storage struct {
	rootDir string
}

func NewStorage(rootDir string) *Storage {
	trimmedRoot := strings.TrimSpace(rootDir)
	if trimmedRoot == "" {
		trimmedRoot = "."
	}
	return &Storage{rootDir: trimmedRoot}
}

var _ usecase.Storage = (*Storage)(nil)

func (s *Storage) Put(ctx context.Context, key string, contentType string, data []byte) error {
	_ = contentType

	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	filePath, err := s.resolvePath(key)
	if err != nil {
		return err
	}
	if err := os.MkdirAll(filepath.Dir(filePath), 0o755); err != nil {
		return err
	}

	tmpFile, err := os.CreateTemp(filepath.Dir(filePath), ".storage-tmp-*")
	if err != nil {
		return err
	}
	tmpFilePath := tmpFile.Name()

	_, copyErr := io.Copy(tmpFile, bytes.NewReader(data))
	closeErr := tmpFile.Close()
	if copyErr != nil {
		_ = os.Remove(tmpFilePath)
		return copyErr
	}
	if closeErr != nil {
		_ = os.Remove(tmpFilePath)
		return closeErr
	}
	if err := os.Rename(tmpFilePath, filePath); err != nil {
		_ = os.Remove(tmpFilePath)
		return err
	}

	return nil
}

func (s *Storage) Get(ctx context.Context, key string) (io.ReadCloser, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}

	filePath, err := s.resolvePath(key)
	if err != nil {
		return nil, err
	}

	file, err := os.Open(filePath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil, usecase.ErrStorageNotFound
		}
		return nil, err
	}
	return file, nil
}

func (s *Storage) Delete(ctx context.Context, key string) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	filePath, err := s.resolvePath(key)
	if err != nil {
		return err
	}

	err = os.Remove(filePath)
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		return err
	}
	return nil
}

func (s *Storage) PresignGet(ctx context.Context, key string, expires time.Duration) (string, error) {
	_ = ctx
	_ = key
	_ = expires
	return "", fmt.Errorf("local storage does not support presigned url")
}

func (s *Storage) resolvePath(key string) (string, error) {
	normalized := filepath.Clean(filepath.FromSlash(strings.TrimSpace(key)))
	if normalized == "" || normalized == "." || normalized == string(filepath.Separator) {
		return "", fmt.Errorf("invalid storage key: %s", key)
	}
	if filepath.IsAbs(normalized) || normalized == ".." || strings.HasPrefix(normalized, ".."+string(filepath.Separator)) {
		return "", fmt.Errorf("invalid storage key: %s", key)
	}
	return filepath.Join(s.rootDir, normalized), nil
}
