package models

import "gorm.io/gorm"

// represents a database table structure of products
type Product struct{
	gorm.Model

	Name string 
	Type string 
	Color string 
	CostPrice float64 
	OriginalPrice float64 
	SalePrice float64 
	DiscountPercentage float64 
	ImageURL string 
	Description string 
	Stock int 
	IsActive bool 
}


// GORM HOOK - AUTO SALE PRICE (trigger)

// CalculateSalePrice calculates discounted price
func (p *Product) CalculateSalePrice() {

	// safety check
	if p.DiscountPercentage < 0 {
		p.DiscountPercentage = 0
	}

	if p.DiscountPercentage > 100 {
		p.DiscountPercentage = 100
	}

	p.SalePrice =p.OriginalPrice - (p.OriginalPrice * p.DiscountPercentage / 100)
}

// BeforeCreate runs automatically before inserting into db.Create(&product)
func (p *Product) BeforeCreate(tx *gorm.DB) (err error) {
	p.CalculateSalePrice()
	return nil
}

// BeforeUpdate runs automatically before updating DB
func (p *Product) BeforeUpdate(tx *gorm.DB) (err error) {
	p.CalculateSalePrice()
	return nil
}

