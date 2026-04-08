package store

import (
	"GithubReleaseNotificationAPI/internal/domain"
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type RepositoryRepository interface {
	Create(ctx context.Context, repository domain.Repository) error
	FindByFullName(ctx context.Context, fullName string) (*domain.Repository, error)
	UpdateLastSeenTag(ctx context.Context, repositoryID int64, tag string) error
	ListTracked(ctx context.Context) ([]domain.Repository, error)
}

type PostgresRepositoryRepository struct {
	pool *pgxpool.Pool
}

func NewRepositoryRepository(pool *pgxpool.Pool) *PostgresRepositoryRepository {
	return &PostgresRepositoryRepository{
		pool: pool,
	}
}

func (r *PostgresRepositoryRepository) Create(ctx context.Context, repository domain.Repository) error {
	return nil
}

func (r *PostgresRepositoryRepository) FindByFullName(ctx context.Context, fullName string) (*domain.Repository, error) {
	return nil, nil
}

func (r *PostgresRepositoryRepository) UpdateLastSeenTag(ctx context.Context, repositoryID int64, tag string) error {
	return nil
}

func (r *PostgresRepositoryRepository) ListTracked(ctx context.Context) ([]domain.Repository, error) {
	return nil, nil
}
