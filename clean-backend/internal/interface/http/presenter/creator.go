package presenter

import (
	usecaseCreator "github.com/tokushun109/tku/clean-backend/internal/usecase/creator"

	"github.com/tokushun109/tku/clean-backend/internal/interface/http/response"
)

func ToCreatorResponse(detail *usecaseCreator.CreatorDetail) *response.CreatorResponse {
	if detail == nil || detail.Creator == nil {
		return &response.CreatorResponse{}
	}

	mimeType := ""
	if detail.Creator.LogoMimeType != nil {
		mimeType = detail.Creator.LogoMimeType.String()
	}

	logoPath := ""
	if detail.Creator.LogoPath != nil {
		logoPath = detail.Creator.LogoPath.String()
	}

	return &response.CreatorResponse{
		Name:         detail.Creator.Name.String(),
		Introduction: detail.Creator.Introduction.String(),
		MimeType:     mimeType,
		Logo:         logoPath,
		APIPath:      detail.APIPath,
	}
}
