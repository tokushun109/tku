package product

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	domainCategory "github.com/tokushun109/tku/clean-backend/internal/domain/category"
	"github.com/tokushun109/tku/clean-backend/internal/domain/primitive"
	domainProduct "github.com/tokushun109/tku/clean-backend/internal/domain/product"
	domainSalesSite "github.com/tokushun109/tku/clean-backend/internal/domain/sales_site"
	domainSiteDetail "github.com/tokushun109/tku/clean-backend/internal/domain/site_detail"
	domainTag "github.com/tokushun109/tku/clean-backend/internal/domain/tag"
	domainTarget "github.com/tokushun109/tku/clean-backend/internal/domain/target"
	"github.com/tokushun109/tku/clean-backend/internal/usecase"
	usecaseProductQuery "github.com/tokushun109/tku/clean-backend/internal/usecase/product/query"
)

const (
	ListModeAll    = "all"
	ListModeActive = "active"

	defaultProductImagePresignTTL = 30 * time.Minute
)

type Usecase interface {
	List(ctx context.Context, mode string, category string, target string) ([]*usecaseProductQuery.Product, error)
	Get(ctx context.Context, productUUID string) (*usecaseProductQuery.Product, error)
	Create(ctx context.Context, input CreateProductInput) (*usecaseProductQuery.Product, error)
	Update(ctx context.Context, productUUID string, input UpdateProductInput) error
	Delete(ctx context.Context, productUUID string) error
	GetProductImageBlob(ctx context.Context, productImageUUID string) (*ProductImageBlob, error)
	CreateProductImages(ctx context.Context, productUUID string, files []ProductImageUploadFile, isChanged bool, orderMap map[int]int) error
	DeleteProductImage(ctx context.Context, productUUID string, productImageUUID string) error
}

type Service struct {
	productRepo      domainProduct.Repository
	productImageRepo domainProduct.ProductImageRepository
	siteDetailRepo   domainSiteDetail.Repository
	categoryRepo     domainCategory.Repository
	targetRepo       domainTarget.Repository
	tagRepo          domainTag.Repository
	salesSiteRepo    domainSalesSite.Repository
	queryReader      usecaseProductQuery.Reader
	storage          usecase.Storage
	uuidGen          usecase.UUIDGenerator
	txManager        usecase.TxManager
}

var _ Usecase = (*Service)(nil)

func New(
	productRepo domainProduct.Repository,
	productImageRepo domainProduct.ProductImageRepository,
	siteDetailRepo domainSiteDetail.Repository,
	categoryRepo domainCategory.Repository,
	targetRepo domainTarget.Repository,
	tagRepo domainTag.Repository,
	salesSiteRepo domainSalesSite.Repository,
	queryReader usecaseProductQuery.Reader,
	storage usecase.Storage,
	uuidGen usecase.UUIDGenerator,
	txManager usecase.TxManager,
) *Service {
	return &Service{
		productRepo:      productRepo,
		productImageRepo: productImageRepo,
		siteDetailRepo:   siteDetailRepo,
		categoryRepo:     categoryRepo,
		targetRepo:       targetRepo,
		tagRepo:          tagRepo,
		salesSiteRepo:    salesSiteRepo,
		queryReader:      queryReader,
		storage:          storage,
		uuidGen:          uuidGen,
		txManager:        txManager,
	}
}

func (s *Service) List(ctx context.Context, mode string, category string, target string) ([]*usecaseProductQuery.Product, error) {
	if !isValidListMode(mode) || strings.TrimSpace(category) == "" || strings.TrimSpace(target) == "" {
		return nil, usecase.NewAppError(usecase.ErrInvalidInput)
	}

	products, err := s.queryReader.ListProducts(ctx, usecaseProductQuery.ListProductsQuery{
		Mode:     mode,
		Category: category,
		Target:   target,
	})
	if err != nil {
		return nil, usecase.NewAppErrorWithMessage(usecase.ErrInternal, err.Error())
	}

	if err := s.attachPresignedImageURLs(ctx, products); err != nil {
		return nil, err
	}

	return products, nil
}

