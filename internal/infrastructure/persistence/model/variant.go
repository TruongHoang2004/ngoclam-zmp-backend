package model

import (
	"time"

	"github.com/TruongHoang2004/ngoclam-zmp-backend/internal/domain/entity"
	"gorm.io/gorm"
)

type VariantModel struct {
	ID        uint   `gorm:"primaryKey;autoIncrement"`
	ProductID uint   `gorm:"index;not null"`
	SKU       string `gorm:"type:varchar(100);not null;unique"`
	Price     int64  `gorm:"not null"`

	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func MapVariantToModel(variant *entity.Variant) *VariantModel {
	return &VariantModel{
		ID:        variant.ID,
		ProductID: variant.ProductID,
		SKU:       variant.SKU,
		Price:     variant.Price,
		CreatedAt: variant.CreatedAt,
		UpdatedAt: variant.UpdatedAt,
	}
}

func (v *VariantModel) ToDomain() *entity.Variant {
	return &entity.Variant{
		ID:        v.ID,
		ProductID: v.ProductID,
		SKU:       v.SKU,
		Price:     v.Price,
		CreatedAt: v.CreatedAt,
		UpdatedAt: v.UpdatedAt,
	}
}
