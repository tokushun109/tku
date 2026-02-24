package usecase

import (
	"context"
	"errors"
	"time"
)

var ErrStorageNotFound = errors.New("storage object is not found")

type Storage interface {
	Put(ctx context.Context, key string, contentType string, data []byte) error
	Get(ctx context.Context, key string) ([]byte, error)
	Delete(ctx context.Context, key string) error
	PresignGet(ctx context.Context, key string, expires time.Duration) (string, error)
}
