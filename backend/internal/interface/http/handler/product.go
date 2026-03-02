package handler

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/tokushun109/tku/backend/internal/interface/http/middleware"
	"github.com/tokushun109/tku/backend/internal/interface/http/presenter"
	"github.com/tokushun109/tku/backend/internal/interface/http/request"
	"github.com/tokushun109/tku/backend/internal/interface/http/response"
	"github.com/tokushun109/tku/backend/internal/usecase"
	usecaseProduct "github.com/tokushun109/tku/backend/internal/usecase/product"
	usecaseProductQuery "github.com/tokushun109/tku/backend/internal/usecase/product/query"
)

const maxProductCSVSize = 5 << 20 // 5MB

type ProductHandler struct {
	productUC usecaseProduct.Usecase
}

func NewProductHandler(productUC usecaseProduct.Usecase) *ProductHandler {
	return &ProductHandler{productUC: productUC}
}

func (h *ProductHandler) List(w http.ResponseWriter, r *http.Request) {
	q, err := request.ParseListProductQuery(r)
	if err != nil {
		response.WriteAppError(w, usecase.NewAppError(usecase.ErrInvalidInput))
		return
	}

	products, err := h.productUC.List(r.Context(), q.Mode, q.Category, q.Target)
	if err != nil {
		response.WriteAppError(w, err)
		return
	}

	res := presenter.ToProductResponses(products)
	response.WriteJSON(w, http.StatusOK, res)
}

func (h *ProductHandler) ListByCategory(w http.ResponseWriter, r *http.Request) {
	q, err := request.ParseListCategoryProductQuery(r)
	if err != nil {
		response.WriteAppError(w, usecase.NewAppError(usecase.ErrInvalidInput))
		return
	}

	categoryProducts, err := h.productUC.ListByCategory(r.Context(), usecaseProductQuery.ListCategoryProductsQuery{
		Category: q.Category,
		Cursor:   q.Cursor,
		Limit:    q.Limit,
		Target:   q.Target,
	})
	if err != nil {
		response.WriteAppError(w, err)
		return
	}

	res := presenter.ToCategoryProductsResponses(categoryProducts)
	response.WriteJSON(w, http.StatusOK, res)
}

func (h *ProductHandler) ExportCSV(w http.ResponseWriter, r *http.Request) {
	rows, err := h.productUC.ExportCSV(r.Context())
	if err != nil {
		response.WriteAppError(w, err)
		return
	}

	var csvBuffer bytes.Buffer
	csvWriter := csv.NewWriter(&csvBuffer)

	if err := csvWriter.Write([]string{"UUID", "Name", "Price", "CategoryName", "TargetName"}); err != nil {
		response.WriteAppError(w, usecase.NewAppErrorWithMessage(usecase.ErrInternal, err.Error()))
		return
	}

	for _, row := range rows {
		if err := csvWriter.Write([]string{
			row.UUID,
			row.Name,
			strconv.Itoa(row.Price),
			row.CategoryName,
			row.TargetName,
		}); err != nil {
			response.WriteAppError(w, usecase.NewAppErrorWithMessage(usecase.ErrInternal, err.Error()))
			return
		}
	}
	csvWriter.Flush()
	if err := csvWriter.Error(); err != nil {
		response.WriteAppError(w, usecase.NewAppErrorWithMessage(usecase.ErrInternal, err.Error()))
		return
	}

	w.Header().Set("Content-Type", "text/csv; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(csvBuffer.Bytes())
}

func (h *ProductHandler) ListCarousel(w http.ResponseWriter, r *http.Request) {
	carouselItems, err := h.productUC.ListCarousel(r.Context())
	if err != nil {
		response.WriteAppError(w, err)
		return
	}

	res := presenter.ToCarouselItemResponses(carouselItems)
	response.WriteJSON(w, http.StatusOK, res)
}

func (h *ProductHandler) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	productUUID := vars["product_uuid"]

	product, err := h.productUC.Get(r.Context(), productUUID)
	if err != nil {
		response.WriteAppError(w, err)
		return
	}
	authUser, _ := middleware.AuthenticatedUserFromContext(r.Context())
	if !product.IsActive && !authUser.IsAdmin {
		response.WriteAppError(w, usecase.NewAppError(usecase.ErrNotFound))
		return
	}

	res := presenter.ToProductResponse(product)
	response.WriteJSON(w, http.StatusOK, res)
}

func (h *ProductHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req request.CreateProductRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.WriteAppError(w, usecase.NewAppError(usecase.ErrInvalidInput))
		return
	}

	createdProductUUID, err := h.productUC.Create(r.Context(), toCreateProductInput(req))
	if err != nil {
		response.WriteAppError(w, err)
		return
	}

	res := response.CreateProductResponse{
		UUID: createdProductUUID.Value(),
	}
	response.WriteJSON(w, http.StatusOK, res)
}

