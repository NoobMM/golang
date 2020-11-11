package healthcheckusecase

import (
	"context"
	healthcheckrepo "github.com/deuanz/golang-with-heroku/app/domain/repos/health_check"
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
