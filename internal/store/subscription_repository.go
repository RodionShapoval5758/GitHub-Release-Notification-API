package store

import (
	"GithubReleaseNotificationAPI/internal/domain"
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type SubscriptionRepository interface {
	Create(ctx context.Context, subscription domain.Subscription) error
	FindByEmailAndRepositoryID(ctx context.Context, email string, repositoryID int64) (*domain.Subscription, error)
	FindByConfirmToken(ctx context.Context, token string) (*domain.Subscription, error)
	FindByUnsubscribeToken(ctx context.Context, token string) (*domain.Subscription, error)
	Confirm(ctx context.Context, token string) error
	DeleteByUnsubscribeToken(ctx context.Context, token string) error
	ListConfirmedByEmail(ctx context.Context, email string) ([]domain.Subscription, error)
}

type PostgresSubscriptionRepository struct {
	pool *pgxpool.Pool
}

func NewSubscriptionRepository(pool *pgxpool.Pool) *PostgresSubscriptionRepository {
	return &PostgresSubscriptionRepository{
		pool: pool,
	}
}

func (r *PostgresSubscriptionRepository) Create(ctx context.Context, subscription domain.Subscription) error {
	return nil
}

func (r *PostgresSubscriptionRepository) FindByEmailAndRepositoryID(ctx context.Context, email string, repositoryID int64) (*domain.Subscription, error) {
	return nil, nil
}

func (r *PostgresSubscriptionRepository) FindByConfirmToken(ctx context.Context, token string) (*domain.Subscription, error) {
	return nil, nil
}

func (r *PostgresSubscriptionRepository) FindByUnsubscribeToken(ctx context.Context, token string) (*domain.Subscription, error) {
	return nil, nil
}

func (r *PostgresSubscriptionRepository) Confirm(ctx context.Context, token string) error {
	return nil
}

func (r *PostgresSubscriptionRepository) DeleteByUnsubscribeToken(ctx context.Context, token string) error {
	return nil
}

func (r *PostgresSubscriptionRepository) ListConfirmedByEmail(ctx context.Context, email string) ([]domain.Subscription, error) {
	return nil, nil
}
