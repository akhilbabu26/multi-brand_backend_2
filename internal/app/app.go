package app

import (
	"log"

	"github.com/akhilbabu26/multi-brand_backend_2/config"
	"github.com/akhilbabu26/multi-brand_backend_2/internal/models"
)

// Init initializes config, database and migrations
func Init() {

	// load config
	cfg, err := config.LoadConfig("config/config.yaml")
	if err != nil {
		log.Fatal("failed to load config:", err)
	}

	config.AppConfig = cfg

	// connect database
	config.ConnectDB(cfg)

	// run migrations
	if err := config.DB.AutoMigrate(&models.User{}); err != nil {
		log.Fatal("migration failed:", err)
	}
}