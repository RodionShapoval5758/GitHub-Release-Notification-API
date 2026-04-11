package service

import (
	"GithubReleaseNotificationAPI/internal/domain"
	gh "GithubReleaseNotificationAPI/internal/github"
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

type githubClient interface {
	CheckRepo(ctx context.Context, fullName string) error
}

type subscriptionService struct {
	subscriptionRepository subscription.Repository
	repositoryRepository   repository.Repository
	githubClient           githubClient
}

func NewSubscriptionService(
	subscriptionRepository subscription.Repository,
	repositoryRepository repository.Repository,
	githubClient githubClient,
) SubscriptionService {
	return &subscriptionService{
		subscriptionRepository: subscriptionRepository,
		repositoryRepository:   repositoryRepository,
		githubClient:           githubClient,
	}
}

const TokenLength = 32

func (s *subscriptionService) Subscribe(ctx context.Context, email string, repo string) error {
	if err := validateEmailFormat(email); err != nil {
		return err
	}

	if err := validateRepoFormat(repo); err != nil {
		return err
	}

	if err := s.githubClient.CheckRepo(ctx, repo); err != nil {
		switch {
		case errors.Is(err, gh.ErrNotFound):
			return ErrRepoNotFound
		case errors.Is(err, gh.ErrRateLimited):
			return ErrTooMuchRequests
		case errors.Is(err, gh.ErrUnexpectedResponse):
			return fmt.Errorf("github repo check failed: %w", err)
		default:
			return fmt.Errorf("github repo check request failed: %w", err)
		}
	}

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

	confirmToken, err := GenerateToken(TokenLength)
	if err != nil {
		return fmt.Errorf("create token: %w", err)
	}

	unsubscribeToken, err := GenerateToken(TokenLength)
	if err != nil {
		return fmt.Errorf("create token: %w", err)
	}

	subscriptionInput := domain.Subscription{
		Email:            email,
		RepositoryID:     repoDomain.ID,
		ConfirmToken:     confirmToken,
		UnsubscribeToken: unsubscribeToken,
	}

	if err := s.subscriptionRepository.Create(ctx, subscriptionInput); err != nil {
		switch {
		case errors.Is(err, store.ErrAlreadyExists):
			return ErrSubscriptionAlreadyExists
		case errors.Is(err, store.ErrTokensAlreadyExists):
			// TODO regenerate tokens
			return fmt.Errorf("create subscription tokens conflict: %w", err)
		default:
			return fmt.Errorf("failed to create subscription: %w", err)
		}
	}

	// TODO sent email confirmation
	return nil
}

func (s *subscriptionService) Confirm(ctx context.Context, token string) error {
	err := s.subscriptionRepository.Confirm(ctx, token)
	if err != nil {
		return err
	}
	return nil
}

func (s *subscriptionService) Unsubscribe(ctx context.Context, token string) error {
	return nil
}

func (s *subscriptionService) ListByEmail(ctx context.Context, email string) ([]domain.Subscription, error) {
	return nil, nil
}
