package command

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"

	domainCategory "github.com/tokushun109/tku/clean-backend/internal/domain/category"
	"github.com/tokushun109/tku/clean-backend/internal/domain/primitive"
	domainProduct "github.com/tokushun109/tku/clean-backend/internal/domain/product"
	domainSalesSite "github.com/tokushun109/tku/clean-backend/internal/domain/sales_site"
	domainSiteDetail "github.com/tokushun109/tku/clean-backend/internal/domain/site_detail"
	domainTag "github.com/tokushun109/tku/clean-backend/internal/domain/tag"
	domainTarget "github.com/tokushun109/tku/clean-backend/internal/domain/target"
	"github.com/tokushun109/tku/clean-backend/internal/usecase"
	usecaseProduct "github.com/tokushun109/tku/clean-backend/internal/usecase/product"
)

type Usecase interface {
	Create(ctx context.Context, input usecaseProduct.CreateProductInput) (primitive.UUID, error)
	Duplicate(ctx context.Context, rawURL string) error
	Update(ctx context.Context, productUUID string, input usecaseProduct.UpdateProductInput) error
	Delete(ctx context.Context, productUUID string) error
	UploadCSV(ctx context.Context, rows []usecaseProduct.ProductCSVInputRow) error
	GetProductImageBlob(ctx context.Context, productImageUUID string) (*usecaseProduct.ProductImageBlob, error)
	CreateProductImages(ctx context.Context, productUUID string, files []usecaseProduct.ProductImageUploadFile, isChanged bool, orderMap map[int]int) error
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
	duplicateSource  usecaseProduct.DuplicateSource
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
	duplicateSource usecaseProduct.DuplicateSource,
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
		duplicateSource:  duplicateSource,
		storage:          storage,
		uuidGen:          uuidGen,
		txManager:        txManager,
	}
}

func (s *Service) Create(ctx context.Context, input usecaseProduct.CreateProductInput) (primitive.UUID, error) {
	categoryID, err := s.resolveCategoryID(ctx, input.CategoryUUID)
	if err != nil {
		return "", err
	}
	targetID, err := s.resolveTargetID(ctx, input.TargetUUID)
	if err != nil {
		return "", err
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
			return "", usecase.NewAppErrorWithMessage(usecase.ErrInvalidInput, err.Error())
		}
		return "", usecase.NewAppErrorWithMessage(usecase.ErrInternal, err.Error())
	}

	tagIDs, err := s.resolveTagIDs(ctx, input.TagUUIDs)
	if err != nil {
		return "", err
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
			return "", usecase.NewAppErrorWithMessage(usecase.ErrInvalidInput, err.Error())
		}
		return "", usecase.NewAppErrorWithMessage(usecase.ErrInternal, err.Error())
	}

	return newProduct.UUID(), nil
}

