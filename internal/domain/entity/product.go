package entity

import "time"

// Product represents a product entity in the domain
type Product struct {
	ID          uint
	Name        string
	Description string
	Price       float64
	CategoryID  uint
	Category    Category
	Images      []Image
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// NewProduct creates a new product instance
func NewProduct(name string, description string, price float64, categoryID uint) *Product {
	now := time.Now()
	return &Product{
		Name:        name,
		Description: description,
		Price:       price,
		CategoryID:  categoryID,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
}

// UpdateInfo updates the product information
func (p *Product) UpdateInfo(name string, description string, price float64, categoryID uint) {
	p.Name = name
	p.Description = description
	p.Price = price
	p.CategoryID = categoryID
	p.UpdatedAt = time.Now()
}