func (s *Service) Get(ctx context.Context, productUUID string) (*usecaseProductQuery.Product, error) {
	if _, err := primitive.NewUUID(productUUID); err != nil {
		return nil, usecase.NewAppError(usecase.ErrInvalidInput)
	}

	product, err := s.queryReader.GetProductByUUID(ctx, productUUID)
	if err != nil {
		return nil, usecase.NewAppErrorWithMessage(usecase.ErrInternal, err.Error())
	}
	if product == nil {
		return nil, usecase.NewAppError(usecase.ErrNotFound)
	}

	if err := s.attachPresignedImageURLs(ctx, []*usecaseProductQuery.Product{product}); err != nil {
		return nil, err
	}

	return product, nil
}

func (s *Service) Create(ctx context.Context, input CreateProductInput) (*usecaseProductQuery.Product, error) {
	categoryID, err := s.resolveCategoryID(ctx, input.CategoryUUID)
	if err != nil {
		return nil, err
	}
	targetID, err := s.resolveTargetID(ctx, input.TargetUUID)
	if err != nil {
		return nil, err
	}

	newUUID := s.uuidGen.New()
	newProduct, err := domainProduct.New(
		newUUID,
		input.Name,
		input.Description,
		input.Price,
		input.IsActive,
		input.IsRecommend,
		categoryID,
		targetID,
	)
	if err != nil {
		if isInvalidProductInputError(err) {
			return nil, usecase.NewAppErrorWithMessage(usecase.ErrInvalidInput, err.Error())
		}
		return nil, usecase.NewAppErrorWithMessage(usecase.ErrInternal, err.Error())
	}

	tagIDs, err := s.resolveTagIDs(ctx, input.TagUUIDs)
	if err != nil {
		return nil, err
	}

	if err := s.txManager.WithinTransaction(ctx, func(txCtx context.Context) error {
		createdProductID, createErr := s.productRepo.Create(txCtx, newProduct)
		if createErr != nil {
			return createErr
		}

		if err := s.productRepo.ReplaceTags(txCtx, createdProductID, tagIDs); err != nil {
			return err
		}

		siteDetails, buildErr := s.buildSiteDetails(txCtx, createdProductID, input.SiteDetails)
		if buildErr != nil {
			return buildErr
		}
		if err := s.siteDetailRepo.ReplaceByProductID(txCtx, createdProductID, siteDetails); err != nil {
			return err
		}
		return nil
	}); err != nil {
		if errors.Is(err, usecase.ErrInvalidInput) || isInvalidProductInputError(err) {
			return nil, usecase.NewAppErrorWithMessage(usecase.ErrInvalidInput, err.Error())
		}
		return nil, usecase.NewAppErrorWithMessage(usecase.ErrInternal, err.Error())
	}

	product, err := s.queryReader.GetProductByUUID(ctx, newUUID)
	if err != nil {
		return nil, usecase.NewAppErrorWithMessage(usecase.ErrInternal, err.Error())
	}
	if product == nil {
		return nil, usecase.NewAppError(usecase.ErrNotFound)
	}

	if err := s.attachPresignedImageURLs(ctx, []*usecaseProductQuery.Product{product}); err != nil {
		return nil, err
	}

	return product, nil
}

