package main

import (
	"log"

	"github.com/akhilbabu26/multi-brand_backend_2/config"
	"github.com/akhilbabu26/multi-brand_backend_2/internal/auth"
	"github.com/akhilbabu26/multi-brand_backend_2/internal/models"

	"github.com/gin-gonic/gin"
)

func main(){
	cfg, err := config.LoadConfig("config/config.yaml")
	if err != nil{
		log.Fatal("failed to load config:", err)
		return
	}

	config.AppConfig = cfg
	config.ConnectDB(cfg)

	if err := config.DB.AutoMigrate(&models.User{}); err != nil {
		log.Fatal("migration failed:", err)
	}

	r := gin.Default()

	authRoute := r.Group("/auth")
	{
		authRoute.POST("/signup", auth.Signup)
		authRoute.POST("/login", auth.Login)
		authRoute.POST("/refresh", auth.RefreshToken)
	}

	r.Run(":8080")
}