package model

type Image struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	ProductID uint   `json:"product_id" gorm:"not null"`
	ImageURL  string `json:"image_url" gorm:"not null;size:255"`
	IsPrimary bool   `json:"is_primary" gorm:"default:false"`
}
