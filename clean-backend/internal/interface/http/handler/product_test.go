package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	usecaseProduct "github.com/tokushun109/tku/clean-backend/internal/usecase/product"
	usecaseProductQuery "github.com/tokushun109/tku/clean-backend/internal/usecase/product/query"
)

type stubProductUC struct {
	listByCategoryErr    error
	listByCategoryRes    []*usecaseProductQuery.CategoryProducts
	listByCategoryCalled bool
	listByCategoryReq    struct {
		mode     string
		category string
		target   string
	}

	createProductImagesErr    error
	createProductImagesCalled bool
	createProductImagesReq    struct {
		productUUID string
		files       []usecaseProduct.ProductImageUploadFile
		isChanged   bool
		orderMap    map[int]int
	}
}

func (s *stubProductUC) List(ctx context.Context, mode string, category string, target string) ([]*usecaseProductQuery.Product, error) {
	return nil, nil
}

func (s *stubProductUC) ListByCategory(
	ctx context.Context,
	mode string,
	category string,
	target string,
) ([]*usecaseProductQuery.CategoryProducts, error) {
	s.listByCategoryCalled = true
	s.listByCategoryReq.mode = mode
	s.listByCategoryReq.category = category
	s.listByCategoryReq.target = target
	if s.listByCategoryErr != nil {
		return nil, s.listByCategoryErr
	}
	return s.listByCategoryRes, nil
}

func (s *stubProductUC) Get(ctx context.Context, productUUID string) (*usecaseProductQuery.Product, error) {
	return nil, nil
}

func (s *stubProductUC) Create(ctx context.Context, input usecaseProduct.CreateProductInput) (*usecaseProductQuery.Product, error) {
	return nil, nil
}

func (s *stubProductUC) Update(ctx context.Context, productUUID string, input usecaseProduct.UpdateProductInput) error {
	return nil
}

func (s *stubProductUC) Delete(ctx context.Context, productUUID string) error {
	return nil
}

func (s *stubProductUC) GetProductImageBlob(ctx context.Context, productImageUUID string) (*usecaseProduct.ProductImageBlob, error) {
	return nil, nil
}

func (s *stubProductUC) CreateProductImages(
	ctx context.Context,
	productUUID string,
	files []usecaseProduct.ProductImageUploadFile,
	isChanged bool,
	orderMap map[int]int,
) error {
	s.createProductImagesCalled = true
	s.createProductImagesReq.productUUID = productUUID
	s.createProductImagesReq.files = files
	s.createProductImagesReq.isChanged = isChanged
	s.createProductImagesReq.orderMap = orderMap
	return s.createProductImagesErr
}

func (s *stubProductUC) DeleteProductImage(ctx context.Context, productUUID string, productImageUUID string) error {
	return nil
}

func TestProductListByCategory(t *testing.T) {
	t.Run("クエリが不正なときバリデーションエラーで失敗する", func(t *testing.T) {
		uc := &stubProductUC{}
		h := NewProductHandler(uc)

		req := httptest.NewRequest(http.MethodGet, "/api/category/product?mode=invalid&category=all&target=all", nil)
		rr := httptest.NewRecorder()

		h.ListByCategory(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Fatalf("expected 400, got %d", rr.Code)
		}
		if uc.listByCategoryCalled {
			t.Fatalf("usecase should not be called on invalid query")
		}
	})

	t.Run("有効なクエリを渡したときカテゴリ別一覧を返す", func(t *testing.T) {
		uc := &stubProductUC{
			listByCategoryRes: []*usecaseProductQuery.CategoryProducts{
				{
					Category: usecaseProductQuery.Classification{
						UUID: "category-uuid",
						Name: "Category",
					},
					Products: []*usecaseProductQuery.Product{},
				},
			},
		}
		h := NewProductHandler(uc)

		req := httptest.NewRequest(http.MethodGet, "/api/category/product?mode=active&category=all&target=all", nil)
		rr := httptest.NewRecorder()

		h.ListByCategory(rr, req)

		if rr.Code != http.StatusOK {
			t.Fatalf("expected 200, got %d", rr.Code)
		}
		if !uc.listByCategoryCalled {
			t.Fatalf("usecase should be called")
		}
		if uc.listByCategoryReq.mode != "active" || uc.listByCategoryReq.category != "all" || uc.listByCategoryReq.target != "all" {
			t.Fatalf("unexpected query args: %+v", uc.listByCategoryReq)
		}

		var res []map[string]any
		if err := json.NewDecoder(rr.Body).Decode(&res); err != nil {
			t.Fatalf("unexpected decode error: %v", err)
		}
		if len(res) != 1 {
			t.Fatalf("expected 1 category, got %d", len(res))
		}
		category, ok := res[0]["category"].(map[string]any)
		if !ok {
			t.Fatalf("expected category object")
		}
		if category["uuid"] != "category-uuid" {
			t.Fatalf("unexpected category uuid: %v", category["uuid"])
		}
	})
}
