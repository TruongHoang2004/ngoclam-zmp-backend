package model

import "github.com/TruongHoang2004/ngoclam-zmp-backend/internal/domain/entity"

type Image struct {
	ID       uint   `gorm:"primaryKey;autoIncrement"`
	URL      string `gorm:"not null;size:512"`             // link public của ảnh trên ImageKit
	IKFileID string `gorm:"not null;size:128;uniqueIndex"` // fileId trong ImageKit
	Hash     string `gorm:"size:64;uniqueIndex;not null"`  // sha256 checksum
}

// Map sang domain entity
func (m *Image) ToDomain() *entity.Image {
	return &entity.Image{
		ID:       m.ID,
		URL:      m.URL,
		IKFileID: m.IKFileID,
		Hash:     m.Hash,
	}
}
