package handler

import (
	"GithubReleaseNotificationAPI/internal/http/models"
	"GithubReleaseNotificationAPI/internal/http/util"
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

	util.WriteJSONResponse(w, http.StatusOK, nil)
}

func (h *Handler) Confirm(w http.ResponseWriter, r *http.Request) {
	token := chi.URLParam(r, "token")

	if token == "" || len(token) < 8 {
		http.Error(w, "invalid token", http.StatusBadRequest)
		return
	}

	err := h.subscriptionService.Confirm(r.Context(), token)
	if err != nil {
		handleError(err)
		return
	}

	util.WriteJSONResponse(w, http.StatusOK, nil)
}

func (h *Handler) Unsubscribe(w http.ResponseWriter, r *http.Request) {
	token := chi.URLParam(r, "token")

	if token == "" || len(token) < 8 {
		http.Error(w, "invalid token", http.StatusBadRequest)
		return
	}

	err := h.subscriptionService.Unsubscribe(r.Context(), token)
	if err != nil {
		handleError(err)
		return
	}
	util.WriteJSONResponse(w, http.StatusOK, nil)
}

func (h *Handler) ListSubscriptions(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")

	if email == "" {
		http.Error(w, "empty email", http.StatusBadRequest)
	}
}
