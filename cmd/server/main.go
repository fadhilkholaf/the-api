package main

import (
	"context"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/fadhilkholaf/go-gorm/internal/config"
	"github.com/fadhilkholaf/go-gorm/internal/database"
	"github.com/fadhilkholaf/go-gorm/internal/router"
)

func main() {
	config.InitEnv()

	db := database.NewConnection()
	defer database.CloseConnection(db)

	r := router.NewRouter(db)

	srv := http.Server{
		Addr:    ":8080",
		Handler: r.Handler(),
	}

	go func() {
		err := srv.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatalf("Error serving server: %s", err.Error())
		}
	}()

	shutdown, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer stop()
	<-shutdown.Done()

	timeout, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := srv.Shutdown(timeout)
	if err != nil {
		log.Fatalf("Error shutting down server: %s", err.Error())
	}
}
