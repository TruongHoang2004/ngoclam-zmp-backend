package application

import (
	"context"
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

func (s *CategoryService) CreateCategory(ctx context.Context, category *entity.Category) (*entity.Category, error) {

	createdCategory, err := s.categoryRepo.Create(ctx, category)
	if err != nil {
		return nil, fmt.Errorf("cannot create category: %w", err)
	}

	return createdCategory, nil
}

func (s *CategoryService) GetCategoryByID(ctx context.Context, id uint) (*entity.Category, error) {
	category, err := s.categoryRepo.FindByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("cannot find category by ID: %w", err)
	}
	return category, nil
}

func (s *CategoryService) GetAllCategories(ctx context.Context) ([]*entity.Category, error) {
	categories, err := s.categoryRepo.FindAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("cannot find all categories: %w", err)
	}
	return categories, nil
}

func (s *CategoryService) UpdateCategory(ctx context.Context, id uint, name, description string) error {
	category, err := s.categoryRepo.FindByID(ctx, id)
	if err != nil {
		return fmt.Errorf("cannot find category by ID: %w", err)
	}

	category.UpdateInfo(name, description)

	if err := s.categoryRepo.Update(ctx, category); err != nil {
		return fmt.Errorf("cannot update category: %w", err)
	}
	return nil
}

func (s *CategoryService) DeleteCategory(ctx context.Context, id uint) error {
	if err := s.categoryRepo.Delete(ctx, id); err != nil {
		return fmt.Errorf("cannot delete category: %w", err)
	}
	return nil
}
