package query

import (
	"context"
	"errors"
	"io"
	"testing"
	"time"

	"github.com/tokushun109/tku/clean-backend/internal/shared/id"
	"github.com/tokushun109/tku/clean-backend/internal/usecase"
)

type stubQueryReader struct {
	listRes            []*Product
	listErr            error
	listByCatRes       []*CategoryProducts
	listByCatErr       error
	listCarouselRes    []*CarouselItem
	listCarouselErr    error
	listCarouselCalled bool
	listCarouselQuery  ListCarouselQuery
	getRes             *Product
	getErr             error
}

func (s *stubQueryReader) ListProducts(ctx context.Context, q ListProductsQuery) ([]*Product, error) {
	if s.listErr != nil {
		return nil, s.listErr
	}
	return s.listRes, nil
}

func (s *stubQueryReader) ListCategoryProducts(ctx context.Context, q ListCategoryProductsQuery) ([]*CategoryProducts, error) {
	if s.listByCatErr != nil {
		return nil, s.listByCatErr
	}
	return s.listByCatRes, nil
}

func (s *stubQueryReader) ListCarouselItems(ctx context.Context, q ListCarouselQuery) ([]*CarouselItem, error) {
	s.listCarouselCalled = true
	s.listCarouselQuery = q
	if s.listCarouselErr != nil {
		return nil, s.listCarouselErr
	}
	return s.listCarouselRes, nil
}

