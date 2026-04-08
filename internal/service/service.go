package service

import (
	"GithubReleaseNotificationAPI/internal/domain"
	"GithubReleaseNotificationAPI/internal/store/repository"
	"GithubReleaseNotificationAPI/internal/store/subscription"
	"context"
)

type SubscriptionService interface {
	Subscribe(ctx context.Context, email string, repo string) error
	Confirm(ctx context.Context, token string) error
	Unsubscribe(ctx context.Context, token string) error
	ListByEmail(ctx context.Context, email string) ([]domain.Subscription, error)
}

type subscriptionService struct {
	subscriptionRepository subscription.Repository
	repositoryRepository   repository.Repository
}

func NewSubscriptionService(
	subscriptionRepository subscription.Repository,
	repositoryRepository repository.Repository,
) SubscriptionService {
	return &subscriptionService{
		subscriptionRepository: subscriptionRepository,
		repositoryRepository:   repositoryRepository,
	}
}

func (s *subscriptionService) Subscribe(ctx context.Context, email string, repo string) error {
	if err := validateEmailFormat(email); err != nil {
		return err
	}

	if err := validateRepoFormat(repo); err != nil {
		return err
	}

	// TODO GitHub repo check
	// TODO no duplicate subscription

	// TODO sent email confirmation
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
