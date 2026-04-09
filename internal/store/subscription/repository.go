package subscription

import (
	"GithubReleaseNotificationAPI/internal/domain"
	"GithubReleaseNotificationAPI/internal/store"
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository interface {
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
	tag, err := r.pool.Exec(
		ctx,
		createSubscriptionQuery,
		subscription.Email,
		subscription.RepositoryID,
		subscription.ConfirmToken,
		subscription.UnsubscribeToken,
	)
	if err != nil {
		var pgErr *pgconn.PgError
		//TODO differentiate unique vialoation with the constraint name
		if errors.As(err, &pgErr) && pgErr.Code == pgerrcode.UniqueViolation {
			return store.ErrAlreadyExists
		}

		return fmt.Errorf(
			"insert subscription for email %s and repository_id %d: %w",
			subscription.Email,
			subscription.RepositoryID,
			err,
		)
	}

	if rowsAffected := tag.RowsAffected(); rowsAffected != 1 {
		return fmt.Errorf("insert subscription row: expected 1 affected row, got %d", rowsAffected)
	}

	return nil
}

func (r *PostgresSubscriptionRepository) FindByEmailAndRepositoryID(ctx context.Context, email string, repositoryID int64) (*domain.Subscription, error) {
	subscription, err := scanSubscription(
		r.pool.QueryRow(
			ctx,
			findSubscriptionByEmailAndRepositoryIDQuery,
			email,
			repositoryID,
		),
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, store.ErrNotFound
		}

		return nil, fmt.Errorf(
			"find subscription by email %s and repository_id %d: %w",
			email,
			repositoryID,
			err,
		)
	}

	return subscription, nil
}

func (r *PostgresSubscriptionRepository) FindByConfirmToken(ctx context.Context, token string) (*domain.Subscription, error) {
	subscription, err := scanSubscription(
		r.pool.QueryRow(
			ctx,
			findSubscriptionByConfirmTokenQuery,
			token,
		),
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, store.ErrNotFound
		}

		return nil, fmt.Errorf("find subscription by confirm token: %w", err)
	}

	return subscription, nil
}

func (r *PostgresSubscriptionRepository) FindByUnsubscribeToken(ctx context.Context, token string) (*domain.Subscription, error) {
	subscription, err := scanSubscription(
		r.pool.QueryRow(
			ctx,
			findSubscriptionByUnsubscribeTokenQuery,
			token,
		),
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, store.ErrNotFound
		}

		return nil, fmt.Errorf("find subscription by unsubscribe token: %w", err)
	}

	return subscription, nil
}

func (r *PostgresSubscriptionRepository) Confirm(ctx context.Context, token string) error {
	tag, err := r.pool.Exec(
		ctx,
		confirmSubscriptionByTokenQuery,
		token,
	)
	if err != nil {
		return fmt.Errorf("confirm subscription by token: %w", err)
	}

	if tag.RowsAffected() == 0 {
		return store.ErrNotFound
	}

	return nil
}

func (r *PostgresSubscriptionRepository) DeleteByUnsubscribeToken(ctx context.Context, token string) error {
	tag, err := r.pool.Exec(
		ctx,
		deleteSubscriptionByUnsubscribeTokenQuery,
		token,
	)
	if err != nil {
		return fmt.Errorf("delete subscription by unsubscribe token: %w", err)
	}

	if tag.RowsAffected() == 0 {
		return store.ErrNotFound
	}

	return nil
}

func (r *PostgresSubscriptionRepository) ListConfirmedByEmail(ctx context.Context, email string) ([]domain.Subscription, error) {
	return nil, nil
}

func scanSubscription(row pgx.Row) (*domain.Subscription, error) {
	var subscription domain.Subscription

	err := row.Scan(
		&subscription.ID,
		&subscription.Email,
		&subscription.RepositoryID,
		&subscription.Confirmed,
		&subscription.ConfirmToken,
		&subscription.UnsubscribeToken,
		&subscription.CreatedAt,
		&subscription.ConfirmedAt,
	)
	if err != nil {
		return nil, err
	}

	return &subscription, nil
}
