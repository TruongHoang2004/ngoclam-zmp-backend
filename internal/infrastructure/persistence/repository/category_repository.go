package repository

import (
	"github.com/TruongHoang2004/ngoclam-zmp-backend/internal/domain/entity"
	"github.com/TruongHoang2004/ngoclam-zmp-backend/internal/infrastructure/persistence/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type CategoryRepositoryImpl struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) entity.CategoryRepository {
	return &CategoryRepositoryImpl{db: db}
}

func (c *CategoryRepositoryImpl) Create(category *entity.Category) (*entity.Category, error) {
	categoryModel := model.MapCategoryToModel(category)

	tx := c.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return nil, err
	}

	// tránh GORM auto-insert ImageRelated
	categoryModel.ImageRelated = nil

	// create category và trả lại record đầy đủ
	if err := tx.Clauses(clause.Returning{}).Create(&categoryModel).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	// nếu có Image thì tạo ImageRelated
	if category.Image.ID != 0 {
		imageRel := model.ImageRelated{
			EntityType: model.EntityTypeCategory,
			EntityID:   categoryModel.ID,
			ImageID:    category.Image.ID,
		}

		if err := tx.Clauses(clause.Returning{}).Create(&imageRel).Error; err != nil {
			tx.Rollback()
			return nil, err
		}

		categoryModel.ImageRelated = &imageRel
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	return categoryModel.ToDomain(nil), nil
}

// FindByID returns category by ID with optional image
func (c *CategoryRepositoryImpl) FindByID(id uint) (*entity.Category, error) {
	var categoryModel model.Category
	if err := c.db.Preload("ImageRelated").First(&categoryModel, id).Error; err != nil {
		return nil, err
	}

	var imageModel *model.Image
	if categoryModel.ImageRelated != nil && categoryModel.ImageRelated.ID != 0 {
		var img model.Image
		if err := c.db.First(&img, categoryModel.ImageRelated.ImageID).Error; err == nil {
			imageModel = &img
		}
	}

	category := categoryModel.ToDomain(imageModel)
	return category, nil
}

// FindAll returns all categories with optional images
func (c *CategoryRepositoryImpl) FindAll() ([]*entity.Category, error) {
	var categoryModels []model.Category
	if err := c.db.Preload("ImageRelated").Find(&categoryModels).Error; err != nil {
		return nil, err
	}

	var categories []*entity.Category
	for _, categoryModel := range categoryModels {
		var imageModel *model.Image
		if categoryModel.ImageRelated != nil && categoryModel.ImageRelated.ID != 0 {
			var img model.Image
			if err := c.db.First(&img, categoryModel.ImageRelated.ImageID).Error; err == nil {
				imageModel = &img
			}
		}
		categories = append(categories, categoryModel.ToDomain(imageModel))
	}
	return categories, nil
}

// Update modifies an existing category
func (c *CategoryRepositoryImpl) Update(category *entity.Category) error {
	categoryModel := model.MapCategoryToModel(category)

	return c.db.Save(categoryModel).Error
}

// Delete removes a category by ID
func (c *CategoryRepositoryImpl) Delete(id uint) error {
	return c.db.Delete(&model.Category{}, id).Error
}
