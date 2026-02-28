package product

import (
	"context"

	"github.com/tokushun109/tku/clean-backend/internal/domain/primitive"
	usecaseProductQuery "github.com/tokushun109/tku/clean-backend/internal/usecase/product/query"
)

const (
	ListModeAll    = "all"
	ListModeActive = "active"
)

type QueryUsecase interface {
	List(ctx context.Context, mode string, category string, target string) ([]*usecaseProductQuery.Product, error)
	ListByCategory(ctx context.Context, mode string, category string, target string) ([]*usecaseProductQuery.CategoryProducts, error)
	ListCarousel(ctx context.Context) ([]*usecaseProductQuery.CarouselItem, error)
	Get(ctx context.Context, productUUID string) (*usecaseProductQuery.Product, error)
	ExportCSV(ctx context.Context) ([]*usecaseProductQuery.ProductCSVRow, error)
}

type CommandUsecase interface {
	Create(ctx context.Context, input CreateProductInput) (primitive.UUID, error)
	Duplicate(ctx context.Context, rawURL string) error
	Update(ctx context.Context, productUUID string, input UpdateProductInput) error
	Delete(ctx context.Context, productUUID string) error
	UploadCSV(ctx context.Context, rows []ProductCSVInputRow) error
	GetProductImageBlob(ctx context.Context, productImageUUID string) (*ProductImageBlob, error)
	CreateProductImages(ctx context.Context, productUUID string, files []ProductImageUploadFile, isChanged bool, orderMap map[int]int) error
	DeleteProductImage(ctx context.Context, productUUID string, productImageUUID string) error
}

type Usecase interface {
	QueryUsecase
	CommandUsecase
}

type Service struct {
	queryUC   QueryUsecase
	commandUC CommandUsecase
}

var _ Usecase = (*Service)(nil)

func New(queryUC QueryUsecase, commandUC CommandUsecase) *Service {
	return &Service{
		queryUC:   queryUC,
		commandUC: commandUC,
	}
}

func (s *Service) List(ctx context.Context, mode string, category string, target string) ([]*usecaseProductQuery.Product, error) {
	return s.queryUC.List(ctx, mode, category, target)
}

func (s *Service) ListByCategory(ctx context.Context, mode string, category string, target string) ([]*usecaseProductQuery.CategoryProducts, error) {
	return s.queryUC.ListByCategory(ctx, mode, category, target)
}

func (s *Service) ListCarousel(ctx context.Context) ([]*usecaseProductQuery.CarouselItem, error) {
	return s.queryUC.ListCarousel(ctx)
}

func (s *Service) Get(ctx context.Context, productUUID string) (*usecaseProductQuery.Product, error) {
	return s.queryUC.Get(ctx, productUUID)
}

func (s *Service) ExportCSV(ctx context.Context) ([]*usecaseProductQuery.ProductCSVRow, error) {
	return s.queryUC.ExportCSV(ctx)
}

func (s *Service) Create(ctx context.Context, input CreateProductInput) (primitive.UUID, error) {
	return s.commandUC.Create(ctx, input)
}

func (s *Service) Duplicate(ctx context.Context, rawURL string) error {
	return s.commandUC.Duplicate(ctx, rawURL)
}

func (s *Service) Update(ctx context.Context, productUUID string, input UpdateProductInput) error {
	return s.commandUC.Update(ctx, productUUID, input)
}

func (s *Service) Delete(ctx context.Context, productUUID string) error {
	return s.commandUC.Delete(ctx, productUUID)
}

func (s *Service) UploadCSV(ctx context.Context, rows []ProductCSVInputRow) error {
	return s.commandUC.UploadCSV(ctx, rows)
}

func (s *Service) GetProductImageBlob(ctx context.Context, productImageUUID string) (*ProductImageBlob, error) {
	return s.commandUC.GetProductImageBlob(ctx, productImageUUID)
}

func (s *Service) CreateProductImages(ctx context.Context, productUUID string, files []ProductImageUploadFile, isChanged bool, orderMap map[int]int) error {
	return s.commandUC.CreateProductImages(ctx, productUUID, files, isChanged, orderMap)
}

func (s *Service) DeleteProductImage(ctx context.Context, productUUID string, productImageUUID string) error {
	return s.commandUC.DeleteProductImage(ctx, productUUID, productImageUUID)
}
