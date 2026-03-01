package command

import (
	"context"
	"errors"
	"testing"

	domainCategory "github.com/tokushun109/tku/backend/internal/domain/category"
	"github.com/tokushun109/tku/backend/internal/domain/primitive"
	domainProduct "github.com/tokushun109/tku/backend/internal/domain/product"
	domainTarget "github.com/tokushun109/tku/backend/internal/domain/target"
	"github.com/tokushun109/tku/backend/internal/shared/id"
	"github.com/tokushun109/tku/backend/internal/usecase"
	usecaseProduct "github.com/tokushun109/tku/backend/internal/usecase/product"
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

func (s *stubProductRepoForCreateImages) ReplaceTags(ctx context.Context, productUUID primitive.UUID, tagUUIDs []primitive.UUID) error {
	return nil
}

type stubTxManager struct {
	called bool
}

func (s *stubTxManager) WithinTransaction(ctx context.Context, fn func(ctx context.Context) error) error {
	s.called = true
	return fn(ctx)
}

type stubProductRepoForCSV struct {
	productsByUUID map[string]*domainProduct.Product
	updated        []*domainProduct.Product
}

func (s *stubProductRepoForCSV) Create(ctx context.Context, p *domainProduct.Product) (primitive.ID, error) {
	return 0, nil
}

func (s *stubProductRepoForCSV) FindByUUID(ctx context.Context, uuid primitive.UUID) (*domainProduct.Product, error) {
	if s.productsByUUID == nil {
		return nil, nil
	}
	return s.productsByUUID[uuid.Value()], nil
}

func (s *stubProductRepoForCSV) Update(ctx context.Context, p *domainProduct.Product) (bool, error) {
	s.updated = append(s.updated, p)
	return true, nil
}

func (s *stubProductRepoForCSV) Delete(ctx context.Context, uuid primitive.UUID) (bool, error) {
	return false, nil
}

func (s *stubProductRepoForCSV) ReplaceTags(ctx context.Context, productUUID primitive.UUID, tagUUIDs []primitive.UUID) error {
	return nil
}

type stubCategoryRepoForCSV struct {
	byName map[string]*domainCategory.Category
	byUUID map[string]*domainCategory.Category
}

func (s *stubCategoryRepoForCSV) Create(ctx context.Context, c *domainCategory.Category) (*domainCategory.Category, error) {
	if s.byName == nil {
		s.byName = map[string]*domainCategory.Category{}
	}
	if s.byUUID == nil {
		s.byUUID = map[string]*domainCategory.Category{}
	}

	rebuilt, err := domainCategory.Rebuild(uint(len(s.byName)+1), c.UUID().Value(), c.Name().Value())
	if err != nil {
		return nil, err
	}
	s.byName[rebuilt.Name().Value()] = rebuilt
	s.byUUID[rebuilt.UUID().Value()] = rebuilt
	return rebuilt, nil
}

func (s *stubCategoryRepoForCSV) FindAll(ctx context.Context) ([]*domainCategory.Category, error) {
	return nil, nil
}

func (s *stubCategoryRepoForCSV) FindUsed(ctx context.Context) ([]*domainCategory.Category, error) {
	return nil, nil
}

func (s *stubCategoryRepoForCSV) FindByUUID(ctx context.Context, uuid primitive.UUID) (*domainCategory.Category, error) {
	if s.byUUID == nil {
		return nil, nil
	}
	return s.byUUID[uuid.Value()], nil
}

func (s *stubCategoryRepoForCSV) FindByName(ctx context.Context, name domainCategory.CategoryName) (*domainCategory.Category, error) {
	if s.byName == nil {
		return nil, nil
	}
	return s.byName[name.Value()], nil
}

func (s *stubCategoryRepoForCSV) ExistsByName(ctx context.Context, name domainCategory.CategoryName) (bool, error) {
	if s.byName == nil {
		return false, nil
	}
	_, ok := s.byName[name.Value()]
	return ok, nil
}

func (s *stubCategoryRepoForCSV) Update(ctx context.Context, c *domainCategory.Category) (bool, error) {
	return false, nil
}

func (s *stubCategoryRepoForCSV) Delete(ctx context.Context, uuid primitive.UUID) (bool, error) {
	return false, nil
}

type stubTargetRepoForCSV struct {
	byName map[string]*domainTarget.Target
	byUUID map[string]*domainTarget.Target
}

func (s *stubTargetRepoForCSV) Create(ctx context.Context, t *domainTarget.Target) (*domainTarget.Target, error) {
	if s.byName == nil {
		s.byName = map[string]*domainTarget.Target{}
	}
	if s.byUUID == nil {
		s.byUUID = map[string]*domainTarget.Target{}
	}

	rebuilt, err := domainTarget.Rebuild(uint(len(s.byName)+1), t.UUID().Value(), t.Name().Value())
	if err != nil {
		return nil, err
	}
	s.byName[rebuilt.Name().Value()] = rebuilt
	s.byUUID[rebuilt.UUID().Value()] = rebuilt
	return rebuilt, nil
}

func (s *stubTargetRepoForCSV) FindAll(ctx context.Context) ([]*domainTarget.Target, error) {
	return nil, nil
}

func (s *stubTargetRepoForCSV) FindUsed(ctx context.Context) ([]*domainTarget.Target, error) {
	return nil, nil
}

func (s *stubTargetRepoForCSV) FindByUUID(ctx context.Context, uuid primitive.UUID) (*domainTarget.Target, error) {
	if s.byUUID == nil {
		return nil, nil
	}
	return s.byUUID[uuid.Value()], nil
}

func (s *stubTargetRepoForCSV) FindByName(ctx context.Context, name domainTarget.TargetName) (*domainTarget.Target, error) {
	if s.byName == nil {
		return nil, nil
	}
	return s.byName[name.Value()], nil
}

func (s *stubTargetRepoForCSV) ExistsByName(ctx context.Context, name domainTarget.TargetName) (bool, error) {
	if s.byName == nil {
		return false, nil
	}
	_, ok := s.byName[name.Value()]
	return ok, nil
}

func (s *stubTargetRepoForCSV) Update(ctx context.Context, t *domainTarget.Target) (bool, error) {
	return false, nil
}

func (s *stubTargetRepoForCSV) Delete(ctx context.Context, uuid primitive.UUID) (bool, error) {
	return false, nil
}

type stubUUIDGenForCSV struct{}

func (g *stubUUIDGenForCSV) New() string {
	return id.GenerateUUID()
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

func TestUploadCSV(t *testing.T) {
	t.Run("有効なCSV行を渡したとき商品更新と分類補完に成功する", func(t *testing.T) {
		product, err := domainProduct.Rebuild(
			1,
			id.GenerateUUID(),
			"old product",
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

		productRepo := &stubProductRepoForCSV{
			productsByUUID: map[string]*domainProduct.Product{
				product.UUID().Value(): product,
			},
		}
		categoryRepo := &stubCategoryRepoForCSV{}
		targetRepo := &stubTargetRepoForCSV{}
		tx := &stubTxManager{}

		s := &Service{
			productRepo:  productRepo,
			categoryRepo: categoryRepo,
			targetRepo:   targetRepo,
			uuidGen:      &stubUUIDGenForCSV{},
			txManager:    tx,
		}

		err = s.UploadCSV(context.Background(), []usecaseProduct.ProductCSVInputRow{
			{
				UUID:         product.UUID().Value(),
				Name:         "new product",
				Price:        2500,
				CategoryName: "new category",
				TargetName:   "new target",
			},
		})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if !tx.called {
			t.Fatalf("expected transaction called")
		}
		if product.Name().Value() != "new product" {
			t.Fatalf("unexpected product name: %s", product.Name().Value())
		}
		if product.Price().Value() != 2500 {
			t.Fatalf("unexpected product price: %d", product.Price().Value())
		}
		if product.CategoryUUID() == nil {
			t.Fatalf("expected category uuid assigned")
		}
		if product.TargetUUID() == nil {
			t.Fatalf("expected target uuid assigned")
		}
		if len(productRepo.updated) != 1 {
			t.Fatalf("expected one update call, got %d", len(productRepo.updated))
		}
		if len(categoryRepo.byName) != 1 {
			t.Fatalf("expected one category created, got %d", len(categoryRepo.byName))
		}
		if len(targetRepo.byName) != 1 {
			t.Fatalf("expected one target created, got %d", len(targetRepo.byName))
		}
	})

	t.Run("行データが不正なときバリデーションエラーで失敗する", func(t *testing.T) {
		tx := &stubTxManager{}
		s := &Service{
			txManager: tx,
		}

		err := s.UploadCSV(context.Background(), []usecaseProduct.ProductCSVInputRow{
			{
				UUID:         "",
				Name:         "product",
				Price:        1000,
				CategoryName: "category",
				TargetName:   "target",
			},
		})
		if err == nil || !errors.Is(err, usecase.ErrInvalidInput) {
			t.Fatalf("expected ErrInvalidInput, got %v", err)
		}
		if tx.called {
			t.Fatalf("transaction should not be called")
		}
	})

	t.Run("分類名が空欄のとき分類IDをNULLに更新する", func(t *testing.T) {
		currentCategoryUUID := "22222222-2222-4222-8222-222222222222"
		currentTargetUUID := "33333333-3333-4333-8333-333333333333"
		product, err := domainProduct.Rebuild(
			1,
			id.GenerateUUID(),
			"old product",
			"",
			1000,
			true,
			false,
			&currentCategoryUUID,
			&currentTargetUUID,
		)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		productRepo := &stubProductRepoForCSV{
			productsByUUID: map[string]*domainProduct.Product{
				product.UUID().Value(): product,
			},
		}
		categoryRepo := &stubCategoryRepoForCSV{}
		targetRepo := &stubTargetRepoForCSV{}
		tx := &stubTxManager{}

		s := &Service{
			productRepo:  productRepo,
			categoryRepo: categoryRepo,
			targetRepo:   targetRepo,
			uuidGen:      &stubUUIDGenForCSV{},
			txManager:    tx,
		}

		err = s.UploadCSV(context.Background(), []usecaseProduct.ProductCSVInputRow{
			{
				UUID:         product.UUID().Value(),
				Name:         "new product",
				Price:        2500,
				CategoryName: "",
				TargetName:   "",
			},
		})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if !tx.called {
			t.Fatalf("expected transaction called")
		}
		if product.CategoryUUID() != nil {
			t.Fatalf("expected category uuid cleared")
		}
		if product.TargetUUID() != nil {
			t.Fatalf("expected target uuid cleared")
		}
		if len(productRepo.updated) != 1 {
			t.Fatalf("expected one update call, got %d", len(productRepo.updated))
		}
		if len(categoryRepo.byName) != 0 {
			t.Fatalf("expected category not created")
		}
		if len(targetRepo.byName) != 0 {
			t.Fatalf("expected target not created")
		}
	})

	t.Run("対象の商品IDが存在しないときバリデーションエラーで失敗する", func(t *testing.T) {
		tx := &stubTxManager{}
		s := &Service{
			productRepo:  &stubProductRepoForCSV{},
			categoryRepo: &stubCategoryRepoForCSV{},
			targetRepo:   &stubTargetRepoForCSV{},
			uuidGen:      &stubUUIDGenForCSV{},
			txManager:    tx,
		}

		err := s.UploadCSV(context.Background(), []usecaseProduct.ProductCSVInputRow{
			{
				UUID:         id.GenerateUUID(),
				Name:         "product",
				Price:        1000,
				CategoryName: "category",
				TargetName:   "target",
			},
		})
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
