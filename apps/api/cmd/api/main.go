package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	appapi "github.com/aldihorn/pact-cfp-demo/apps/api/internal/api"
	"github.com/aldihorn/pact-cfp-demo/apps/api/internal/config"
	"github.com/aldihorn/pact-cfp-demo/apps/api/internal/db"
	"github.com/aldihorn/pact-cfp-demo/apps/api/internal/repository"
	"github.com/aldihorn/pact-cfp-demo/apps/api/internal/service"
)

func main() {
	cfg := config.Load()

	database, err := db.ConnectFromConfig(cfg)
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	repo := repository.NewTodoRepository(database)
	svc := service.NewTodoService(repo)
	router := appapi.NewRouter(svc, cfg)

	srv := &http.Server{
		Addr:              ":" + cfg.Port,
		Handler:           router,
		ReadHeaderTimeout: 5 * time.Second,
	}

	go func() {
		log.Printf("api listening on :%s", cfg.Port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("server failed: %v", err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Printf("shutdown error: %v", err)
	}
}
