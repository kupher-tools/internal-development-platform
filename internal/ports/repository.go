package ports

import (
	"context"
	"internal-development-platform/internal/domain"
)

type ContainerRepoService interface {
	CreateRepository(ctx context.Context, repoRequest domain.CreateContainerRepositoryRequest) (*domain.RepoResponse, error)
}

type ContainerRepositoryStore interface {
	Create(
		ctx context.Context,
		repo domain.ContainerRepository,
	) error
}
