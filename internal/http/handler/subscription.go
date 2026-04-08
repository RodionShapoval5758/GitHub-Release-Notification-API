package handler

import (
	"GithubReleaseNotificationAPI/internal/http/models"
	"GithubReleaseNotificationAPI/internal/service"
	"net/http"

	"github.com/go-chi/chi/v5"
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
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "invalid form", http.StatusBadRequest)
		return
	}

	req := models.SubscriptionRequest{
		Email: r.Form.Get("email"),
		Repo:  r.Form.Get("repo"),
	}

	if req.Email == "" || req.Repo == "" {
		http.Error(w, "email/repo is empty", http.StatusBadRequest)
		return
	}

	if err := h.subscriptionService.Subscribe(r.Context(), req.Email, req.Repo); err != nil {
		handleError(err)
		return
	}
}

func (h *Handler) Confirm(w http.ResponseWriter, r *http.Request) {
	token := chi.URLParam(r, "token")
	_ = token
}

func (h *Handler) Unsubscribe(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

func (h *Handler) ListSubscriptions(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}
