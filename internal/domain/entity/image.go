package entity

import (
	"mime/multipart"
)

// Image represents an image associated with a product
type Image struct {
	ID   uint   `json:"id" `  // Unique identifier for the image
	Path string `json:"path"` // URL path to the image
}

// NewImage creates a new Image entity
func NewImage(imagePath string, isPrimary bool) *Image {
	return &Image{
		ID:   0,
		Path: imagePath,
	}
}

type ImageRepository interface {
	SaveFile(file *multipart.FileHeader) (*Image, error)
	FindByID(id uint) (*Image, error)
	FindAll() ([]*Image, error)
	Delete(id uint) error
}
