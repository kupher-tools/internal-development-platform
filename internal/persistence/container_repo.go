package persistence

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"

	"internal-development-platform/internal/domain"
)

type ContainerRepositoryStore struct {
	db *pgxpool.Pool
}

func NewContainerRepositoryStore(
	db *pgxpool.Pool,
) *ContainerRepositoryStore {
	return &ContainerRepositoryStore{
		db: db,
	}
}

func (s *ContainerRepositoryStore) Create(
	ctx context.Context,
	repo domain.ContainerRepository,
) error {

	query := `
	INSERT INTO container_repositories (
		id,
		name,
		provider,
		visibility,
		description,
		desired_state,
		actual_state,
		status,
		generation,
		observed_generation,
		created_at,
		updated_at
	)
	VALUES (
		$1, $2, $3, $4, $5,
		$6, $7, $8, $9, $10,
		NOW(), NOW()
	)
	`

	_, err := s.db.Exec(
		ctx,
		query,
		uuid.New(),
		repo.Name,
		repo.Provider,
		repo.Visibility,
		repo.Description,
		repo.DesiredState,
		repo.ActualState,
		repo.Status,
		repo.Generation,
		repo.ObservedGeneration,
	)

	if err != nil {
		return fmt.Errorf(
			"failed to create container repository: %w",
			err,
		)
	}

	return nil
}
