package service

import (
	"context"
	appErrors "internal-development-platform/internal/errors"
)

type HealthService struct {
	version string
}

func NewHealthService(version string) *HealthService {
	return &HealthService{version: version}
}

func (s *HealthService) CheckLiveness(ctx context.Context) error {
	dbHealth := true
	if !dbHealth {
		return appErrors.ErrDatabaseDown
	}
	return nil
}

func (s *HealthService) CheckReadiness(ctx context.Context) error {
	appHealth := false
	if !appHealth {
		return appErrors.ErrDatabaseDown
	}
	return nil
}
