package models

import "gorm.io/gorm"

// represents a database table structure of card
type Cart struct{
	gorm.Model

	UserID uint
	ProductID uint
	Quantity int
}