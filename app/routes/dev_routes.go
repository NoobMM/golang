package routes

import (
	"github.com/gin-gonic/gin"
)

// ApplyDevTools add dev's api to the app
func ApplyDevTools(r *gin.Engine, routes *HTTPRoutes) {
	//apiRoute := r.Group("/dev")
	{
		//apiRoute.POST("/seed", routes.Seed.Init)
		//apiRoute.POST("/apikeys.migrate", routes.Auth.MigrateAPISecret)
	}
}
