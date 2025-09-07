package application

import (
	"fmt"
	"mime/multipart"

	"github.com/TruongHoang2004/ngoclam-zmp-backend/internal/domain/entity"
)

type ImageService struct {
	imageRepo entity.ImageRepository
}

func NewImageService(imageRepo entity.ImageRepository) *ImageService {
	return &ImageService{
		imageRepo: imageRepo,
	}
}

func (s *ImageService) UploadImage(fileHeader *multipart.FileHeader) (*entity.Image, error) {
	// Validate file size (e.g., max 5MB)
	const maxFileSize = 5 << 20 // 5MB
	if fileHeader.Size > maxFileSize {
		return nil, NewUnsupportedMediaTypeError(fmt.Sprintf("file size exceeds the limit of %d bytes", maxFileSize))
	}
	image, err := s.imageRepo.SaveFile(fileHeader)
	if err != nil {
		if err.Error() == "ERROR: duplicate key value violates unique constraint \"idx_images_hash\" (SQLSTATE 23505)" {
			return nil, NewConflictError("image already exists")
		}
		return nil, NewInternalServerError(fmt.Sprintf("cannot save image: %v", err))
	}
	return image, nil
}

func (s *ImageService) GetImageByID(id uint) (*entity.Image, error) {
	image, err := s.imageRepo.FindByID(id)
	if err != nil {
		return nil, fmt.Errorf("cannot find image by ID: %w", err)
	}
	return image, nil
}

func (s *ImageService) ListImages() ([]*entity.Image, error) {
	images, err := s.imageRepo.FindAll()
	if err != nil {
		return nil, fmt.Errorf("cannot list images: %w", err)
	}
	return images, nil
}

func (s *ImageService) DeleteImage(id uint) error {
	if err := s.imageRepo.Delete(id); err != nil {
		return fmt.Errorf("cannot delete image: %w", err)
	}
	return nil
}
