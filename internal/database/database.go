package database

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/fadhilkholaf/go-gorm/internal/model"
)

func NewConnection() *gorm.DB {
	db, err := gorm.Open(postgres.Open(os.Getenv("DB_URL")), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})

	if err != nil {
		log.Fatalf("Database connection error: %s", err.Error())
	}

	err = db.AutoMigrate(&model.User{})

	if err != nil {
		log.Fatalf("Model migration error: %s", err.Error())
	}

	return db
}

func CloseConnection(db *gorm.DB) {
	sqlDB, err := db.DB()

	if err != nil {
		log.Fatalf("Error getting sqlDB: %s", err.Error())
	}

	err = sqlDB.Close()

	if err != nil {
		log.Fatalf("Error closing database: %s", err.Error())
	}
}
