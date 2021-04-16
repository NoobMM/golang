package healthcheckusecase

import "context"

func (uc *useCase) CheckReadiness(ctx context.Context) error {
	return uc.HealthCheckRepo.CheckDatabaseReadiness(ctx)
}
