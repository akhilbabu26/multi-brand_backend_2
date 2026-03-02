package models

import "gorm.io/gorm"

// represents a database table structure of users
type User struct{
	gorm.Model
	Name string
	Email string `gorm:"unique"`
	Password string
	Role string
	IsBlocked bool
}