package database

import (
	"ngoclam-zmp-be/internal/infrastructure/persistence/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB() {
	dsn := "host=localhost user=postgres password=truonghoang2004 dbname=ngoclam port=5432 sslmode=disable TimeZone=Asia/Ho_Chi_Minh"
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}

	// Auto Migrate
	db.AutoMigrate(&model.Category{}, &model.Product{}, &model.Image{})
}

func GetDB() *gorm.DB {
	return db
}

func CloseDB() error {
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}
