package healthcheckusecase

import (
	"context"

	healthcheckrepo "github.com/NoobMM/golang/app/domain/repos/health_check"
)

// UseCase is a interface for healthcheckusecase
type UseCase interface {
	CheckReadiness(ctx context.Context) error
}

type useCase struct {
	HealthCheckRepo healthcheckrepo.Repo
}

// New is a constructor method of UseCase
func New(healthCheckRepo healthcheckrepo.Repo) UseCase {
	return &useCase{
		HealthCheckRepo: healthCheckRepo,
	}
}