func (h *ProductHandler) Duplicate(w http.ResponseWriter, r *http.Request) {
	var req request.DuplicateProductRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.WriteAppError(w, usecase.NewAppError(usecase.ErrInvalidInput))
		return
	}

	if err := h.productUC.Duplicate(r.Context(), req.URL); err != nil {
		response.WriteAppError(w, err)
		return
	}

	response.WriteSuccess(w)
}

func (h *ProductHandler) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	productUUID := vars["product_uuid"]

	var req request.UpdateProductRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.WriteAppError(w, usecase.NewAppError(usecase.ErrInvalidInput))
		return
	}

	if err := h.productUC.Update(r.Context(), productUUID, toUpdateProductInput(req)); err != nil {
		response.WriteAppError(w, err)
		return
	}

	response.WriteSuccess(w)
}

func (h *ProductHandler) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	productUUID := vars["product_uuid"]

	if err := h.productUC.Delete(r.Context(), productUUID); err != nil {
		response.WriteAppError(w, err)
		return
	}

	response.WriteSuccess(w)
}

func (h *ProductHandler) UploadCSV(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("csv")
	if err != nil {
		response.WriteAppError(w, usecase.NewAppError(usecase.ErrInvalidInput))
		return
	}
	defer func() {
		_ = file.Close()
	}()

	csvBytes, err := io.ReadAll(io.LimitReader(file, maxProductCSVSize+1))
	if err != nil {
		response.WriteAppError(w, usecase.NewAppErrorWithMessage(usecase.ErrInternal, err.Error()))
		return
	}
	if len(csvBytes) == 0 || len(csvBytes) > maxProductCSVSize {
		response.WriteAppError(w, usecase.NewAppError(usecase.ErrInvalidInput))
		return
	}

	rows, err := request.ParseProductCSV(bytes.NewReader(csvBytes))
	if err != nil {
		response.WriteAppError(w, usecase.NewAppErrorWithMessage(usecase.ErrInvalidInput, err.Error()))
		return
	}

	inputRows := make([]usecaseProduct.ProductCSVInputRow, 0, len(rows))
	for _, row := range rows {
		inputRows = append(inputRows, usecaseProduct.ProductCSVInputRow{
			UUID:         row.UUID,
			Name:         row.Name,
			Price:        row.Price,
			CategoryName: row.CategoryName,
			TargetName:   row.TargetName,
		})
	}

	if err := h.productUC.UploadCSV(r.Context(), inputRows); err != nil {
		response.WriteAppError(w, err)
		return
	}

	response.WriteSuccess(w)
}

func toCreateProductInput(req request.CreateProductRequest) usecaseProduct.CreateProductInput {
	tagUUIDs := make([]string, 0, len(req.Tags))
	for _, tag := range req.Tags {
		tagUUIDs = append(tagUUIDs, tag.UUID)
	}

	siteDetails := make([]usecaseProduct.SiteDetailInput, 0, len(req.SiteDetails))
	for _, siteDetail := range req.SiteDetails {
		siteDetails = append(siteDetails, usecaseProduct.SiteDetailInput{
			SalesSiteUUID: siteDetail.SalesSite.UUID,
			DetailURL:     siteDetail.DetailURL,
		})
	}

	return usecaseProduct.CreateProductInput{
		Name:         req.Name,
		Description:  req.Description,
		Price:        req.Price,
		IsActive:     req.IsActive,
		IsRecommend:  req.IsRecommend,
		CategoryUUID: req.Category.UUID,
		TargetUUID:   req.Target.UUID,
		TagUUIDs:     tagUUIDs,
		SiteDetails:  siteDetails,
	}
}

func toUpdateProductInput(req request.UpdateProductRequest) usecaseProduct.UpdateProductInput {
	tagUUIDs := make([]string, 0, len(req.Tags))
	for _, tag := range req.Tags {
		tagUUIDs = append(tagUUIDs, tag.UUID)
	}

	siteDetails := make([]usecaseProduct.SiteDetailInput, 0, len(req.SiteDetails))
	for _, siteDetail := range req.SiteDetails {
		siteDetails = append(siteDetails, usecaseProduct.SiteDetailInput{
			SalesSiteUUID: siteDetail.SalesSite.UUID,
			DetailURL:     siteDetail.DetailURL,
		})
	}

	productImages := make([]usecaseProduct.ProductImageUpdateInput, 0, len(req.ProductImages))
	for _, image := range req.ProductImages {
		productImages = append(productImages, usecaseProduct.ProductImageUpdateInput{
			UUID:         image.UUID,
			DisplayOrder: image.DisplayOrder,
		})
	}

	return usecaseProduct.UpdateProductInput{
		Name:          req.Name,
		Description:   req.Description,
		Price:         req.Price,
		IsActive:      req.IsActive,
		IsRecommend:   req.IsRecommend,
		CategoryUUID:  req.Category.UUID,
		TargetUUID:    req.Target.UUID,
		TagUUIDs:      tagUUIDs,
		SiteDetails:   siteDetails,
		ProductImages: productImages,
	}
}
