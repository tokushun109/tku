package handler

import (
	"bytes"
	"context"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	usecaseProduct "github.com/tokushun109/tku/clean-backend/internal/usecase/product"
	usecaseProductQuery "github.com/tokushun109/tku/clean-backend/internal/usecase/product/query"
)

type stubProductUC struct {
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

func TestProductCreateImage(t *testing.T) {
	t.Run("空ファイルを渡したときバリデーションエラーで失敗する", func(t *testing.T) {
		uc := &stubProductUC{}
		h := NewProductHandler(uc)

		req := newProductImageUploadRequest(t, "empty.png", []byte{})
		rr := httptest.NewRecorder()

		h.CreateImage(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Fatalf("expected 400, got %d", rr.Code)
		}
		if uc.createProductImagesCalled {
			t.Fatalf("usecase should not be called when file is empty")
		}
	})

	t.Run("20MBを超える画像ファイルを渡したときバリデーションエラーで失敗する", func(t *testing.T) {
		uc := &stubProductUC{}
		h := NewProductHandler(uc)

		req := newProductImageUploadRequest(t, "too-large.png", bytes.Repeat([]byte("a"), maxProductImageSize+1))
		rr := httptest.NewRecorder()

		h.CreateImage(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Fatalf("expected 400, got %d", rr.Code)
		}
		if uc.createProductImagesCalled {
			t.Fatalf("usecase should not be called when payload is too large")
		}
	})

	t.Run("有効な画像ファイルを渡したとき画像追加に成功する", func(t *testing.T) {
		uc := &stubProductUC{}
		h := NewProductHandler(uc)

		req := newProductImageUploadRequest(t, "valid.png", []byte{0x89, 'P', 'N', 'G', '\r', '\n', 0x1a, '\n'})
		rr := httptest.NewRecorder()

		h.CreateImage(rr, req)

		if rr.Code != http.StatusOK {
			t.Fatalf("expected 200, got %d", rr.Code)
		}
		if !uc.createProductImagesCalled {
			t.Fatalf("usecase should be called")
		}
		if uc.createProductImagesReq.productUUID != "product-uuid" {
			t.Fatalf("unexpected product uuid: %s", uc.createProductImagesReq.productUUID)
		}
		if len(uc.createProductImagesReq.files) != 1 {
			t.Fatalf("expected 1 file, got %d", len(uc.createProductImagesReq.files))
		}
		if uc.createProductImagesReq.files[0].Name != "valid.png" {
			t.Fatalf("unexpected file name: %s", uc.createProductImagesReq.files[0].Name)
		}
		if len(uc.createProductImagesReq.files[0].Data) == 0 {
			t.Fatalf("expected non-empty file bytes")
		}
	})
}

func newProductImageUploadRequest(t *testing.T, fileName string, fileData []byte) *http.Request {
	t.Helper()

	var body bytes.Buffer
	writer := multipart.NewWriter(&body)
	part, err := writer.CreateFormFile("file0", fileName)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if _, err := io.Copy(part, bytes.NewReader(fileData)); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if err := writer.Close(); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	req := httptest.NewRequest(http.MethodPost, "/api/product/product-uuid/product_image", &body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req = mux.SetURLVars(req, map[string]string{"product_uuid": "product-uuid"})
	return req
}
