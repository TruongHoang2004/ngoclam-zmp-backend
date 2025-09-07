package dto

import "github.com/TruongHoang2004/ngoclam-zmp-backend/internal/domain/entity"

type CreateImageRequest struct {
	URL string `json:"url" binding:"required"`
}

type UpdateImageRequest struct {
	URL string `json:"url" binding:"required"`
}

type ImageResponse struct {
	ID  uint   `json:"id"`
	URL string `json:"url"`
}

func NewImageResponse(image *entity.Image) *ImageResponse {
	return &ImageResponse{
		ID:  uint(image.ID),
		URL: image.Path,
	}
}
