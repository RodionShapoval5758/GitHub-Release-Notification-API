package service

import (
	"GithubReleaseNotificationAPI/internal/domain"
	"GithubReleaseNotificationAPI/internal/store"
	"context"
)

type SubscriptionService interface {
	Subscribe(ctx context.Context, email string, repo string) error
	Confirm(ctx context.Context, token string) error
	Unsubscribe(ctx context.Context, token string) error
	ListByEmail(ctx context.Context, email string) ([]domain.Subscription, error)
}

type subscriptionService struct {
	subscriptionRepository store.SubscriptionRepository
	repositoryRepository   store.RepositoryRepository
}

func NewSubscriptionService(
	subscriptionRepository store.SubscriptionRepository,
	repositoryRepository store.RepositoryRepository,
) SubscriptionService {
	return &subscriptionService{
		subscriptionRepository: subscriptionRepository,
		repositoryRepository:   repositoryRepository,
	}
}

func (s *subscriptionService) Subscribe(ctx context.Context, email string, repo string) error {
	return nil
}

func (s *subscriptionService) Confirm(ctx context.Context, token string) error {
	return nil
}

func (s *subscriptionService) Unsubscribe(ctx context.Context, token string) error {
	return nil
}

func (s *subscriptionService) ListByEmail(ctx context.Context, email string) ([]domain.Subscription, error) {
	return nil, nil
}