func (s *Service) Update(ctx context.Context, productUUID string, input UpdateProductInput) error {
	uuid, err := primitive.NewUUID(productUUID)
	if err != nil {
		return usecase.NewAppError(usecase.ErrInvalidInput)
	}

	current, err := s.productRepo.FindByUUID(ctx, uuid)
	if err != nil {
		return usecase.NewAppErrorWithMessage(usecase.ErrInternal, err.Error())
	}
	if current == nil {
		return usecase.NewAppError(usecase.ErrNotFound)
	}

	categoryID, err := s.resolveCategoryID(ctx, input.CategoryUUID)
	if err != nil {
		return err
	}
	targetID, err := s.resolveTargetID(ctx, input.TargetUUID)
	if err != nil {
		return err
	}

	if err := current.ChangeProduct(
		input.Name,
		input.Description,
		input.Price,
		input.IsActive,
		input.IsRecommend,
		categoryID,
		targetID,
	); err != nil {
		if isInvalidProductInputError(err) {
			return usecase.NewAppErrorWithMessage(usecase.ErrInvalidInput, err.Error())
		}
		return usecase.NewAppErrorWithMessage(usecase.ErrInternal, err.Error())
	}

	tagIDs, err := s.resolveTagIDs(ctx, input.TagUUIDs)
	if err != nil {
		return err
	}

	requestedImageOrders := make(map[string]int, len(input.ProductImages))
	for _, image := range input.ProductImages {
		if _, err := primitive.NewUUID(image.UUID); err != nil {
			return usecase.NewAppError(usecase.ErrInvalidInput)
		}
		if _, err := domainProduct.NewProductImageOrder(image.Order); err != nil {
			return usecase.NewAppErrorWithMessage(usecase.ErrInvalidInput, err.Error())
		}
		if _, exists := requestedImageOrders[image.UUID]; exists {
			return usecase.NewAppError(usecase.ErrInvalidInput)
		}
		requestedImageOrders[image.UUID] = image.Order
	}

	deletedImagePaths := make([]string, 0)
	if err := s.txManager.WithinTransaction(ctx, func(txCtx context.Context) error {
		updated, err := s.productRepo.Update(txCtx, current)
		if err != nil {
			return err
		}
		if !updated {
			return usecase.ErrNotFound
		}

		if err := s.productRepo.ReplaceTags(txCtx, current.ID(), tagIDs); err != nil {
			return err
		}

		siteDetails, buildErr := s.buildSiteDetails(txCtx, current.ID(), input.SiteDetails)
		if buildErr != nil {
			return buildErr
		}
		if err := s.siteDetailRepo.ReplaceByProductID(txCtx, current.ID(), siteDetails); err != nil {
			return err
		}

		currentImages, err := s.productImageRepo.FindByProductID(txCtx, current.ID())
		if err != nil {
			return err
		}
		imageMap := make(map[string]*domainProduct.ProductImage, len(currentImages))
		for _, image := range currentImages {
			imageMap[image.UUID().String()] = image
		}

		for imageUUID, order := range requestedImageOrders {
			image, ok := imageMap[imageUUID]
			if !ok {
				return usecase.ErrInvalidInput
			}
			updated, err := s.productImageRepo.UpdateOrder(txCtx, image.UUID(), order)
			if err != nil {
				return err
			}
			if !updated {
				return usecase.ErrNotFound
			}
		}

		for _, image := range currentImages {
			if _, keep := requestedImageOrders[image.UUID().String()]; keep {
				continue
			}
			deleted, err := s.productImageRepo.DeleteByUUID(txCtx, image.UUID())
			if err != nil {
				return err
			}
			if deleted {
				deletedImagePaths = append(deletedImagePaths, image.Path().String())
			}
		}

		return nil
	}); err != nil {
		if errors.Is(err, usecase.ErrNotFound) {
			return usecase.NewAppError(usecase.ErrNotFound)
		}
		if errors.Is(err, usecase.ErrInvalidInput) || isInvalidProductInputError(err) {
			return usecase.NewAppErrorWithMessage(usecase.ErrInvalidInput, err.Error())
		}
		return usecase.NewAppErrorWithMessage(usecase.ErrInternal, err.Error())
	}

	for _, path := range deletedImagePaths {
		if delErr := s.storage.Delete(ctx, path); delErr != nil {
			log.Printf("[WARN] product update delete image object failed: path=%s err=%v", path, delErr)
		}
	}

	return nil
}

