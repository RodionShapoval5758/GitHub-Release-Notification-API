package repository

import (
	"GithubReleaseNotificationAPI/internal/domain"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository interface {
	Create(ctx context.Context, repositoryName string) error
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

func (r *PostgresRepositoryRepository) Create(ctx context.Context, repositoryName string) error {
	tag, err := r.pool.Exec(ctx, createRepoQuery, repositoryName)
	if err != nil {
		return fmt.Errorf("repositories: insert row with name %s: %v", repositoryName, err)
	}
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
