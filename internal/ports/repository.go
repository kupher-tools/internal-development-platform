package ports

import (
	"context"
	"internal-development-platform/internal/domain"
)

type ContainerRepoService interface {
	CreateRepository(ctx context.Context, repoRequest domain.RepoRequest) (*domain.RepoResponse, error)
}
