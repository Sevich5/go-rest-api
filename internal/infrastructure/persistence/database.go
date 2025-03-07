package persistence

import (
	"app/internal/infrastructure/configuration"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func ConnectDatabase(cfg *configuration.Config) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.Database.Host, cfg.Database.Username, cfg.Database.Password, cfg.Database.Database, cfg.Database.Port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}

	return db
}
