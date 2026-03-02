package presenter

import (
	"github.com/tokushun109/tku/backend/internal/interface/http/response"
	usecaseProductQuery "github.com/tokushun109/tku/backend/internal/usecase/product/query"
)

func ToProductResponse(product *usecaseProductQuery.Product) *response.ProductResponse {
	if product == nil {
		return &response.ProductResponse{}
	}

	tags := make([]response.ProductClassificationResponse, 0, len(product.Tags))
	for _, tag := range product.Tags {
		tags = append(tags, response.ProductClassificationResponse{UUID: tag.UUID, Name: tag.Name})
	}

	images := make([]response.ProductImageResponse, 0, len(product.ProductImages))
	for _, image := range product.ProductImages {
		images = append(images, response.ProductImageResponse{
			UUID:         image.UUID,
			Name:         image.Name,
			DisplayOrder: image.DisplayOrder,
			APIPath:      image.APIPath,
		})
	}

	siteDetails := make([]response.ProductSiteDetailResponse, 0, len(product.SiteDetails))
	for _, detail := range product.SiteDetails {
		siteDetails = append(siteDetails, response.ProductSiteDetailResponse{
			UUID:      detail.UUID,
			DetailURL: detail.DetailURL,
			SalesSite: response.ProductSalesSiteResponse{
				UUID: detail.SalesSite.UUID,
				Name: detail.SalesSite.Name,
			},
		})
	}

	return &response.ProductResponse{
		UUID:        product.UUID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		IsRecommend: product.IsRecommend,
		IsActive:    product.IsActive,
		Category: response.ProductClassificationResponse{
			UUID: product.Category.UUID,
			Name: product.Category.Name,
		},
		Target: response.ProductClassificationResponse{
			UUID: product.Target.UUID,
			Name: product.Target.Name,
		},
		Tags:          tags,
		ProductImages: images,
		SiteDetails:   siteDetails,
	}
}

func ToProductResponses(products []*usecaseProductQuery.Product) []*response.ProductResponse {
	result := make([]*response.ProductResponse, 0, len(products))
	for _, product := range products {
		result = append(result, ToProductResponse(product))
	}
	return result
}

func ToCategoryProductsResponse(categoryProducts *usecaseProductQuery.CategoryProducts) *response.CategoryProductsResponse {
	if categoryProducts == nil {
		return &response.CategoryProductsResponse{
			Category: response.ProductClassificationResponse{},
			PageInfo: response.CursorPageInfoResponse{},
			Products: []*response.ProductResponse{},
		}
	}

	return &response.CategoryProductsResponse{
		Category: response.ProductClassificationResponse{
			UUID: categoryProducts.Category.UUID,
			Name: categoryProducts.Category.Name,
		},
		PageInfo: response.CursorPageInfoResponse{
			HasMore:    categoryProducts.PageInfo.HasMore,
			NextCursor: categoryProducts.PageInfo.NextCursor,
		},
		Products: ToProductResponses(categoryProducts.Products),
	}
}

func ToCategoryProductsResponses(categoryProductsList []*usecaseProductQuery.CategoryProducts) []*response.CategoryProductsResponse {
	result := make([]*response.CategoryProductsResponse, 0, len(categoryProductsList))
	for _, categoryProducts := range categoryProductsList {
		result = append(result, ToCategoryProductsResponse(categoryProducts))
	}
	return result
}

func ToCarouselItemResponse(item *usecaseProductQuery.CarouselItem) *response.CarouselItemResponse {
	if item == nil {
		return &response.CarouselItemResponse{
			Product: &response.ProductResponse{},
		}
	}

	return &response.CarouselItemResponse{
		Product: ToProductResponse(item.Product),
		APIPath: item.APIPath,
	}
}

func ToCarouselItemResponses(items []*usecaseProductQuery.CarouselItem) []*response.CarouselItemResponse {
	result := make([]*response.CarouselItemResponse, 0, len(items))
	for _, item := range items {
		result = append(result, ToCarouselItemResponse(item))
	}
	return result
}
