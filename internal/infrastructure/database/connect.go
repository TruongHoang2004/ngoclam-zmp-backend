package database

import (
	"fmt"
	"log"

	"github.com/TruongHoang2004/ngoclam-zmp-backend/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDatabase(config *config.Config) (*gorm.DB, error) {
	log.Printf("🔌 Connecting to database %s", config.DBUrl)
	db, err := gorm.Open(postgres.Open(config.DBUrl), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect database: %w", err)
	}

	log.Println("✅ Database connected")
	return db.Debug(), nil
}
