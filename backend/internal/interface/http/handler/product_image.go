package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tokushun109/tku/backend/internal/interface/http/response"
	"github.com/tokushun109/tku/backend/internal/usecase"
	usecaseProduct "github.com/tokushun109/tku/backend/internal/usecase/product"
)

const maxProductImageSize = 20 << 20 // 20MB

type ProductImageHandler struct {
	productUC usecaseProduct.Usecase
}

type productImageDisplayOrderParams struct {
	IsChanged    bool        `json:"isChanged"`
	DisplayOrder map[int]int `json:"displayOrder"`
}

func NewProductImageHandler(productUC usecaseProduct.Usecase) *ProductImageHandler {
	return &ProductImageHandler{productUC: productUC}
}

func (h *ProductImageHandler) GetBlob(w http.ResponseWriter, r *http.Request) {
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

func (h *ProductImageHandler) Create(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	productUUID := vars["product_uuid"]

	displayOrderParams := productImageDisplayOrderParams{DisplayOrder: map[int]int{}}
	displayOrderJSON := r.FormValue("displayOrder")
	if displayOrderJSON != "" {
		if err := json.Unmarshal([]byte(displayOrderJSON), &displayOrderParams); err != nil {
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

	if err := h.productUC.CreateProductImages(r.Context(), productUUID, files, displayOrderParams.IsChanged, displayOrderParams.DisplayOrder); err != nil {
		response.WriteAppError(w, err)
		return
	}

	response.WriteSuccess(w)
}

func (h *ProductImageHandler) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	productUUID := vars["product_uuid"]
	productImageUUID := vars["product_image_uuid"]

	if err := h.productUC.DeleteProductImage(r.Context(), productUUID, productImageUUID); err != nil {
		response.WriteAppError(w, err)
		return
	}

	response.WriteSuccess(w)
}
