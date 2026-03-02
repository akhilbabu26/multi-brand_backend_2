package products

// create product
type CreateProductDOT struct{
	Name string `json:"name"`
	Type string `json:"type"`
	Color string `json:"color"`
	CostPrice float64 `json:"cost_price"`
	OriginalPrice float64 `json:"original_price"`
	DiscountPercentage float64 `json:"discount_percentage"`
	ImageURL string `json:"image_url"`
	Description string `json:"description"`
	Stock int `json:"stock"`
	IsActive bool `json:"is_active"`
}

// Update Product
type UpdateProductDTO struct {
	Name *string `json:"name"`
	Type *string `json:"type"`
	Color *string `json:"color"`
	CostPrice *float64 `json:"cost_price"`
	OriginalPrice *float64 `json:"original_price"`
	DiscountPercentage *float64 `json:"discount_percentage"`
	ImageURL *string `json:"image_url"`
	Description *string `json:"description"`
	Stock *int `json:"stock"`
	IsActive *bool `json:"isActive"`
}