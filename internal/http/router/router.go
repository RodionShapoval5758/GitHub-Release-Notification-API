package router

import (
	"GithubReleaseNotificationAPI/internal/http/handler"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func New(handler *handler.Handler) http.Handler {
	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Post("/api/subscribe", handler.Subscribe)
	router.Get("/api/confirm/{token}", handler.Confirm)
	router.Get("/api/unsubscribe/{token}", handler.Unsubscribe)
	router.Get("/api/subscriptions", handler.ListSubscriptions)

	return router
}
