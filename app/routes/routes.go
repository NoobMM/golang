package routes

import (
	healthcheckhttp "github.com/NoobMM/golang/app/domain/deliveries/http/health_check"
	"github.com/gin-gonic/gin"
)

// HTTPRoutes ...
type HTTPRoutes struct {
	HealthCheck healthcheckhttp.HTTPHandler
}

// ApplyHealthCheckRoutes is a function for grouping and applying health check routes
func ApplyHealthCheckRoutes(r *gin.Engine, httpRoutes *HTTPRoutes) {
	apiRoute := r.Group("/_healthz")
	{
		apiRoute.GET("/liveness", httpRoutes.HealthCheck.CheckLiveness)
		apiRoute.GET("/readiness", httpRoutes.HealthCheck.CheckReadiness)
	}
}
