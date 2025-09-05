package entity

// Image represents an image associated with a product
type Image struct {
	ID        uint   // Unique identifier for the image
	ImageURL  string // URL path to the image
	IsPrimary bool   // Flag to indicate if this is the primary/main image for the product
}

// NewImage creates a new Image entity
func NewImage(productID uint, imageURL string, isPrimary bool) *Image {
	return &Image{
		ImageURL:  imageURL,
		IsPrimary: isPrimary,
	}
}

// SetAsPrimary marks this image as the primary image
func (i *Image) SetAsPrimary() {
	i.IsPrimary = true
}

// UnsetAsPrimary marks this image as not the primary image
func (i *Image) UnsetAsPrimary() {
	i.IsPrimary = false
}
