package command

import (
	"context"
	"errors"
	"io"
	"strings"
	"testing"
	"time"

	"github.com/tokushun109/tku/backend/internal/domain/primitive"
	domainProduct "github.com/tokushun109/tku/backend/internal/domain/product"
	domainSalesSite "github.com/tokushun109/tku/backend/internal/domain/sales_site"
	domainSiteDetail "github.com/tokushun109/tku/backend/internal/domain/site_detail"
	domainTag "github.com/tokushun109/tku/backend/internal/domain/tag"
	"github.com/tokushun109/tku/backend/internal/shared/id"
	"github.com/tokushun109/tku/backend/internal/usecase"
	usecaseProduct "github.com/tokushun109/tku/backend/internal/usecase/product"
)

type stubDuplicateSource struct {
	result *usecaseProduct.DuplicateProductData
	err    error
	called bool
	rawURL string
}

func (s *stubDuplicateSource) Duplicate(ctx context.Context, rawURL string) (*usecaseProduct.DuplicateProductData, error) {
	s.called = true
	s.rawURL = rawURL
	if s.err != nil {
		return nil, s.err
	}
	return s.result, nil
}

type stubUUIDGenerator struct{}

func (g *stubUUIDGenerator) New() string {
	return id.GenerateUUID()
}

type stubProductRepoForDuplicate struct {
	created         *domainProduct.Product
	createErr       error
	createdID       primitive.ID
	replacedTagIDs  []primitive.ID
	replacedProduct primitive.ID
}

func (s *stubProductRepoForDuplicate) Create(ctx context.Context, p *domainProduct.Product) (primitive.ID, error) {
	s.created = p
	if s.createErr != nil {
		return 0, s.createErr
	}
	if s.createdID == 0 {
		s.createdID = primitive.ID(10)
	}
	return s.createdID, nil
}

func (s *stubProductRepoForDuplicate) FindByUUID(ctx context.Context, uuid primitive.UUID) (*domainProduct.Product, error) {
	return nil, nil
}

func (s *stubProductRepoForDuplicate) FindByID(ctx context.Context, id primitive.ID) (*domainProduct.Product, error) {
	return nil, nil
}

func (s *stubProductRepoForDuplicate) Update(ctx context.Context, p *domainProduct.Product) (bool, error) {
	return false, nil
}

func (s *stubProductRepoForDuplicate) Delete(ctx context.Context, uuid primitive.UUID) (bool, error) {
	return false, nil
}

func (s *stubProductRepoForDuplicate) ReplaceTags(ctx context.Context, productID primitive.ID, tagIDs []primitive.ID) error {
	s.replacedProduct = productID
	s.replacedTagIDs = tagIDs
	return nil
}

type stubTagRepoForDuplicate struct {
	byName  map[string]*domainTag.Tag
	created []*domainTag.Tag
	findErr error
}

func (s *stubTagRepoForDuplicate) Create(ctx context.Context, tag *domainTag.Tag) (*domainTag.Tag, error) {
	if s.byName == nil {
		s.byName = map[string]*domainTag.Tag{}
	}

	rebuilt, err := domainTag.Rebuild(uint(len(s.byName)+1), tag.UUID().Value(), tag.Name().Value())
	if err != nil {
		return nil, err
	}
	s.byName[strings.ToLower(rebuilt.Name().Value())] = rebuilt
	s.created = append(s.created, rebuilt)
	return rebuilt, nil
}

func (s *stubTagRepoForDuplicate) FindAll(ctx context.Context) ([]*domainTag.Tag, error) {
	return nil, nil
}

func (s *stubTagRepoForDuplicate) FindByName(ctx context.Context, name domainTag.TagName) (*domainTag.Tag, error) {
	if s.findErr != nil {
		return nil, s.findErr
	}
	if s.byName == nil {
		return nil, nil
	}
	return s.byName[strings.ToLower(name.Value())], nil
}