func (s *Service) Delete(ctx context.Context, productUUID string) error {
	uuid, err := primitive.NewUUID(productUUID)
	if err != nil {
		return usecase.NewAppError(usecase.ErrInvalidInput)
	}

	current, err := s.productRepo.FindByUUID(ctx, uuid)
	if err != nil {
		return usecase.NewAppErrorWithMessage(usecase.ErrInternal, err.Error())
	}
	if current == nil {
		return usecase.NewAppError(usecase.ErrNotFound)
	}

	deletedImagePaths := make([]string, 0)
	if err := s.txManager.WithinTransaction(ctx, func(txCtx context.Context) error {
		images, err := s.productImageRepo.FindByProductID(txCtx, current.ID())
		if err != nil {
			return err
		}
		for _, image := range images {
			deleted, err := s.productImageRepo.DeleteByUUID(txCtx, image.UUID())
			if err != nil {
				return err
			}
			if deleted {
				deletedImagePaths = append(deletedImagePaths, image.Path().String())
			}
		}

		if err := s.siteDetailRepo.DeleteByProductID(txCtx, current.ID()); err != nil {
			return err
		}
		if err := s.productRepo.ReplaceTags(txCtx, current.ID(), nil); err != nil {
			return err
		}

		deleted, err := s.productRepo.Delete(txCtx, uuid)
		if err != nil {
			return err
		}
		if !deleted {
			return usecase.ErrNotFound
		}
		return nil
	}); err != nil {
		if errors.Is(err, usecase.ErrNotFound) {
			return usecase.NewAppError(usecase.ErrNotFound)
		}
		return usecase.NewAppErrorWithMessage(usecase.ErrInternal, err.Error())
	}

	for _, path := range deletedImagePaths {
		if delErr := s.storage.Delete(ctx, path); delErr != nil {
			log.Printf("[WARN] product delete delete image object failed: path=%s err=%v", path, delErr)
		}
	}

	return nil
}

func (s *Service) GetProductImageBlob(ctx context.Context, productImageUUID string) (*ProductImageBlob, error) {
	uuid, err := primitive.NewUUID(productImageUUID)
	if err != nil {
		return nil, usecase.NewAppError(usecase.ErrInvalidInput)
	}

	image, err := s.productImageRepo.FindByUUID(ctx, uuid)
	if err != nil {
		return nil, usecase.NewAppErrorWithMessage(usecase.ErrInternal, err.Error())
	}
	if image == nil {
		return nil, usecase.NewAppError(usecase.ErrNotFound)
	}

	body, err := s.storage.Get(ctx, image.Path().String())
	if err != nil {
		if errors.Is(err, usecase.ErrStorageNotFound) {
			return nil, usecase.NewAppError(usecase.ErrNotFound)
		}
		return nil, usecase.NewAppErrorWithMessage(usecase.ErrInternal, err.Error())
	}

	return &ProductImageBlob{
		ContentType: image.MimeType().String(),
		Body:        body,
	}, nil
}

func (s *Service) CreateProductImages(ctx context.Context, productUUID string, files []ProductImageUploadFile, isChanged bool, orderMap map[int]int) error {
	uuid, err := primitive.NewUUID(productUUID)
	if err != nil {
		return usecase.NewAppError(usecase.ErrInvalidInput)
	}
	product, err := s.productRepo.FindByUUID(ctx, uuid)
	if err != nil {
		return usecase.NewAppErrorWithMessage(usecase.ErrInternal, err.Error())
	}
	if product == nil {
		return usecase.NewAppError(usecase.ErrNotFound)
	}
	if len(files) == 0 {
		return nil
	}

	uploadedPaths := make([]string, 0, len(files))
	if err := s.txManager.WithinTransaction(ctx, func(txCtx context.Context) error {
		for i, file := range files {
			mimeType := http.DetectContentType(file.Data)
			imageMimeType, err := domainProduct.NewProductImageMimeType(mimeType)
			if err != nil {
				return err
			}

			imageUUID := s.uuidGen.New()
			imagePath, err := buildProductImagePath(imageUUID, imageMimeType)
			if err != nil {
				return err
			}

			if err := s.storage.Put(txCtx, imagePath.String(), imageMimeType.String(), file.Data); err != nil {
				return err
			}
			uploadedPaths = append(uploadedPaths, imagePath.String())

			order := 0
			if isChanged {
				order = orderMap[i]
			}

			image, err := domainProduct.NewProductImage(
				imageUUID,
				file.Name,
				imageMimeType.String(),
				imagePath.String(),
				order,
				product.ID().Uint(),
			)
			if err != nil {
				return err
			}

			if err := s.productImageRepo.Create(txCtx, image); err != nil {
				return err
			}
		}
		return nil
	}); err != nil {
		for _, path := range uploadedPaths {
			if delErr := s.storage.Delete(ctx, path); delErr != nil {
				log.Printf("[WARN] product image create rollback delete failed: path=%s err=%v", path, delErr)
			}
		}

		if isInvalidProductInputError(err) {
			return usecase.NewAppErrorWithMessage(usecase.ErrInvalidInput, err.Error())
		}
		return usecase.NewAppErrorWithMessage(usecase.ErrInternal, err.Error())
	}

	return nil
}

