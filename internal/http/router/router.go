package router

import (
	"GithubReleaseNotificationAPI/internal/http/handler"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func New(handler *handler.Handler) http.Handler {
	router := chi.NewRouter()

	router.Post("/api/subscribe", handler.Subscribe)
	router.Get("/api/confirm/{token}", handler.Confirm)
	router.Get("/api/unsubscribe/{token}", handler.Unsubscribe)
	router.Get("/api/subscriptions", handler.ListSubscriptions)

	return router
}
