package ports

import (
	"context"
)

type HealthService interface {
	CheckLiveness(ctx context.Context) error
	CheckReadiness(ctx context.Context) error
}