func (s *Service) DeleteProductImage(ctx context.Context, productUUID string, productImageUUID string) error {
	parsedProductUUID, err := primitive.NewUUID(productUUID)
	if err != nil {
		return usecase.NewAppError(usecase.ErrInvalidInput)
	}
	parsedImageUUID, err := primitive.NewUUID(productImageUUID)
	if err != nil {
		return usecase.NewAppError(usecase.ErrInvalidInput)
	}

	product, err := s.productRepo.FindByUUID(ctx, parsedProductUUID)
	if err != nil {
		return usecase.NewAppErrorWithMessage(usecase.ErrInternal, err.Error())
	}
	if product == nil {
		return usecase.NewAppError(usecase.ErrNotFound)
	}

	image, err := s.productImageRepo.FindByUUID(ctx, parsedImageUUID)
	if err != nil {
		return usecase.NewAppErrorWithMessage(usecase.ErrInternal, err.Error())
	}
	if image == nil {
		return usecase.NewAppError(usecase.ErrNotFound)
	}
	if image.ProductID() != product.ID().Uint() {
		return usecase.NewAppError(usecase.ErrNotFound)
	}

	if err := s.txManager.WithinTransaction(ctx, func(txCtx context.Context) error {
		deleted, err := s.productImageRepo.DeleteByUUID(txCtx, image.UUID())
		if err != nil {
			return err
		}
		if !deleted {
			return usecase.ErrNotFound
		}
		return nil
	}); err != nil {
		if errors.Is(err, usecase.ErrNotFound) {
			return usecase.NewAppError(usecase.ErrNotFound)
		}
		return usecase.NewAppErrorWithMessage(usecase.ErrInternal, err.Error())
	}

	if delErr := s.storage.Delete(ctx, image.Path().String()); delErr != nil {
		log.Printf("[WARN] product image delete object failed: path=%s err=%v", image.Path().String(), delErr)
	}

	return nil
}

func (s *Service) resolveCategoryID(ctx context.Context, rawUUID string) (*uint, error) {
	trimmed := strings.TrimSpace(rawUUID)
	if trimmed == "" {
		return nil, nil
	}

	uuid, err := primitive.NewUUID(trimmed)
	if err != nil {
		return nil, usecase.NewAppError(usecase.ErrInvalidInput)
	}

	category, err := s.categoryRepo.FindByUUID(ctx, uuid)
	if err != nil {
		return nil, usecase.NewAppErrorWithMessage(usecase.ErrInternal, err.Error())
	}
	if category == nil {
		return nil, usecase.NewAppError(usecase.ErrInvalidInput)
	}

	id := category.ID().Uint()
	return &id, nil
}

func (s *Service) resolveTargetID(ctx context.Context, rawUUID string) (*uint, error) {
	trimmed := strings.TrimSpace(rawUUID)
	if trimmed == "" {
		return nil, nil
	}

	uuid, err := primitive.NewUUID(trimmed)
	if err != nil {
		return nil, usecase.NewAppError(usecase.ErrInvalidInput)
	}

	target, err := s.targetRepo.FindByUUID(ctx, uuid)
	if err != nil {
		return nil, usecase.NewAppErrorWithMessage(usecase.ErrInternal, err.Error())
	}
	if target == nil {
		return nil, usecase.NewAppError(usecase.ErrInvalidInput)
	}

	id := target.ID().Uint()
	return &id, nil
}

