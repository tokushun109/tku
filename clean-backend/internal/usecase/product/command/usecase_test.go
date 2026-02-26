package command

import (
	"context"
	"errors"
	"testing"

	"github.com/tokushun109/tku/clean-backend/internal/domain/primitive"
	domainProduct "github.com/tokushun109/tku/clean-backend/internal/domain/product"
	"github.com/tokushun109/tku/clean-backend/internal/shared/id"
	"github.com/tokushun109/tku/clean-backend/internal/usecase"
	usecaseProduct "github.com/tokushun109/tku/clean-backend/internal/usecase/product"
)

type stubProductRepoForCreateImages struct {
	findByUUIDRes *domainProduct.Product
	findByUUIDErr error

	findByUUIDCalled int
}

func (s *stubProductRepoForCreateImages) Create(ctx context.Context, p *domainProduct.Product) (primitive.ID, error) {
	return 0, nil
}

func (s *stubProductRepoForCreateImages) FindByUUID(ctx context.Context, uuid primitive.UUID) (*domainProduct.Product, error) {
	s.findByUUIDCalled++
	if s.findByUUIDErr != nil {
		return nil, s.findByUUIDErr
	}
	return s.findByUUIDRes, nil
}

func (s *stubProductRepoForCreateImages) Update(ctx context.Context, p *domainProduct.Product) (bool, error) {
	return false, nil
}

func (s *stubProductRepoForCreateImages) Delete(ctx context.Context, uuid primitive.UUID) (bool, error) {
	return false, nil
}

func (s *stubProductRepoForCreateImages) ReplaceTags(ctx context.Context, productID primitive.ID, tagIDs []primitive.ID) error {
	return nil
}

type stubTxManager struct {
	called bool
}

func (s *stubTxManager) WithinTransaction(ctx context.Context, fn func(ctx context.Context) error) error {
	s.called = true
	return fn(ctx)
}

func TestCreateProductImagesWithEmptyFiles(t *testing.T) {
	t.Run("filesが空でもproductUUIDが不正ならバリデーションエラーで失敗する", func(t *testing.T) {
		repo := &stubProductRepoForCreateImages{}
		s := &Service{productRepo: repo}

		err := s.CreateProductImages(context.Background(), "invalid-uuid", []usecaseProduct.ProductImageUploadFile{}, false, nil)
		if err == nil || !errors.Is(err, usecase.ErrInvalidInput) {
			t.Fatalf("expected ErrInvalidInput, got %v", err)
		}
		if repo.findByUUIDCalled != 0 {
			t.Fatalf("expected FindByUUID not called, got %d", repo.findByUUIDCalled)
		}
	})

	t.Run("filesが空でも商品が存在しないならNotFoundエラーを返す", func(t *testing.T) {
		repo := &stubProductRepoForCreateImages{findByUUIDRes: nil}
		s := &Service{productRepo: repo}

		err := s.CreateProductImages(context.Background(), id.GenerateUUID(), []usecaseProduct.ProductImageUploadFile{}, false, nil)
		if err == nil || !errors.Is(err, usecase.ErrNotFound) {
			t.Fatalf("expected ErrNotFound, got %v", err)
		}
		if repo.findByUUIDCalled != 1 {
			t.Fatalf("expected FindByUUID called once, got %d", repo.findByUUIDCalled)
		}
	})

	t.Run("filesが空でも商品が存在するなら成功する", func(t *testing.T) {
		productUUID := id.GenerateUUID()
		product, err := domainProduct.New(
			productUUID,
			"sample product",
			"",
			1000,
			true,
			false,
			nil,
			nil,
		)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		repo := &stubProductRepoForCreateImages{findByUUIDRes: product}
		s := &Service{productRepo: repo}

		err = s.CreateProductImages(context.Background(), productUUID, []usecaseProduct.ProductImageUploadFile{}, false, nil)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if repo.findByUUIDCalled != 1 {
			t.Fatalf("expected FindByUUID called once, got %d", repo.findByUUIDCalled)
		}
	})
}

func TestCreateProductImages(t *testing.T) {
	t.Run("画像MIMEが不正なときバリデーションエラーを返す", func(t *testing.T) {
		productUUID := id.GenerateUUID()
		product := mustNewProduct(t, productUUID)
		repo := &stubProductRepoForCreateImages{findByUUIDRes: product}
		tx := &stubTxManager{}
		s := &Service{
			productRepo: repo,
			txManager:   tx,
		}

		err := s.CreateProductImages(context.Background(), productUUID, []usecaseProduct.ProductImageUploadFile{
			{
				Name: "plain.txt",
				Data: []byte("this is plain text"),
			},
		}, false, nil)
		if err == nil || !errors.Is(err, usecase.ErrInvalidInput) {
			t.Fatalf("expected ErrInvalidInput, got %v", err)
		}
		if !tx.called {
			t.Fatalf("expected transaction called")
		}
	})
}

func mustNewProduct(t *testing.T, uuid string) *domainProduct.Product {
	t.Helper()

	product, err := domainProduct.New(
		uuid,
		"sample product",
		"",
		1000,
		true,
		false,
		nil,
		nil,
	)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	return product
}
