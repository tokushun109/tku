package s3

import (
	"bytes"
	"context"
	"errors"
	"io"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsconfig "github.com/aws/aws-sdk-go-v2/config"
	awss3 "github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/aws/smithy-go"

	"github.com/tokushun109/tku/clean-backend/internal/usecase"
)

type Storage struct {
	bucket        string
	client        *awss3.Client
	presignClient *awss3.PresignClient
}

var _ usecase.Storage = (*Storage)(nil)

func NewStorage(bucket string) (*Storage, error) {
	trimmedBucket := strings.TrimSpace(bucket)
	if trimmedBucket == "" {
		return nil, errors.New("api bucket name is empty")
	}

	cfg, err := awsconfig.LoadDefaultConfig(context.Background())
	if err != nil {
		return nil, err
	}

	client := awss3.NewFromConfig(cfg)
	return &Storage{
		bucket:        trimmedBucket,
		client:        client,
		presignClient: awss3.NewPresignClient(client),
	}, nil
}

func (s *Storage) Put(ctx context.Context, key string, contentType string, data []byte) error {
	_, err := s.client.PutObject(ctx, &awss3.PutObjectInput{
		Bucket:      aws.String(s.bucket),
		Key:         aws.String(key),
		Body:        bytes.NewReader(data),
		ContentType: aws.String(contentType),
	})
	return err
}

func (s *Storage) Get(ctx context.Context, key string) (io.ReadCloser, error) {
	output, err := s.client.GetObject(ctx, &awss3.GetObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		return nil, mapS3NotFoundError(err)
	}
	return output.Body, nil
}

func (s *Storage) Delete(ctx context.Context, key string) error {
	_, err := s.client.DeleteObject(ctx, &awss3.DeleteObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(key),
	})
	return err
}

func (s *Storage) PresignGet(ctx context.Context, key string, expires time.Duration) (string, error) {
	req, err := s.presignClient.PresignGetObject(
		ctx,
		&awss3.GetObjectInput{
			Bucket: aws.String(s.bucket),
			Key:    aws.String(key),
		},
		awss3.WithPresignExpires(expires),
	)
	if err != nil {
		return "", err
	}
	return req.URL, nil
}

func mapS3NotFoundError(err error) error {
	var noSuchKey *types.NoSuchKey
	if errors.As(err, &noSuchKey) {
		return usecase.ErrStorageNotFound
	}

	var apiErr smithy.APIError
	if errors.As(err, &apiErr) {
		code := strings.TrimSpace(apiErr.ErrorCode())
		if code == "NoSuchKey" || code == "NotFound" {
			return usecase.ErrStorageNotFound
		}
	}
	return err
}
