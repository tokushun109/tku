package presenter

import (
	"github.com/tokushun109/tku/clean-backend/internal/interface/http/response"
	usecaseProductQuery "github.com/tokushun109/tku/clean-backend/internal/usecase/product/query"
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
			UUID:    image.UUID,
			Name:    image.Name,
			Order:   image.Order,
			APIPath: image.APIPath,
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
