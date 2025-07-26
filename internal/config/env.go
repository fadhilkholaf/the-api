package config

import (
	"log"

	"github.com/joho/godotenv"
)

func InitEnv() {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading environment file: %s", err.Error())
	}
}
