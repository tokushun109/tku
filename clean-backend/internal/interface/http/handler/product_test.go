package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/tokushun109/tku/clean-backend/internal/domain/primitive"
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

	listCarouselErr    error
	listCarouselRes    []*usecaseProductQuery.CarouselItem
	listCarouselCalled bool

	createErr       error
	createRes       primitive.UUID
	createCalled    bool
	duplicateErr    error
	duplicateReq    string
	duplicateCalled bool

	exportCSVErr error
	exportCSVRes []*usecaseProductQuery.ProductCSVRow

	uploadCSVErr    error
	uploadCSVCalled bool
	uploadCSVRows   []usecaseProduct.ProductCSVInputRow

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

func (s *stubProductUC) ListCarousel(ctx context.Context) ([]*usecaseProductQuery.CarouselItem, error) {
	s.listCarouselCalled = true
	if s.listCarouselErr != nil {
		return nil, s.listCarouselErr
	}
	return s.listCarouselRes, nil
}

func (s *stubProductUC) Get(ctx context.Context, productUUID string) (*usecaseProductQuery.Product, error) {
	return nil, nil
}

func (s *stubProductUC) Create(ctx context.Context, input usecaseProduct.CreateProductInput) (primitive.UUID, error) {
	s.createCalled = true
	if s.createErr != nil {
		return "", s.createErr
	}
	return s.createRes, nil
}

func (s *stubProductUC) Duplicate(ctx context.Context, rawURL string) error {
	s.duplicateCalled = true
	s.duplicateReq = rawURL
	return s.duplicateErr
}

func (s *stubProductUC) Update(ctx context.Context, productUUID string, input usecaseProduct.UpdateProductInput) error {
	return nil
}

func (s *stubProductUC) Delete(ctx context.Context, productUUID string) error {
	return nil
}

func (s *stubProductUC) ExportCSV(ctx context.Context) ([]*usecaseProductQuery.ProductCSVRow, error) {
	if s.exportCSVErr != nil {
		return nil, s.exportCSVErr
	}
	return s.exportCSVRes, nil
}