func (s *stubTagRepoForDuplicate) FindByUUID(ctx context.Context, uuid primitive.UUID) (*domainTag.Tag, error) {
	if s.findErr != nil {
		return nil, s.findErr
	}
	for _, tag := range s.byName {
		if tag != nil && tag.UUID() == uuid {
			return tag, nil
		}
	}
	return nil, nil
}

func (s *stubTagRepoForDuplicate) ExistsByName(ctx context.Context, name domainTag.TagName) (bool, error) {
	if s.byName == nil {
		return false, nil
	}
	_, ok := s.byName[strings.ToLower(name.Value())]
	return ok, nil
}

func (s *stubTagRepoForDuplicate) Update(ctx context.Context, tag *domainTag.Tag) (bool, error) {
	return false, nil
}

func (s *stubTagRepoForDuplicate) Delete(ctx context.Context, uuid primitive.UUID) (bool, error) {
	return false, nil
}

type stubSalesSiteRepoForDuplicate struct {
	byName map[string]*domainSalesSite.SalesSite
	err    error
}

func (s *stubSalesSiteRepoForDuplicate) Create(ctx context.Context, salesSite *domainSalesSite.SalesSite) (*domainSalesSite.SalesSite, error) {
	return salesSite, nil
}

func (s *stubSalesSiteRepoForDuplicate) FindAll(ctx context.Context) ([]*domainSalesSite.SalesSite, error) {
	return nil, nil
}

func (s *stubSalesSiteRepoForDuplicate) FindByName(ctx context.Context, name domainSalesSite.SalesSiteName) (*domainSalesSite.SalesSite, error) {
	if s.err != nil {
		return nil, s.err
	}
	if s.byName == nil {
		return nil, nil
	}
	return s.byName[strings.ToLower(name.Value())], nil
}

func (s *stubSalesSiteRepoForDuplicate) FindByUUID(ctx context.Context, uuid primitive.UUID) (*domainSalesSite.SalesSite, error) {
	return nil, nil
}

func (s *stubSalesSiteRepoForDuplicate) Update(ctx context.Context, salesSite *domainSalesSite.SalesSite) (bool, error) {
	return false, nil
}

func (s *stubSalesSiteRepoForDuplicate) Delete(ctx context.Context, uuid primitive.UUID) (bool, error) {
	return false, nil
}

type stubSiteDetailRepoForDuplicate struct {
	replacedProductID primitive.ID
	replacedDetails   []*domainSiteDetail.SiteDetail
}

func (s *stubSiteDetailRepoForDuplicate) ReplaceByProductID(ctx context.Context, productID primitive.ID, details []*domainSiteDetail.SiteDetail) error {
	s.replacedProductID = productID
	s.replacedDetails = details
	return nil
}

func (s *stubSiteDetailRepoForDuplicate) DeleteByProductID(ctx context.Context, productID primitive.ID) error {
	return nil
}

type stubProductImageRepoForDuplicate struct {
	created []*domainProduct.ProductImage
}

func (s *stubProductImageRepoForDuplicate) Create(ctx context.Context, image *domainProduct.ProductImage) (*domainProduct.ProductImage, error) {
	s.created = append(s.created, image)
	return image, nil
}

func (s *stubProductImageRepoForDuplicate) FindByUUID(ctx context.Context, uuid primitive.UUID) (*domainProduct.ProductImage, error) {
	return nil, nil
}

func (s *stubProductImageRepoForDuplicate) FindByProductID(ctx context.Context, productID primitive.ID) ([]*domainProduct.ProductImage, error) {
	return nil, nil
}

func (s *stubProductImageRepoForDuplicate) UpdateOrder(ctx context.Context, uuid primitive.UUID, order int) (bool, error) {
	return false, nil
}

func (s *stubProductImageRepoForDuplicate) DeleteByUUID(ctx context.Context, uuid primitive.UUID) (bool, error) {
	return false, nil
}

func (s *stubProductImageRepoForDuplicate) DeleteByProductID(ctx context.Context, productID primitive.ID) error {
	return nil
}

type stubStorageForDuplicate struct {
	puts      []string
	deletes   []string
	putErrAt  int
	putCalled int
}

