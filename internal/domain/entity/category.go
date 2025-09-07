package entity

import "time"

// Category represents a product category in the domain
type Category struct {
	ID          uint
	Name        string
	Description string
	Image       Image

	CreatedAt time.Time
	UpdatedAt time.Time
	// Products reference is optional in the domain model depending on your needs
}

type CategoryRepository interface {
	Create(category *Category) (*Category, error)
	FindByID(id uint) (*Category, error)
	FindAll() ([]*Category, error)
	Update(category *Category) error
	Delete(id uint) error
}

// NewCategory creates a new category instance
func NewCategory(name string, description string) *Category {
	now := time.Now()
	return &Category{
		Name:        name,
		Description: description,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
}

// UpdateInfo updates the category information
func (c *Category) UpdateInfo(name string, description string) {
	c.Name = name
	c.Description = description
	c.UpdatedAt = time.Now()
}
