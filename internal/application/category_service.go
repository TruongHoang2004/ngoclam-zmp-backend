package application

import (
	"fmt"

	"github.com/TruongHoang2004/ngoclam-zmp-backend/internal/domain/entity"
)

type CategoryService struct {
	categoryRepo entity.CategoryRepository
	// imageRepo    entity.ImageRepository
}

func NewCategoryService(categoryRepo entity.CategoryRepository) *CategoryService {
	return &CategoryService{
		categoryRepo: categoryRepo,
		// imageRepo:    imageRepo,
	}
}

func (s *CategoryService) CreateCategory(category *entity.Category) (*entity.Category, error) {

	createdCategory, err := s.categoryRepo.Create(category)
	if err != nil {
		return nil, fmt.Errorf("cannot create category: %w", err)
	}

	return createdCategory, nil
}

func (s *CategoryService) GetCategoryByID(id uint) (*entity.Category, error) {
	category, err := s.categoryRepo.FindByID(id)
	if err != nil {
		return nil, fmt.Errorf("cannot find category by ID: %w", err)
	}
	return category, nil
}

func (s *CategoryService) GetAllCategories() ([]*entity.Category, error) {
	categories, err := s.categoryRepo.FindAll()
	if err != nil {
		return nil, fmt.Errorf("cannot find all categories: %w", err)
	}
	return categories, nil
}

func (s *CategoryService) UpdateCategory(id uint, name, description string) error {
	category, err := s.categoryRepo.FindByID(id)
	if err != nil {
		return fmt.Errorf("cannot find category by ID: %w", err)
	}

	category.UpdateInfo(name, description)

	if err := s.categoryRepo.Update(category); err != nil {
		return fmt.Errorf("cannot update category: %w", err)
	}
	return nil
}

func (s *CategoryService) DeleteCategory(id uint) error {
	if err := s.categoryRepo.Delete(id); err != nil {
		return fmt.Errorf("cannot delete category: %w", err)
	}
	return nil
}
