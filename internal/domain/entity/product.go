package entity

import (
	"time"
)

// Product represents a product entity in the domain
type Product struct {
	ID          uint
	Name        string
	Description string
	Price       int64
	CategoryID  uint
	Category    Category
	Images      []Image
	Variants    []Variant
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// NewProduct creates a new product instance
func NewProduct(name, description string, price int64, categoryID uint) *Product {
	now := time.Now()
	return &Product{
		Name:        name,
		Description: description,
		Price:       price,
		CategoryID:  categoryID,
		Images:      []Image{},
		CreatedAt:   now,
		UpdatedAt:   now,
	}
}

// UpdateInfo updates the product information
func (p *Product) UpdateInfo(name string, description string, price int64, categoryID uint) {
	p.Name = name
	p.Description = description
	p.Price = price
	p.CategoryID = categoryID
	p.UpdatedAt = time.Now()
}

type ProductRepository interface {
	Create(product Product) (*Product, error)
	FindByID(id uint) (*Product, error)
	FindAll() ([]*Product, error)
	Update(product Product) (*Product, error)
	Delete(id uint) error
	FindByCategoryID(categoryID uint) ([]*Product, error)
}
