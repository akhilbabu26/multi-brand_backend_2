// package config

// import (
// 	"fmt"
// 	"gorm.io/driver/postgres"
// 	"gorm.io/gorm"
// )

// var DB *gorm.DB

// func ConnectDB(cfg *Config) {
// 	dsn := fmt.Sprintf(
// 		"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
// 		cfg.DB.Host,
// 		cfg.DB.User,
// 		cfg.DB.Password,
// 		cfg.DB.DBName,
// 		cfg.DB.Port,
// 		cfg.DB.SSLMode,
// 	)

// 	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

// 	if err != nil{
// 		panic("DB connection failed")
// 	}

// 	DB = db
// }

package config

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//
// ======================================================
// DATABASE INSTANCE (GLOBAL)
// ======================================================
//

// DB is the global database connection used across the app
var DB *gorm.DB

//
// ======================================================
// DATABASE CONNECTION
// ======================================================
//

// ConnectDB creates and initializes the database connection
func ConnectDB(cfg *Config) {

	// build postgres DSN (Data Source Name)
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
		cfg.DB.Host,
		cfg.DB.User,
		cfg.DB.Password,
		cfg.DB.DBName,
		cfg.DB.Port,
		cfg.DB.SSLMode,
	)

	// open database connection using GORM
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("DB connection failed: %v", err))
	}

	// assign global DB instance
	DB = db
}