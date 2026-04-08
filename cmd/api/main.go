package main

import (
	"GithubReleaseNotificationAPI/internal/config"
	"GithubReleaseNotificationAPI/internal/db"
	httpHandler "GithubReleaseNotificationAPI/internal/http/handler"
	httpRouter "GithubReleaseNotificationAPI/internal/http/router"
	"GithubReleaseNotificationAPI/internal/service"
	"GithubReleaseNotificationAPI/internal/store"
	"context"
	"errors"
	"log"
	"net/http"
	"time"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load env variables: %v", err)
	}

	if err := db.RunMigrations(cfg.DatabaseURL); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	dbPool, err := db.NewPool(ctx, cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Pool creation error: %v", err)
	}
	defer dbPool.Close()

	subscriptionRepository := store.NewSubscriptionRepository(dbPool)
	repositoryRepository := store.NewRepositoryRepository(dbPool)

	subscriptionService := service.NewSubscriptionService(
		subscriptionRepository,
		repositoryRepository,
	)

	handler := httpHandler.New(subscriptionService)
	router := httpRouter.New(handler)

	server := &http.Server{
		Addr:    ":" + cfg.Port,
		Handler: router,
	}

	log.Printf("starting HTTP server on :%s", cfg.Port)

	if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatal(err)
	}
}
