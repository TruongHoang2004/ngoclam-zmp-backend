package entity

import (
	"context"
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
	Create(ctx context.Context, product Product) (*Product, error)
	FindByID(ctx context.Context, id uint) (*Product, error)
	FindAll(ctx context.Context) ([]*Product, error)
	Update(ctx context.Context, product Product) (*Product, error)
	Delete(ctx context.Context, id uint) error
	FindByCategoryID(ctx context.Context, categoryID uint) ([]*Product, error)
}
