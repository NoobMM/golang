package main

import (
	"fmt"
	"log"
	"os"

	healthcheckhttp "github.com/NoobMM/golang/app/domain/deliveries/http/health_check"
	"github.com/NoobMM/golang/app/domain/interfaces/connectors"
	healthcheckrepo "github.com/NoobMM/golang/app/domain/repos/health_check"
	healthcheckusecase "github.com/NoobMM/golang/app/domain/usecases/health_check"
	"github.com/NoobMM/golang/app/environments"
	walletrepo "github.com/NoobMM/golang/app/intrastructure/repos/wallet"
	"github.com/NoobMM/golang/app/migration"
	wallethttp "github.com/NoobMM/golang/app/presentasion/http/wallet"
	"github.com/NoobMM/golang/app/routes"
	walletusecase "github.com/NoobMM/golang/app/usecases/wallet"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
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

	log.Printf("MIGRATE APP...")
	err := migration.Migrate()
	if err != nil {
		log.Printf("migration failed %s",err.Error())
		return 
	}
	log.Println("migration success")

	db := connectors.ConnectPostgresDB(false)
	log.Println("starting server...")

	// Repos
	healthCheckRepo := healthcheckrepo.New(db)
	walletRepo := walletrepo.New(db)

	// Usecases
	healthCheckUseCase := healthcheckusecase.New(healthCheckRepo)
	walletUseCase := walletusecase.New(walletRepo)

	// Delivery
	healthCheckHTTP := healthcheckhttp.New(healthCheckUseCase)
	walletHTTP := wallethttp.New(walletUseCase)

	app := gin.Default()
	app.NoRoute(func(c *gin.Context) {
		fmt.Printf("! NO ROUTE")
	})

	routes.ApplyHealthCheckRoutes(app, &routes.HTTPRoutes{
		HealthCheck: healthCheckHTTP,
	})

	routes.ApplyAPIRoutes(app, &routes.HTTPRoutes{
		Wallet: walletHTTP,
	})

	var port = os.Getenv("PORT")
	if port == "" {
		port = "8001"
	}
	_ = app.Run(":" + port)
}