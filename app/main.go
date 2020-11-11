package main

import (
	"fmt"
	healthcheckhttp "github.com/deuanz/golang-with-heroku/app/domain/deliveries/http/health_check"
	"github.com/deuanz/golang-with-heroku/app/domain/interfaces/connectors"
	healthcheckrepo "github.com/deuanz/golang-with-heroku/app/domain/repos/health_check"
	healthcheckusecase "github.com/deuanz/golang-with-heroku/app/domain/usecases/health_check"
	"github.com/deuanz/golang-with-heroku/app/environments"
	"github.com/deuanz/golang-with-heroku/app/routes"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"log"
	"os"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "app",
		Short: "Start app",
		Run:   startApp,
	}

	if err := rootCmd.Execute(); err != nil {
		log.Println(err)
	}
}

func startApp(_ *cobra.Command, _ []string) {
	environments.Init()

	//log.Printf("MIGRATE APP...")
	//connectors.RunMigration()

	db := connectors.ConnectPostgresDB(false)
	log.Println("starting server...")

	// Repos
	healthCheckRepo := healthcheckrepo.New(db)

	// Usecases
	healthCheckUseCase := healthcheckusecase.New(healthCheckRepo)

	// Delivery
	healthCheckHTTP := healthcheckhttp.New(healthCheckUseCase)

	app := gin.Default()
	app.NoRoute(func(c *gin.Context) {
		fmt.Printf("! NO ROUTE")
	})

	routes.ApplyHealthCheckRoutes(app, &routes.HTTPRoutes{
		HealthCheck: healthCheckHTTP,
	})

	var port = os.Getenv("PORT")
	if port == "" {
		port = "8001"
	}
	_ = app.Run(":" + port)
}
