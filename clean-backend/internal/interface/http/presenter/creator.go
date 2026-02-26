package presenter

import (
	domainVO "github.com/tokushun109/tku/clean-backend/internal/domain/vo"
	usecaseCreator "github.com/tokushun109/tku/clean-backend/internal/usecase/creator"

	"github.com/tokushun109/tku/clean-backend/internal/interface/http/response"
)

func ToCreatorResponse(detail *usecaseCreator.CreatorDetail) *response.CreatorResponse {
	if detail == nil || detail.Creator == nil {
		return &response.CreatorResponse{}
	}

	mimeType := ""
	if detail.Creator.LogoMimeType() != nil {
		mimeType = detail.Creator.LogoMimeType().Value()
	}

	logoPath := ""
	if detail.Creator.LogoPath() != nil {
		logoPath = detail.Creator.LogoPath().Value()
	}

	return &response.CreatorResponse{
		Name:         detail.Creator.Name().Value(),
		Introduction: domainVO.ToValueOrEmpty(detail.Creator.Introduction()),
		MimeType:     mimeType,
		Logo:         logoPath,
		APIPath:      detail.APIPath,
	}
}
