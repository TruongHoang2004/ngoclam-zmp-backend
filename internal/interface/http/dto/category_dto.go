package dto

import "github.com/TruongHoang2004/ngoclam-zmp-backend/internal/domain/entity"

type CreateCategoryRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	ImageID     uint   `json:"image_id"`
}

func (r *CreateCategoryRequest) ToDomain() *entity.Category {
	var image entity.Image
	if r.ImageID != 0 {
		image = entity.Image{
			ID: r.ImageID,
		}
	}
	return &entity.Category{
		Name:        r.Name,
		Description: r.Description,
		Image:       image,
	}
}

type UpdateCategoryRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	ImageID     uint   `json:"image_id"`
}

type CategoryResponse struct {
	ID          uint           `json:"id"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Image       *ImageResponse `json:"image,omitempty"`
}

func NewCategoryResponse(category entity.Category) CategoryResponse {
	return CategoryResponse{
		ID:          category.ID,
		Name:        category.Name,
		Description: category.Description,
		Image:       NewImageResponse(&category.Image),
	}
}
