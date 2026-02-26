package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tokushun109/tku/clean-backend/internal/interface/http/presenter"
	"github.com/tokushun109/tku/clean-backend/internal/interface/http/request"
	"github.com/tokushun109/tku/clean-backend/internal/interface/http/response"
	"github.com/tokushun109/tku/clean-backend/internal/usecase"
	usecaseProduct "github.com/tokushun109/tku/clean-backend/internal/usecase/product"
)

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

	categoryProducts, err := h.productUC.ListByCategory(r.Context(), q.Mode, q.Category, q.Target)
	if err != nil {
		response.WriteAppError(w, err)
		return
	}

	res := presenter.ToCategoryProductsResponses(categoryProducts)
	response.WriteJSON(w, http.StatusOK, res)
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

	res := presenter.ToProductResponse(product)
	response.WriteJSON(w, http.StatusOK, res)
}

func (h *ProductHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req request.CreateProductRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.WriteAppError(w, usecase.NewAppError(usecase.ErrInvalidInput))
		return
	}

	createdProduct, err := h.productUC.Create(r.Context(), toCreateProductInput(req))
	if err != nil {
		response.WriteAppError(w, err)
		return
	}

	res := presenter.ToProductResponse(createdProduct)
	response.WriteJSON(w, http.StatusOK, res)
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
			UUID:  image.UUID,
			Order: image.Order,
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