func (s *Service) Duplicate(ctx context.Context, rawURL string) error {
	if s.duplicateSource == nil {
		return usecase.NewAppErrorWithMessage(usecase.ErrInternal, "duplicate source is nil")
	}

	duplicateURL, err := normalizeDuplicateProductURL(rawURL)
	if err != nil {
		if errors.Is(err, usecase.ErrInvalidInput) || errors.Is(err, primitive.ErrInvalidURL) {
			return usecase.NewAppErrorWithMessage(usecase.ErrInvalidInput, err.Error())
		}
		return usecase.NewAppErrorWithMessage(usecase.ErrInternal, err.Error())
	}

	duplicatedProduct, err := s.duplicateSource.Duplicate(ctx, duplicateURL)
	if err != nil {
		if errors.Is(err, usecase.ErrInvalidInput) || errors.Is(err, primitive.ErrInvalidURL) {
			return usecase.NewAppErrorWithMessage(usecase.ErrInvalidInput, err.Error())
		}
		return usecase.NewAppErrorWithMessage(usecase.ErrInternal, err.Error())
	}
	if duplicatedProduct == nil {
		return usecase.NewAppErrorWithMessage(usecase.ErrInternal, "duplicated product data is nil")
	}

	newProduct, err := domainProduct.New(
		s.uuidGen.New(),
		duplicatedProduct.Name,
		duplicatedProduct.Description,
		duplicatedProduct.Price,
		false,
		false,
		nil,
		nil,
	)
	if err != nil {
		if isInvalidProductInputError(err) {
			return usecase.NewAppErrorWithMessage(usecase.ErrInvalidInput, err.Error())
		}
		return usecase.NewAppErrorWithMessage(usecase.ErrInternal, err.Error())
	}

	salesSite, err := s.findSalesSiteByName(ctx, "creema")
	if err != nil {
		return usecase.NewAppErrorWithMessage(usecase.ErrInternal, err.Error())
	}
	if salesSite == nil {
		return usecase.NewAppErrorWithMessage(usecase.ErrInternal, "sales site is not found: creema")
	}

	uploadedPaths := make([]string, 0, len(duplicatedProduct.Images))
	if err := s.txManager.WithinTransaction(ctx, func(txCtx context.Context) error {
		tagIDs, err := s.resolveOrCreateTagIDsByNames(txCtx, duplicatedProduct.Tags)
		if err != nil {
			return err
		}

		createdProductID, err := s.productRepo.Create(txCtx, newProduct)
		if err != nil {
			return err
		}

		if err := s.productRepo.ReplaceTags(txCtx, createdProductID, tagIDs); err != nil {
			return err
		}

		siteDetail, err := domainSiteDetail.New(
			s.uuidGen.New(),
			duplicateURL,
			createdProductID.Value(),
			salesSite.ID().Value(),
		)
		if err != nil {
			return err
		}
		if err := s.siteDetailRepo.ReplaceByProductID(txCtx, createdProductID, []*domainSiteDetail.SiteDetail{siteDetail}); err != nil {
			return err
		}

		for i, image := range duplicatedProduct.Images {
			imageMimeType, err := domainProduct.NewProductImageMimeType(http.DetectContentType(image.Data))
			if err != nil {
				return err
			}

			imageUUID := s.uuidGen.New()
			imagePath, err := buildProductImagePath(imageUUID, imageMimeType)
			if err != nil {
				return err
			}
			if err := s.storage.Put(txCtx, imagePath.Value(), imageMimeType.Value(), image.Data); err != nil {
				return err
			}
			uploadedPaths = append(uploadedPaths, imagePath.Value())

			imageName := strings.TrimSpace(image.Name)
			if imageName == "" {
				imageName = fmt.Sprintf("image-%d%s", i, imageMimeType.Extension())
			}

			productImage, err := domainProduct.NewProductImage(
				imageUUID,
				imageName,
				imageMimeType.Value(),
				imagePath.Value(),
				len(duplicatedProduct.Images)-i,
				createdProductID.Value(),
			)
			if err != nil {
				return err
			}
			if err := s.productImageRepo.Create(txCtx, productImage); err != nil {
				return err
			}
		}

		return nil
	}); err != nil {
		for _, path := range uploadedPaths {
			if delErr := s.storage.Delete(ctx, path); delErr != nil {
				log.Printf("[WARN] product duplicate rollback delete failed: path=%s err=%v", path, delErr)
			}
		}

		if errors.Is(err, usecase.ErrInvalidInput) || errors.Is(err, domainTag.ErrInvalidName) || isInvalidProductInputError(err) {
			return usecase.NewAppErrorWithMessage(usecase.ErrInvalidInput, err.Error())
		}
		return usecase.NewAppErrorWithMessage(usecase.ErrInternal, err.Error())
	}

	return nil
}