func (s *Service) resolveTagIDs(ctx context.Context, rawUUIDs []string) ([]primitive.ID, error) {
	if len(rawUUIDs) == 0 {
		return []primitive.ID{}, nil
	}

	seen := map[primitive.ID]struct{}{}
	tagIDs := make([]primitive.ID, 0, len(rawUUIDs))
	for _, rawUUID := range rawUUIDs {
		uuid, err := primitive.NewUUID(strings.TrimSpace(rawUUID))
		if err != nil {
			return nil, usecase.NewAppError(usecase.ErrInvalidInput)
		}

		tag, err := s.tagRepo.FindByUUID(ctx, uuid)
		if err != nil {
			return nil, usecase.NewAppErrorWithMessage(usecase.ErrInternal, err.Error())
		}
		if tag == nil {
			return nil, usecase.NewAppError(usecase.ErrInvalidInput)
		}
		tagID := tag.ID()
		if _, exists := seen[tagID]; exists {
			continue
		}
		seen[tagID] = struct{}{}
		tagIDs = append(tagIDs, tagID)
	}
	return tagIDs, nil
}

func (s *Service) buildSiteDetails(ctx context.Context, productID primitive.ID, inputs []SiteDetailInput) ([]*domainSiteDetail.SiteDetail, error) {
	details := make([]*domainSiteDetail.SiteDetail, 0, len(inputs))
	for _, input := range inputs {
		salesSiteUUID := strings.TrimSpace(input.SalesSiteUUID)
		if salesSiteUUID == "" {
			return nil, usecase.ErrInvalidInput
		}
		uuid, err := primitive.NewUUID(salesSiteUUID)
		if err != nil {
			return nil, usecase.ErrInvalidInput
		}

		salesSite, err := s.salesSiteRepo.FindByUUID(ctx, uuid)
		if err != nil {
			return nil, err
		}
		if salesSite == nil {
			return nil, usecase.ErrInvalidInput
		}

		siteDetail, err := domainSiteDetail.New(
			s.uuidGen.New(),
			input.DetailURL,
			productID.Uint(),
			salesSite.ID().Uint(),
		)
		if err != nil {
			return nil, err
		}
		details = append(details, siteDetail)
	}
	return details, nil
}

func (s *Service) attachPresignedImageURLs(ctx context.Context, products []*usecaseProductQuery.Product) error {
	for _, product := range products {
		for i := range product.ProductImages {
			if strings.TrimSpace(product.ProductImages[i].Path) == "" {
				continue
			}
			url, err := s.storage.PresignGet(ctx, product.ProductImages[i].Path, defaultProductImagePresignTTL)
			if err != nil {
				return usecase.NewAppErrorWithMessage(usecase.ErrInternal, err.Error())
			}
			product.ProductImages[i].APIPath = url
		}
	}
	return nil
}

func isValidListMode(mode string) bool {
	switch mode {
	case ListModeAll, ListModeActive:
		return true
	default:
		return false
	}
}

func buildProductImagePath(uuidStr string, mimeType domainProduct.ProductImageMimeType) (domainProduct.ProductImagePath, error) {
	if len(uuidStr) < 2 {
		return "", fmt.Errorf("invalid uuid length: %d", len(uuidStr))
	}

	rawPath := fmt.Sprintf(
		"img/product/%s/%s/%s%s",
		uuidStr[0:1],
		uuidStr[1:2],
		uuidStr,
		mimeType.Extension(),
	)

	return domainProduct.NewProductImagePath(rawPath)
}

func isInvalidProductInputError(err error) bool {
	if err == nil {
		return false
	}

	return errors.Is(err, primitive.ErrInvalidUUID) ||
		errors.Is(err, primitive.ErrInvalidURL) ||
		errors.Is(err, domainProduct.ErrInvalidName) ||
		errors.Is(err, domainProduct.ErrInvalidPrice) ||
		errors.Is(err, domainProduct.ErrInvalidCategoryID) ||
		errors.Is(err, domainProduct.ErrInvalidTargetID) ||
		errors.Is(err, domainProduct.ErrInvalidImageName) ||
		errors.Is(err, domainProduct.ErrInvalidImageMimeType) ||
		errors.Is(err, domainProduct.ErrInvalidImagePath) ||
		errors.Is(err, domainProduct.ErrInvalidImageOrder) ||
		errors.Is(err, domainProduct.ErrInvalidImageProductID) ||
		errors.Is(err, domainSiteDetail.ErrInvalidDetailURL) ||
		errors.Is(err, domainSiteDetail.ErrInvalidProductID) ||
		errors.Is(err, domainSiteDetail.ErrInvalidSalesSiteID)
}
