package handler

import (
	"GithubReleaseNotificationAPI/internal/service"
	"net/http"
)

type Handler struct {
	subscriptionService service.SubscriptionService
}

func New(subscriptionService service.SubscriptionService) *Handler {
	return &Handler{
		subscriptionService: subscriptionService,
	}
}

func (h *Handler) Subscribe(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

func (h *Handler) Confirm(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

func (h *Handler) Unsubscribe(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

func (h *Handler) ListSubscriptions(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}
