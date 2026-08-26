package database

import (
	"fmt"

	"github.com/karl-luyima/fullstack-vue-go-app/backend/config"
	"github.com/karl-luyima/fullstack-vue-go-app/backend/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Connect opens a connection to Postgres using the given config.
func Connect(cfg *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBSSLMode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}
	return db, nil
}

// Migrate creates/updates tables to match your Go structs. This is the
// GORM equivalent of writing CREATE TABLE by hand.
func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(&models.User{})
}