package routes

import (
	healthcheckhttp "github.com/NoobMM/golang/app/presentation/http/health_check"
	wallethttp "github.com/NoobMM/golang/app/presentation/http/wallet"
	"github.com/gin-gonic/gin"
)

// HTTPRoutes ...
type HTTPRoutes struct {
	HealthCheck healthcheckhttp.HTTPHandler
	Wallet      wallethttp.HTTPHandler
}

// ApplyAPIRoutes is a function for grouping and applying api routes
func ApplyAPIRoutes(r *gin.Engine, httpRoutes *HTTPRoutes) {
	apiRoute := r.Group("/api")
	{
		apiRoute.POST("/wallets", httpRoutes.Wallet.CreateWallet)
		apiRoute.POST("/wallets.add-balance", httpRoutes.Wallet.UpdateWallet)
		apiRoute.GET("/wallets/:walletID", httpRoutes.Wallet.FindWallet)
	}
}

// ApplyHealthCheckRoutes is a function for grouping and applying health check routes
func ApplyHealthCheckRoutes(r *gin.Engine, httpRoutes *HTTPRoutes) {
	apiRoute := r.Group("/_healthz")
	{
		apiRoute.GET("/liveness", httpRoutes.HealthCheck.CheckLiveness)
		apiRoute.GET("/readiness", httpRoutes.HealthCheck.CheckReadiness)
	}
}