func (s *stubProductUC) UploadCSV(ctx context.Context, rows []usecaseProduct.ProductCSVInputRow) error {
	s.uploadCSVCalled = true
	s.uploadCSVRows = rows
	return s.uploadCSVErr
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

func TestProductListCarousel(t *testing.T) {
	t.Run("ユースケースがエラーを返したとき内部エラーで失敗する", func(t *testing.T) {
		uc := &stubProductUC{listCarouselErr: context.DeadlineExceeded}
		h := NewProductHandler(uc)

		req := httptest.NewRequest(http.MethodGet, "/api/carousel_image", nil)
		rr := httptest.NewRecorder()

		h.ListCarousel(rr, req)

		if rr.Code != http.StatusInternalServerError {
			t.Fatalf("expected 500, got %d", rr.Code)
		}
		if !uc.listCarouselCalled {
			t.Fatalf("usecase should be called")
		}
	})

	t.Run("有効な入力を渡したときカルーセル一覧を返す", func(t *testing.T) {
		uc := &stubProductUC{
			listCarouselRes: []*usecaseProductQuery.CarouselItem{
				{
					APIPath: "https://signed.example.com/carousel.png",
					Product: &usecaseProductQuery.Product{
						UUID: "product-uuid",
						Name: "product name",
						Target: usecaseProductQuery.Classification{
							UUID: "target-uuid",
							Name: "target",
						},
					},
				},
			},
		}
		h := NewProductHandler(uc)

		req := httptest.NewRequest(http.MethodGet, "/api/carousel_image", nil)
		rr := httptest.NewRecorder()

		h.ListCarousel(rr, req)

		if rr.Code != http.StatusOK {
			t.Fatalf("expected 200, got %d", rr.Code)
		}
		if !uc.listCarouselCalled {
			t.Fatalf("usecase should be called")
		}

		var res []map[string]any
		if err := json.NewDecoder(rr.Body).Decode(&res); err != nil {
			t.Fatalf("unexpected decode error: %v", err)
		}
		if len(res) != 1 {
			t.Fatalf("expected 1 item, got %d", len(res))
		}
		if res[0]["apiPath"] != "https://signed.example.com/carousel.png" {
			t.Fatalf("unexpected apiPath: %v", res[0]["apiPath"])
		}
		product, ok := res[0]["product"].(map[string]any)
		if !ok {
			t.Fatalf("expected product object")
		}
		if product["uuid"] != "product-uuid" {
			t.Fatalf("unexpected product uuid: %v", product["uuid"])
		}
	})
}

func TestProductCreate(t *testing.T) {
	t.Run("不正なリクエストボディならバリデーションエラーで失敗する", func(t *testing.T) {
		uc := &stubProductUC{}
		h := NewProductHandler(uc)

		req := httptest.NewRequest(http.MethodPost, "/api/product", nil)
		rr := httptest.NewRecorder()

		h.Create(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Fatalf("expected 400, got %d", rr.Code)
		}
		if uc.createCalled {
			t.Fatalf("usecase should not be called")
		}
	})

	t.Run("有効な入力を渡したとき作成した商品UUIDを返す", func(t *testing.T) {
		uc := &stubProductUC{
			createRes: primitive.UUID("b3d2a889-e2aa-4430-a030-1dcf2dbf13af"),
		}
		h := NewProductHandler(uc)

		reqBody := `{
			"name":"sample product",
			"description":"desc",
			"price":1000,
			"isRecommend":true,
			"isActive":true,
			"category":{"uuid":"category-uuid","name":"category"},
			"target":{"uuid":"target-uuid","name":"target"},
			"tags":[{"uuid":"tag-uuid","name":"tag"}],
			"siteDetails":[{"uuid":"detail-uuid","detailUrl":"https://example.com","salesSite":{"uuid":"sales-site-uuid","name":"site"}}]
		}`
		req := httptest.NewRequest(http.MethodPost, "/api/product", strings.NewReader(reqBody))
		rr := httptest.NewRecorder()

		h.Create(rr, req)

		if rr.Code != http.StatusOK {
			t.Fatalf("expected 200, got %d", rr.Code)
		}
		if !uc.createCalled {
			t.Fatalf("usecase should be called")
		}

		var res map[string]any
		if err := json.NewDecoder(rr.Body).Decode(&res); err != nil {
			t.Fatalf("unexpected decode error: %v", err)
		}
		if res["uuid"] != "b3d2a889-e2aa-4430-a030-1dcf2dbf13af" {
			t.Fatalf("unexpected uuid: %v", res["uuid"])
		}
	})
}

func TestProductExportCSV(t *testing.T) {
	t.Run("ユースケースがエラーを返したとき内部エラーで失敗する", func(t *testing.T) {
		uc := &stubProductUC{exportCSVErr: context.DeadlineExceeded}
		h := NewProductHandler(uc)

		req := httptest.NewRequest(http.MethodGet, "/api/csv/product", nil)
		rr := httptest.NewRecorder()

		h.ExportCSV(rr, req)

		if rr.Code != http.StatusInternalServerError {
			t.Fatalf("expected 500, got %d", rr.Code)
		}
	})

	t.Run("有効な入力を渡したときCSVを返す", func(t *testing.T) {
		uc := &stubProductUC{
			exportCSVRes: []*usecaseProductQuery.ProductCSVRow{
				{
					ID:           1,
					Name:         "product",
					Price:        1200,
					CategoryName: "category",
					TargetName:   "target",
				},
			},
		}
		h := NewProductHandler(uc)

		req := httptest.NewRequest(http.MethodGet, "/api/csv/product", nil)
		rr := httptest.NewRecorder()

		h.ExportCSV(rr, req)

		if rr.Code != http.StatusOK {
			t.Fatalf("expected 200, got %d", rr.Code)
		}
		if contentType := rr.Header().Get("Content-Type"); contentType != "text/csv; charset=utf-8" {
			t.Fatalf("unexpected content type: %s", contentType)
		}
		got := rr.Body.String()
		if !strings.Contains(got, "ID,Name,Price,CategoryName,TargetName") {
			t.Fatalf("unexpected csv header: %s", got)
		}
		if !strings.Contains(got, "1,product,1200,category,target") {
			t.Fatalf("unexpected csv body: %s", got)
		}
	})
}

func TestProductUploadCSV(t *testing.T) {
	t.Run("csvファイルがないときバリデーションエラーで失敗する", func(t *testing.T) {
		uc := &stubProductUC{}
		h := NewProductHandler(uc)

		req := httptest.NewRequest(http.MethodPost, "/api/csv/product", nil)
		rr := httptest.NewRecorder()

		h.UploadCSV(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Fatalf("expected 400, got %d", rr.Code)
		}
		if uc.uploadCSVCalled {
			t.Fatalf("usecase should not be called")
		}
	})

	t.Run("有効なcsvを渡したときCSVアップロードに成功する", func(t *testing.T) {
		uc := &stubProductUC{}
		h := NewProductHandler(uc)

		req := newProductCSVUploadRequest(
			t,
			"products.csv",
			[]byte("id,name,price,categoryName,targetName\n1,new product,1200,new category,new target\n"),
		)
		rr := httptest.NewRecorder()

		h.UploadCSV(rr, req)

		if rr.Code != http.StatusOK {
			t.Fatalf("expected 200, got %d", rr.Code)
		}
		if !uc.uploadCSVCalled {
			t.Fatalf("usecase should be called")
		}
		if len(uc.uploadCSVRows) != 1 {
			t.Fatalf("expected 1 row, got %d", len(uc.uploadCSVRows))
		}
		if uc.uploadCSVRows[0].ID != 1 {
			t.Fatalf("unexpected id: %d", uc.uploadCSVRows[0].ID)
		}
		if uc.uploadCSVRows[0].CategoryName != "new category" {
			t.Fatalf("unexpected category name: %s", uc.uploadCSVRows[0].CategoryName)
		}
	})
}

func newProductCSVUploadRequest(t *testing.T, fileName string, fileData []byte) *http.Request {
	t.Helper()

	var body bytes.Buffer
	writer := multipart.NewWriter(&body)

	part, err := writer.CreateFormFile("csv", fileName)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if _, err := io.Copy(part, bytes.NewReader(fileData)); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if err := writer.Close(); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	req := httptest.NewRequest(http.MethodPost, "/api/csv/product", &body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	return req
}