func (s *stubQueryReader) GetProductByUUID(ctx context.Context, productUUID string) (*Product, error) {
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
			queryReader: &stubQueryReader{
				listErr: errors.New("db error"),
			},
		}

		_, err := s.List(context.Background(), "all", "all", "all")
		if err == nil || !errors.Is(err, usecase.ErrInternal) {
			t.Fatalf("expected ErrInternal, got %v", err)
		}
	})

	t.Run("presignに失敗したとき内部エラーを返す", func(t *testing.T) {
		s := &Service{
			queryReader: &stubQueryReader{
				listRes: []*Product{
					{
						UUID: "product-uuid",
						ProductImages: []ProductImage{
							{Path: "img/product/path.png"},
						},
					},
				},
			},
			storage: &stubStorage{presignErr: errors.New("s3 error")},
		}

		_, err := s.List(context.Background(), "all", "all", "all")
		if err == nil || !errors.Is(err, usecase.ErrInternal) {
			t.Fatalf("expected ErrInternal, got %v", err)
		}
	})

	t.Run("有効な入力を渡したとき一覧取得に成功しapiPathが設定される", func(t *testing.T) {
		s := &Service{
			queryReader: &stubQueryReader{
				listRes: []*Product{
					{
						UUID: "product-uuid",
						ProductImages: []ProductImage{
							{Path: "img/product/path.png"},
						},
					},
				},
			},
			storage: &stubStorage{presignURL: "https://signed.example.com/path"},
		}

		products, err := s.List(context.Background(), "all", "all", "all")
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

func TestListCarousel(t *testing.T) {
	t.Run("query readerがエラーを返したとき内部エラーを返す", func(t *testing.T) {
		s := &Service{
			queryReader: &stubQueryReader{
				listCarouselErr: errors.New("db error"),
			},
		}

		_, err := s.ListCarousel(context.Background())
		if err == nil || !errors.Is(err, usecase.ErrInternal) {
			t.Fatalf("expected ErrInternal, got %v", err)
		}
	})

	t.Run("presignに失敗したとき内部エラーを返す", func(t *testing.T) {
		s := &Service{
			queryReader: &stubQueryReader{
				listCarouselRes: []*CarouselItem{
					{
						Product: &Product{
							UUID: "product-uuid",
							ProductImages: []ProductImage{
								{Path: "img/product/path.png"},
							},
						},
					},
				},
			},
			storage: &stubStorage{presignErr: errors.New("s3 error")},
		}

		_, err := s.ListCarousel(context.Background())
		if err == nil || !errors.Is(err, usecase.ErrInternal) {
			t.Fatalf("expected ErrInternal, got %v", err)
		}
	})

	t.Run("有効な入力を渡したときカルーセル一覧取得に成功しapiPathが設定される", func(t *testing.T) {
		reader := &stubQueryReader{
			listCarouselRes: []*CarouselItem{
				{
					Product: &Product{
						UUID: "product-1",
						ProductImages: []ProductImage{
							{Path: "img/product/first.png"},
						},
					},
				},
				{
					Product: &Product{
						UUID:          "product-2",
						ProductImages: []ProductImage{},
					},
				},
			},
		}
		s := &Service{
			queryReader: reader,
			storage:     &stubStorage{presignURL: "https://signed.example.com/path"},
		}

		items, err := s.ListCarousel(context.Background())
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if !reader.listCarouselCalled {
			t.Fatalf("expected ListCarouselItems called")
		}
		if reader.listCarouselQuery.Limit != defaultCarouselLimit {
			t.Fatalf("unexpected carousel limit: %d", reader.listCarouselQuery.Limit)
		}
		if len(items) != 1 {
			t.Fatalf("expected 1 item, got %d", len(items))
		}
		if items[0].APIPath != "https://signed.example.com/path" {
			t.Fatalf("unexpected api path: %s", items[0].APIPath)
		}
		if len(items[0].Product.ProductImages) != 1 || items[0].Product.ProductImages[0].APIPath != "https://signed.example.com/path" {
			t.Fatalf("unexpected product image api path: %+v", items[0].Product.ProductImages)
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
			queryReader: &stubQueryReader{
				listByCatErr: ErrCategoryNotFound,
			},
		}

		_, err := s.ListByCategory(context.Background(), "all", id.GenerateUUID(), "all")
		if err == nil || !errors.Is(err, usecase.ErrNotFound) {
			t.Fatalf("expected ErrNotFound, got %v", err)
		}
	})

	t.Run("query readerがエラーを返したとき内部エラーを返す", func(t *testing.T) {
		s := &Service{
			queryReader: &stubQueryReader{
				listByCatErr: errors.New("db error"),
			},
		}

		_, err := s.ListByCategory(context.Background(), "all", "all", "all")
		if err == nil || !errors.Is(err, usecase.ErrInternal) {
			t.Fatalf("expected ErrInternal, got %v", err)
		}
	})

	t.Run("presignに失敗したとき内部エラーを返す", func(t *testing.T) {
		s := &Service{
			queryReader: &stubQueryReader{
				listByCatRes: []*CategoryProducts{
					{
						Category: Classification{UUID: "category-uuid", Name: "category"},
						Products: []*Product{
							{
								UUID: "product-uuid",
								ProductImages: []ProductImage{
									{Path: "img/product/path.png"},
								},
							},
						},
					},
				},
			},
			storage: &stubStorage{presignErr: errors.New("s3 error")},
		}

		_, err := s.ListByCategory(context.Background(), "all", "all", "all")
		if err == nil || !errors.Is(err, usecase.ErrInternal) {
			t.Fatalf("expected ErrInternal, got %v", err)
		}
	})

	t.Run("有効な入力を渡したときカテゴリ別一覧取得に成功し空カテゴリも維持される", func(t *testing.T) {
		s := &Service{
			queryReader: &stubQueryReader{
				listByCatRes: []*CategoryProducts{
					{
						Category: Classification{UUID: "category-1", Name: "Category 1"},
						Products: []*Product{
							{
								UUID: "product-1",
								ProductImages: []ProductImage{
									{Path: "img/product/path.png"},
								},
							},
						},
					},
					{
						Category: Classification{UUID: "category-2", Name: "Category 2"},
						Products: []*Product{},
					},
				},
			},
			storage: &stubStorage{presignURL: "https://signed.example.com/path"},
		}

		categoryProducts, err := s.ListByCategory(context.Background(), "active", "all", "all")
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
			queryReader: &stubQueryReader{getRes: nil},
		}

		_, err := s.Get(context.Background(), id.GenerateUUID())
		if err == nil || !errors.Is(err, usecase.ErrNotFound) {
			t.Fatalf("expected ErrNotFound, got %v", err)
		}
	})

	t.Run("有効な入力を渡したとき詳細取得に成功しapiPathが設定される", func(t *testing.T) {
		s := &Service{
			queryReader: &stubQueryReader{
				getRes: &Product{
					UUID: "product-uuid",
					ProductImages: []ProductImage{
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
