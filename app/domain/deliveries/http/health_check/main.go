package healthcheckhttp

import (
	healthcheckusecase "github.com/NoobMM/golang/app/domain/usecases/health_check"
	"github.com/gin-gonic/gin"
)

// A HTTPHandler for HealthCheck endpoints
type HTTPHandler interface {
	CheckLiveness(c *gin.Context)
	CheckReadiness(c *gin.Context)
}

type httpHandler struct {
	HealthCheckUseCase healthcheckusecase.UseCase
}

// New is a function to create a new HealthCheck endpoint
func New(healthCheckUseCase healthcheckusecase.UseCase) HTTPHandler {
	return &httpHandler{
		HealthCheckUseCase: healthCheckUseCase,
	}
}
