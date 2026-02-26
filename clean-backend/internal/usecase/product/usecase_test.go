package product

import (
	"context"
	"errors"
	"io"
	"testing"
	"time"

	"github.com/tokushun109/tku/clean-backend/internal/domain/primitive"
	domainProduct "github.com/tokushun109/tku/clean-backend/internal/domain/product"
	"github.com/tokushun109/tku/clean-backend/internal/shared/id"
	"github.com/tokushun109/tku/clean-backend/internal/usecase"
	usecaseProductQuery "github.com/tokushun109/tku/clean-backend/internal/usecase/product/query"
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

type stubProductQueryReader struct {
	listRes      []*usecaseProductQuery.Product
	listErr      error
	listByCatRes []*usecaseProductQuery.CategoryProducts
	listByCatErr error
	getRes       *usecaseProductQuery.Product
	getErr       error
}

func (s *stubProductQueryReader) ListProducts(ctx context.Context, q usecaseProductQuery.ListProductsQuery) ([]*usecaseProductQuery.Product, error) {
	if s.listErr != nil {
		return nil, s.listErr
	}
	return s.listRes, nil
}

func (s *stubProductQueryReader) ListCategoryProducts(
	ctx context.Context,
	q usecaseProductQuery.ListCategoryProductsQuery,
) ([]*usecaseProductQuery.CategoryProducts, error) {
	if s.listByCatErr != nil {
		return nil, s.listByCatErr
	}
	return s.listByCatRes, nil
}

func (s *stubProductQueryReader) GetProductByUUID(ctx context.Context, productUUID string) (*usecaseProductQuery.Product, error) {
	if s.getErr != nil {
		return nil, s.getErr
	}
	return s.getRes, nil
}

type stubStorage struct {
	presignURL string
	presignErr error
}

func (s *stubStorage) Put(ctx context.Context, key string, contentType string, data []byte) error {
	return nil
}

func (s *stubStorage) Get(ctx context.Context, key string) (io.ReadCloser, error) {
	return nil, nil
}

func (s *stubStorage) Delete(ctx context.Context, key string) error {
	return nil
}

func (s *stubStorage) PresignGet(ctx context.Context, key string, expires time.Duration) (string, error) {
	if s.presignErr != nil {
		return "", s.presignErr
	}
	return s.presignURL, nil
}

type stubTxManager struct {
	called bool
}

func (s *stubTxManager) WithinTransaction(ctx context.Context, fn func(ctx context.Context) error) error {
	s.called = true
	return fn(ctx)
}

func TestListProducts(t *testing.T) {
	t.Run("modeが不正なときバリデーションエラーで失敗する", func(t *testing.T) {
		s := &Service{}

		_, err := s.List(context.Background(), "invalid", "all", "all")
		if err == nil || !errors.Is(err, usecase.ErrInvalidInput) {
			t.Fatalf("expected ErrInvalidInput, got %v", err)
		}
	})

	t.Run("query readerがエラーを返したとき内部エラーを返す", func(t *testing.T) {
		s := &Service{
			queryReader: &stubProductQueryReader{
				listErr: errors.New("db error"),
			},
		}

		_, err := s.List(context.Background(), ListModeAll, "all", "all")
		if err == nil || !errors.Is(err, usecase.ErrInternal) {
			t.Fatalf("expected ErrInternal, got %v", err)
		}
	})

	t.Run("presignに失敗したとき内部エラーを返す", func(t *testing.T) {
		s := &Service{
			queryReader: &stubProductQueryReader{
				listRes: []*usecaseProductQuery.Product{
					{
						UUID: "product-uuid",
						ProductImages: []usecaseProductQuery.ProductImage{
							{Path: "img/product/path.png"},
						},
					},
				},
			},
			storage: &stubStorage{presignErr: errors.New("s3 error")},
		}

		_, err := s.List(context.Background(), ListModeAll, "all", "all")
		if err == nil || !errors.Is(err, usecase.ErrInternal) {
			t.Fatalf("expected ErrInternal, got %v", err)
		}
	})

	t.Run("有効な入力を渡したとき一覧取得に成功しapiPathが設定される", func(t *testing.T) {
		s := &Service{
			queryReader: &stubProductQueryReader{
				listRes: []*usecaseProductQuery.Product{
					{
						UUID: "product-uuid",
						ProductImages: []usecaseProductQuery.ProductImage{
							{Path: "img/product/path.png"},
						},
					},
				},
			},
			storage: &stubStorage{presignURL: "https://signed.example.com/path"},
		}

		products, err := s.List(context.Background(), ListModeAll, "all", "all")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(products) != 1 || len(products[0].ProductImages) != 1 {
			t.Fatalf("unexpected products: %+v", products)
		}
		if products[0].ProductImages[0].APIPath != "https://signed.example.com/path" {
			t.Fatalf("unexpected api path: %s", products[0].ProductImages[0].APIPath)
		}
	})
}

func TestListProductsByCategory(t *testing.T) {
	t.Run("modeが不正なときバリデーションエラーで失敗する", func(t *testing.T) {
		s := &Service{}

		_, err := s.ListByCategory(context.Background(), "invalid", "all", "all")
		if err == nil || !errors.Is(err, usecase.ErrInvalidInput) {
			t.Fatalf("expected ErrInvalidInput, got %v", err)
		}
	})

	t.Run("categoryが存在しないときNotFoundエラーを返す", func(t *testing.T) {
		s := &Service{
			queryReader: &stubProductQueryReader{
				listByCatErr: usecaseProductQuery.ErrCategoryNotFound,
			},
		}

		_, err := s.ListByCategory(context.Background(), ListModeAll, id.GenerateUUID(), "all")
		if err == nil || !errors.Is(err, usecase.ErrNotFound) {
			t.Fatalf("expected ErrNotFound, got %v", err)
		}
	})

	t.Run("query readerがエラーを返したとき内部エラーを返す", func(t *testing.T) {
		s := &Service{
			queryReader: &stubProductQueryReader{
				listByCatErr: errors.New("db error"),
			},
		}

		_, err := s.ListByCategory(context.Background(), ListModeAll, "all", "all")
		if err == nil || !errors.Is(err, usecase.ErrInternal) {
			t.Fatalf("expected ErrInternal, got %v", err)
		}
	})

	t.Run("presignに失敗したとき内部エラーを返す", func(t *testing.T) {
		s := &Service{
			queryReader: &stubProductQueryReader{
				listByCatRes: []*usecaseProductQuery.CategoryProducts{
					{
						Category: usecaseProductQuery.Classification{UUID: "category-uuid", Name: "category"},
						Products: []*usecaseProductQuery.Product{
							{
								UUID: "product-uuid",
								ProductImages: []usecaseProductQuery.ProductImage{
									{Path: "img/product/path.png"},
								},
							},
						},
					},
				},
			},
			storage: &stubStorage{presignErr: errors.New("s3 error")},
		}

		_, err := s.ListByCategory(context.Background(), ListModeAll, "all", "all")
		if err == nil || !errors.Is(err, usecase.ErrInternal) {
			t.Fatalf("expected ErrInternal, got %v", err)
		}
	})

	t.Run("有効な入力を渡したときカテゴリ別一覧取得に成功し空カテゴリも維持される", func(t *testing.T) {
		s := &Service{
			queryReader: &stubProductQueryReader{
				listByCatRes: []*usecaseProductQuery.CategoryProducts{
					{
						Category: usecaseProductQuery.Classification{UUID: "category-1", Name: "Category 1"},
						Products: []*usecaseProductQuery.Product{
							{
								UUID: "product-1",
								ProductImages: []usecaseProductQuery.ProductImage{
									{Path: "img/product/path.png"},
								},
							},
						},
					},
					{
						Category: usecaseProductQuery.Classification{UUID: "category-2", Name: "Category 2"},
						Products: []*usecaseProductQuery.Product{},
					},
				},
			},
			storage: &stubStorage{presignURL: "https://signed.example.com/path"},
		}

		categoryProducts, err := s.ListByCategory(context.Background(), ListModeActive, "all", "all")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(categoryProducts) != 2 {
			t.Fatalf("expected 2 categories, got %d", len(categoryProducts))
		}
		if len(categoryProducts[0].Products) != 1 {
			t.Fatalf("expected first category products length 1, got %d", len(categoryProducts[0].Products))
		}
		if len(categoryProducts[1].Products) != 0 {
			t.Fatalf("expected second category products length 0, got %d", len(categoryProducts[1].Products))
		}
		if categoryProducts[0].Products[0].ProductImages[0].APIPath != "https://signed.example.com/path" {
			t.Fatalf("unexpected api path: %s", categoryProducts[0].Products[0].ProductImages[0].APIPath)
		}
	})
}

func TestGetProduct(t *testing.T) {
	t.Run("productUUIDが不正なときバリデーションエラーで失敗する", func(t *testing.T) {
		s := &Service{}

		_, err := s.Get(context.Background(), "invalid-uuid")
		if err == nil || !errors.Is(err, usecase.ErrInvalidInput) {
			t.Fatalf("expected ErrInvalidInput, got %v", err)
		}
	})

	t.Run("対象が見つからないときNotFoundエラーを返す", func(t *testing.T) {
		s := &Service{
			queryReader: &stubProductQueryReader{getRes: nil},
		}

		_, err := s.Get(context.Background(), id.GenerateUUID())
		if err == nil || !errors.Is(err, usecase.ErrNotFound) {
			t.Fatalf("expected ErrNotFound, got %v", err)
		}
	})

	t.Run("有効な入力を渡したとき詳細取得に成功しapiPathが設定される", func(t *testing.T) {
		s := &Service{
			queryReader: &stubProductQueryReader{
				getRes: &usecaseProductQuery.Product{
					UUID: "product-uuid",
					ProductImages: []usecaseProductQuery.ProductImage{
						{Path: "img/product/detail.png"},
					},
				},
			},
			storage: &stubStorage{presignURL: "https://signed.example.com/detail"},
		}

		product, err := s.Get(context.Background(), id.GenerateUUID())
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if product == nil || len(product.ProductImages) != 1 {
			t.Fatalf("unexpected product: %+v", product)
		}
		if product.ProductImages[0].APIPath != "https://signed.example.com/detail" {
			t.Fatalf("unexpected api path: %s", product.ProductImages[0].APIPath)
		}
	})
}

func TestCreateProductImagesWithEmptyFiles(t *testing.T) {
	t.Run("filesが空でもproductUUIDが不正ならバリデーションエラーで失敗する", func(t *testing.T) {
		repo := &stubProductRepoForCreateImages{}
		s := &Service{productRepo: repo}

		err := s.CreateProductImages(context.Background(), "invalid-uuid", []ProductImageUploadFile{}, false, nil)
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

		err := s.CreateProductImages(context.Background(), id.GenerateUUID(), []ProductImageUploadFile{}, false, nil)
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

		err = s.CreateProductImages(context.Background(), productUUID, []ProductImageUploadFile{}, false, nil)
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

		err := s.CreateProductImages(context.Background(), productUUID, []ProductImageUploadFile{
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
