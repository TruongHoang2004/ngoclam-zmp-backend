package application

import "github.com/TruongHoang2004/ngoclam-zmp-backend/internal/domain/entity"

type ProductService struct {
	ProductRepository entity.ProductRepository
}

func NewProductService(productRepo entity.ProductRepository) *ProductService {
	return &ProductService{
		ProductRepository: productRepo,
	}
}

func (s *ProductService) CreateProduct(product entity.Product) (*entity.Product, error) {
	return s.ProductRepository.Create(product)
}

func (s *ProductService) GetProductByID(id uint) (*entity.Product, error) {
	return s.ProductRepository.FindByID(id)
}

func (s *ProductService) GetAllProducts() ([]*entity.Product, error) {
	return s.ProductRepository.FindAll()
}

func (s *ProductService) UpdateProduct(product entity.Product) (*entity.Product, error) {
	return s.ProductRepository.Update(product)
}

func (s *ProductService) DeleteProduct(id uint) error {
	return s.ProductRepository.Delete(id)
}

func (s *ProductService) GetProductsByCategoryID(categoryID uint) ([]*entity.Product, error) {
	return s.ProductRepository.FindByCategoryID(categoryID)
}
