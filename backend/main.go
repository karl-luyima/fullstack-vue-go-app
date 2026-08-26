package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/karl-luyima/fullstack-vue-go-app/backend/config"
	"github.com/karl-luyima/fullstack-vue-go-app/backend/controllers"
	"github.com/karl-luyima/fullstack-vue-go-app/backend/database"
	"github.com/karl-luyima/fullstack-vue-go-app/backend/repositories"
	"github.com/karl-luyima/fullstack-vue-go-app/backend/services"
	"github.com/karl-luyima/fullstack-vue-go-app/backend/utils"
)

func main() {
	cfg := config.Load()

	db, err := database.Connect(cfg)
	if err != nil {
		log.Fatal(err)
	}
	if err := database.Migrate(db); err != nil {
		log.Fatal(err)
	}

	jwtManager := utils.NewJWTManager(cfg.JWTSecret, 24*time.Hour)

	userRepo := repositories.NewUserRepository(db)
	authService := services.NewAuthService(userRepo, jwtManager)
	authController := controllers.NewAuthController(authService)

	router := gin.Default()

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	api := router.Group("/api")
	auth := api.Group("/auth")
	auth.POST("/register", authController.Register)
	auth.POST("/login", authController.Login)

	router.Run(":" + cfg.AppPort)
}