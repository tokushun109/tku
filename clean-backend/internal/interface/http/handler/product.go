package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tokushun109/tku/clean-backend/internal/interface/http/presenter"
	"github.com/tokushun109/tku/clean-backend/internal/interface/http/request"
	"github.com/tokushun109/tku/clean-backend/internal/interface/http/response"
	"github.com/tokushun109/tku/clean-backend/internal/usecase"
	usecaseProduct "github.com/tokushun109/tku/clean-backend/internal/usecase/product"
)

const maxProductImageSize = 20 << 20 // 20MB

type ProductHandler struct {
	productUC usecaseProduct.Usecase
}

type productImageOrderParams struct {
	IsChanged bool        `json:"isChanged"`
	Order     map[int]int `json:"order"`
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

func (h *ProductHandler) GetImageBlob(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	productImageUUID := vars["product_image_uuid"]

	blob, err := h.productUC.GetProductImageBlob(r.Context(), productImageUUID)
	if err != nil {
		response.WriteAppError(w, err)
		return
	}
	defer func() {
		_ = blob.Body.Close()
	}()

	w.Header().Set("Content-Type", blob.ContentType)
	w.WriteHeader(http.StatusOK)
	_, _ = io.Copy(w, blob.Body)
}

func (h *ProductHandler) CreateImage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	productUUID := vars["product_uuid"]

	orderParams := productImageOrderParams{Order: map[int]int{}}
	orderJSON := r.FormValue("order")
	if orderJSON != "" {
		if err := json.Unmarshal([]byte(orderJSON), &orderParams); err != nil {
			response.WriteAppError(w, usecase.NewAppError(usecase.ErrInvalidInput))
			return
		}
	}

	files := make([]usecaseProduct.ProductImageUploadFile, 0)
	for i := 0; ; i++ {
		fileField := fmt.Sprintf("file%d", i)
		file, fileHeader, err := r.FormFile(fileField)
		if err != nil {
			if errors.Is(err, http.ErrMissingFile) {
				break
			}
			response.WriteAppError(w, usecase.NewAppError(usecase.ErrInvalidInput))
			return
		}

		fileBytes, err := io.ReadAll(io.LimitReader(file, maxProductImageSize+1))
		_ = file.Close()
		if err != nil {
			response.WriteAppError(w, usecase.NewAppErrorWithMessage(usecase.ErrInternal, err.Error()))
			return
		}
		if len(fileBytes) == 0 || len(fileBytes) > maxProductImageSize {
			response.WriteAppError(w, usecase.NewAppError(usecase.ErrInvalidInput))
			return
		}

		files = append(files, usecaseProduct.ProductImageUploadFile{
			Name: fileHeader.Filename,
			Data: fileBytes,
		})
	}

	if err := h.productUC.CreateProductImages(r.Context(), productUUID, files, orderParams.IsChanged, orderParams.Order); err != nil {
		response.WriteAppError(w, err)
		return
	}

	response.WriteSuccess(w)
}

func (h *ProductHandler) DeleteImage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	productUUID := vars["product_uuid"]
	productImageUUID := vars["product_image_uuid"]

	if err := h.productUC.DeleteProductImage(r.Context(), productUUID, productImageUUID); err != nil {
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
