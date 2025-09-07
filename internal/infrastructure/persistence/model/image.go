package model

import "github.com/TruongHoang2004/ngoclam-zmp-backend/internal/domain/entity"

type Image struct {
	ID   uint   `gorm:"primaryKey;autoIncrement"`
	Path string `gorm:"not null;size:255"`
	Hash string `gorm:"size:64;uniqueIndex;not null default:''"`
}

func (Image) TableName() string {
	return "images"
}

func NewImage(path string) *Image {
	return &Image{
		Path: path,
		Hash: "",
	}
}

func (img *Image) ToDomain() *entity.Image {
	return &entity.Image{
		ID:   img.ID,
		Path: "./uploads/" + img.Path,
	}
}
func MapImageToModel(image *entity.Image) *Image {
	return &Image{
		ID:   image.ID,
		Path: image.Path,
	}
}