func (s *stubStorageForDuplicate) Put(ctx context.Context, key string, contentType string, data []byte) error {
	s.putCalled++
	if s.putErrAt > 0 && s.putCalled == s.putErrAt {
		return errors.New("put failed")
	}
	s.puts = append(s.puts, key)
	return nil
}

func (s *stubStorageForDuplicate) Get(ctx context.Context, key string) (io.ReadCloser, error) {
	return io.NopCloser(strings.NewReader("")), nil
}

func (s *stubStorageForDuplicate) Delete(ctx context.Context, key string) error {
	s.deletes = append(s.deletes, key)
	return nil
}

func (s *stubStorageForDuplicate) PresignGet(ctx context.Context, key string, expires time.Duration) (string, error) {
	return "", nil
}

func TestDuplicateProduct(t *testing.T) {
	t.Run("URLが不正なときバリデーションエラーで失敗する", func(t *testing.T) {
		source := &stubDuplicateSource{}
		s := &Service{
			duplicateSource: source,
		}

		err := s.Duplicate(context.Background(), "example.com/items/1")
		if err == nil || !errors.Is(err, usecase.ErrInvalidInput) {
			t.Fatalf("expected ErrInvalidInput, got %v", err)
		}
		if source.called {
			t.Fatalf("duplicate source should not be called")
		}
	})

	t.Run("duplicate sourceが入力エラーを返したときバリデーションエラーで失敗する", func(t *testing.T) {
		source := &stubDuplicateSource{err: usecase.ErrInvalidInput}
		s := &Service{
			duplicateSource: source,
		}

		err := s.Duplicate(context.Background(), "https://www.creema.jp/items/1")
		if err == nil || !errors.Is(err, usecase.ErrInvalidInput) {
			t.Fatalf("expected ErrInvalidInput, got %v", err)
		}
		if !source.called {
			t.Fatalf("duplicate source should be called")
		}
	})

	t.Run("duplicate sourceがnot foundを返したときnot foundで失敗する", func(t *testing.T) {
		source := &stubDuplicateSource{err: usecase.ErrNotFound}
		s := &Service{
			duplicateSource: source,
		}

		err := s.Duplicate(context.Background(), "https://www.creema.jp/items/1")
		if err == nil || !errors.Is(err, usecase.ErrNotFound) {
			t.Fatalf("expected ErrNotFound, got %v", err)
		}
		if !source.called {
			t.Fatalf("duplicate source should be called")
		}
	})

	t.Run("有効な入力を渡したとき商品と関連データの複製に成功する", func(t *testing.T) {
		source := &stubDuplicateSource{
			result: &usecaseProduct.DuplicateProductData{
				Name:        "sample product",
				Description: "desc",
				Price:       1200,
				Tags:        []string{"Silver", "silver", "Ring"},
				Images: []usecaseProduct.DuplicateProductImage{
					{Name: "main.png", Data: []byte{0x89, 'P', 'N', 'G', '\r', '\n', 0x1a, '\n'}},
					{Name: "sub.png", Data: []byte{0x89, 'P', 'N', 'G', '\r', '\n', 0x1a, '\n'}},
				},
			},
		}
		productRepo := &stubProductRepoForDuplicate{}
		tagRepo := &stubTagRepoForDuplicate{}
		salesSiteRepo := &stubSalesSiteRepoForDuplicate{
			byName: map[string]*domainSalesSite.SalesSite{
				"creema": mustDuplicateSalesSite(t, "11111111-1111-4111-8111-111111111111"),
			},
		}
		siteDetailRepo := &stubSiteDetailRepoForDuplicate{}
		productImageRepo := &stubProductImageRepoForDuplicate{}
		storage := &stubStorageForDuplicate{}
		tx := &stubTxManager{}

		s := &Service{
			productRepo:      productRepo,
			productImageRepo: productImageRepo,
			siteDetailRepo:   siteDetailRepo,
			tagRepo:          tagRepo,
			salesSiteRepo:    salesSiteRepo,
			duplicateSource:  source,
			storage:          storage,
			uuidGen:          &stubUUIDGenerator{},
			txManager:        tx,
		}

		err := s.Duplicate(context.Background(), "https://www.creema.jp/items/123")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if !source.called {
			t.Fatalf("duplicate source should be called")
		}
		if !tx.called {
			t.Fatalf("transaction should be called")
		}
		if productRepo.created == nil {
			t.Fatalf("product should be created")
		}
		if productRepo.created.IsActive() {
			t.Fatalf("duplicated product should be inactive")
		}
		if productRepo.created.IsRecommend() {
			t.Fatalf("duplicated product should not be recommend")
		}
		if len(tagRepo.created) != 2 {
			t.Fatalf("expected 2 created tags, got %d", len(tagRepo.created))
		}
		if len(productRepo.replacedTagIDs) != 2 {
			t.Fatalf("expected 2 tag ids, got %d", len(productRepo.replacedTagIDs))
		}
		if len(siteDetailRepo.replacedDetails) != 1 {
			t.Fatalf("expected 1 site detail, got %d", len(siteDetailRepo.replacedDetails))
		}
		if siteDetailRepo.replacedDetails[0].DetailURL().Value() != "https://www.creema.jp/items/123" {
			t.Fatalf("unexpected detail url: %s", siteDetailRepo.replacedDetails[0].DetailURL().Value())
		}
		if len(productImageRepo.created) != 2 {
			t.Fatalf("expected 2 product images, got %d", len(productImageRepo.created))
		}
		if productImageRepo.created[0].Order().Value() != 2 || productImageRepo.created[1].Order().Value() != 1 {
			t.Fatalf("unexpected image order: %d, %d", productImageRepo.created[0].Order().Value(), productImageRepo.created[1].Order().Value())
		}
		if len(storage.puts) != 2 {
			t.Fatalf("expected 2 storage puts, got %d", len(storage.puts))
		}
	})

	t.Run("画像保存に失敗したときロールバック用の削除を試行する", func(t *testing.T) {
		source := &stubDuplicateSource{
			result: &usecaseProduct.DuplicateProductData{
				Name:        "sample product",
				Description: "desc",
				Price:       1200,
				Images: []usecaseProduct.DuplicateProductImage{
					{Name: "main.png", Data: []byte{0x89, 'P', 'N', 'G', '\r', '\n', 0x1a, '\n'}},
					{Name: "sub.png", Data: []byte{0x89, 'P', 'N', 'G', '\r', '\n', 0x1a, '\n'}},
				},
			},
		}
		storage := &stubStorageForDuplicate{putErrAt: 2}
		tx := &stubTxManager{}

		s := &Service{
			productRepo:      &stubProductRepoForDuplicate{},
			productImageRepo: &stubProductImageRepoForDuplicate{},
			siteDetailRepo:   &stubSiteDetailRepoForDuplicate{},
			tagRepo:          &stubTagRepoForDuplicate{},
			salesSiteRepo: &stubSalesSiteRepoForDuplicate{
				byName: map[string]*domainSalesSite.SalesSite{
					"creema": mustDuplicateSalesSite(t, "11111111-1111-4111-8111-111111111111"),
				},
			},
			duplicateSource: source,
			storage:         storage,
			uuidGen:         &stubUUIDGenerator{},
			txManager:       tx,
		}

		err := s.Duplicate(context.Background(), "https://www.creema.jp/items/123")
		if err == nil || !errors.Is(err, usecase.ErrInternal) {
			t.Fatalf("expected ErrInternal, got %v", err)
		}
		if !tx.called {
			t.Fatalf("transaction should be called")
		}
		if len(storage.deletes) != 1 {
			t.Fatalf("expected 1 rollback delete, got %d", len(storage.deletes))
		}
	})
}

func mustDuplicateSalesSite(t *testing.T, uuid string) *domainSalesSite.SalesSite {
	t.Helper()

	salesSite, err := domainSalesSite.Rebuild(1, uuid, "Creema", "https://www.creema.jp", "")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	return salesSite
}
