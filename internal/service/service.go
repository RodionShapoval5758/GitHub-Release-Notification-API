package service

import (
	"GithubReleaseNotificationAPI/internal/domain"
	"GithubReleaseNotificationAPI/internal/store"
	"GithubReleaseNotificationAPI/internal/store/repository"
	"GithubReleaseNotificationAPI/internal/store/subscription"
	"context"
	"errors"
	"fmt"
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

	// Manual race condition handling
	repoDomain, err := s.repositoryRepository.FindByFullName(ctx, repo)
	if err != nil {
		if errors.Is(err, store.ErrNotFound) {
			repoDomain, err = s.repositoryRepository.Create(ctx, repo)
			if err != nil {
				if errors.Is(err, store.ErrAlreadyExists) {
					repoDomain, err = s.repositoryRepository.FindByFullName(ctx, repo)
					if err != nil {
						return fmt.Errorf("find repository %s after create conflict: %w", repo, err)
					}
				} else {
					return fmt.Errorf("create repository %s: %w", repo, err)
				}
			}
		} else {
			return fmt.Errorf("find repository %s: %w", repo, err)
		}
	}

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
