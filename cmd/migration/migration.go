package main

import (
	"github.com/fadhilkholaf/go-gorm/internal/config"
	"github.com/fadhilkholaf/go-gorm/internal/database"
)

func main() {
	config.InitEnv()

	db := database.NewConnection()
	database.Migrate(db)
	defer database.CloseConnection(db)
}
