package application

import (
	"context"

	"github.com/TruongHoang2004/ngoclam-zmp-backend/internal/domain/entity"
)

type ProductService struct {
	ProductRepository entity.ProductRepository
}

func NewProductService(productRepo entity.ProductRepository) *ProductService {
	return &ProductService{
		ProductRepository: productRepo,
	}
}

func (s *ProductService) CreateProduct(ctx context.Context, product entity.Product) (*entity.Product, error) {
	return s.ProductRepository.Create(ctx, product)
}

func (s *ProductService) GetProductByID(ctx context.Context, id uint) (*entity.Product, error) {
	return s.ProductRepository.FindByID(ctx, id)
}

func (s *ProductService) GetAllProducts(ctx context.Context) ([]*entity.Product, error) {
	return s.ProductRepository.FindAll(ctx)
}

func (s *ProductService) UpdateProduct(ctx context.Context, product entity.Product) (*entity.Product, error) {
	return s.ProductRepository.Update(ctx, product)
}

func (s *ProductService) DeleteProduct(ctx context.Context, id uint) error {
	return s.ProductRepository.Delete(ctx, id)
}

func (s *ProductService) GetProductsByCategoryID(ctx context.Context, categoryID uint) ([]*entity.Product, error) {
	return s.ProductRepository.FindByCategoryID(ctx, categoryID)
}