func (s *Service) Update(ctx context.Context, productUUID string, input usecaseProduct.UpdateProductInput) error {
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
			imageMap[image.UUID().Value()] = image
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
			if _, keep := requestedImageOrders[image.UUID().Value()]; keep {
				continue
			}
			deleted, err := s.productImageRepo.DeleteByUUID(txCtx, image.UUID())
			if err != nil {
				return err
			}
			if deleted {
				deletedImagePaths = append(deletedImagePaths, image.Path().Value())
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
				deletedImagePaths = append(deletedImagePaths, image.Path().Value())
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

func (s *Service) UploadCSV(ctx context.Context, rows []usecaseProduct.ProductCSVInputRow) error {
	normalizedRows, err := normalizeProductCSVRows(rows)
	if err != nil {
		return usecase.NewAppErrorWithMessage(usecase.ErrInvalidInput, err.Error())
	}

	if err := s.txManager.WithinTransaction(ctx, func(txCtx context.Context) error {
		categoryCache := make(map[string]*domainCategory.Category)
		targetCache := make(map[string]*domainTarget.Target)

		for _, row := range normalizedRows {
			productID, err := primitive.NewID(row.id)
			if err != nil {
				return fmt.Errorf("product id is invalid: id=%d: %w", row.id, usecase.ErrInvalidInput)
			}

			current, err := s.productRepo.FindByID(txCtx, productID)
			if err != nil {
				return err
			}
			if current == nil {
				return fmt.Errorf("product not found: id=%d: %w", row.id, usecase.ErrInvalidInput)
			}

			var categoryID *uint
			if row.categoryName != nil {
				categoryEntity, err := s.findOrCreateCategoryByName(txCtx, *row.categoryName, categoryCache)
				if err != nil {
					return err
				}
				categoryIDValue := categoryEntity.ID().Value()
				categoryID = &categoryIDValue
			}

			var targetID *uint
			if row.targetName != nil {
				targetEntity, err := s.findOrCreateTargetByName(txCtx, *row.targetName, targetCache)
				if err != nil {
					return err
				}
				targetIDValue := targetEntity.ID().Value()
				targetID = &targetIDValue
			}

			description := ""
			if current.Description() != nil {
				description = current.Description().Value()
			}

			if err := current.ChangeProduct(
				row.name,
				description,
				row.price,
				current.IsActive(),
				current.IsRecommend(),
				categoryID,
				targetID,
			); err != nil {
				return fmt.Errorf("row update failed: id=%d: %w", row.id, err)
			}

			updated, err := s.productRepo.Update(txCtx, current)
			if err != nil {
				return err
			}
			if !updated {
				return fmt.Errorf("product update failed: id=%d: %w", row.id, usecase.ErrInvalidInput)
			}
		}

		return nil
	}); err != nil {
		if errors.Is(err, usecase.ErrInvalidInput) ||
			errors.Is(err, domainCategory.ErrInvalidName) ||
			errors.Is(err, domainTarget.ErrInvalidName) ||
			isInvalidProductInputError(err) {
			return usecase.NewAppErrorWithMessage(usecase.ErrInvalidInput, err.Error())
		}

		return usecase.NewAppErrorWithMessage(usecase.ErrInternal, err.Error())
	}

	return nil
}

func (s *Service) GetProductImageBlob(ctx context.Context, productImageUUID string) (*usecaseProduct.ProductImageBlob, error) {
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

	body, err := s.storage.Get(ctx, image.Path().Value())
	if err != nil {
		if errors.Is(err, usecase.ErrStorageNotFound) {
			return nil, usecase.NewAppError(usecase.ErrNotFound)
		}
		return nil, usecase.NewAppErrorWithMessage(usecase.ErrInternal, err.Error())
	}

	return &usecaseProduct.ProductImageBlob{
		ContentType: image.MimeType().Value(),
		Body:        body,
	}, nil
}

func (s *Service) CreateProductImages(ctx context.Context, productUUID string, files []usecaseProduct.ProductImageUploadFile, isChanged bool, orderMap map[int]int) error {
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

			if err := s.storage.Put(txCtx, imagePath.Value(), imageMimeType.Value(), file.Data); err != nil {
				return err
			}
			uploadedPaths = append(uploadedPaths, imagePath.Value())

			order := 0
			if isChanged {
				order = orderMap[i]
			}

			image, err := domainProduct.NewProductImage(
				imageUUID,
				file.Name,
				imageMimeType.Value(),
				imagePath.Value(),
				order,
				product.ID().Value(),
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
	if image.ProductID() != product.ID().Value() {
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

	if delErr := s.storage.Delete(ctx, image.Path().Value()); delErr != nil {
		log.Printf("[WARN] product image delete object failed: path=%s err=%v", image.Path().Value(), delErr)
	}

	return nil
}
