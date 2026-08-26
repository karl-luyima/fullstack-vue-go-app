package main

import (
	"log"
	"net/http"

	"github.com/karl-luyima/fullstack-vue-go-app/backend/config"
	"github.com/karl-luyima/fullstack-vue-go-app/backend/database"
	"github.com/gin-gonic/gin"
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

	router := gin.Default()

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	router.Run(":" + cfg.AppPort)
}