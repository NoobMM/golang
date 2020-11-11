package connectors

import (
	//"database/sql"
	"fmt"
	"github.com/deuanz/golang-with-heroku/app/environments"
	"github.com/filipemendespi/newrelic-context/nrgorm"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/jinzhu/gorm"
	"log"
	"os"
)

var (
	// postgresDBInstance is an instance for of postgres db in gorm format
	postgresDBInstance *gorm.DB
)

// ConnectPostgresDB is a connector function for connecting postgres
func ConnectPostgresDB(dbLogMode bool) *gorm.DB {
	return connect(
		environments.PostgresHost,
		environments.PostgresPort,
		environments.PostgresUser,
		environments.PostgresPassword,
		environments.PostgresDB,
		dbLogMode,
	)
}

// RunMigration will run migration script in postgres
func RunMigration() {
	dir, _ := os.Getwd()
	dir += "/migrations"
	if environments.Environment == "heroku" {
		runMigrationUp(
			"schema_migrations",
			"app/migrations",
		)
	} else {
		runMigrationUp(
			"schema_migrations",
			"",
		)
	}

}

func connect(host, port, user, password, dbName string, dbLogMode bool) *gorm.DB {
	var err error
	connection := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host,
		port,
		user,
		password,
		dbName)
	mut.Lock()
	defer mut.Unlock()
	if postgresDBInstance == nil {
		postgresDBInstance, err = gorm.Open("postgres", connection)
		if err != nil {
			panic(err.Error())
		}
		log.Println("database is connected")
	}
	postgresDBInstance.LogMode(dbLogMode)
	nrgorm.AddGormCallbacks(postgresDBInstance)
	return postgresDBInstance
}

func runMigrationUp(migrationTable, dir string) {
	err := getMigrationEngine(migrationTable, dir).Up()
	if err != nil {
		if err.Error() == "no change" {
			log.Printf("Already latest version, Nothing to migrate.")
			return
		}
		log.Println("error occurred :<")
		panic(err)
	} else {
		log.Printf("Migrate up success")
	}
}

func runMigrationDown(migrationTable, dir string) {
	err := getMigrationEngine(migrationTable, dir).Down()
	if err != nil {
		if err.Error() == "no change" {
			log.Printf("Already latest version, Nothing to migrate.")
			return
		}
		log.Println("error occurred :<")
		panic(err)
	} else {
		log.Printf("Migrate down success")
	}
}

func getMigrationEngine(migrationTable, dir string) *migrate.Migrate {
	if dir == "" {
		dir, _ = os.Getwd()
		dir += "/migrations"
	}
	driver, err := postgres.WithInstance(ConnectPostgresDB(false).DB(), &postgres.Config{
		MigrationsTable: migrationTable,
	})
	if err != nil {
		panic(err)
	}
	//log.Printf("DRIVER %+v : ", driver)
	migrationEngine, err := migrate.NewWithDatabaseInstance(
		"file://"+dir,
		"postgres",
		driver,
	)
	if err != nil {
		panic(err)
	}
	return migrationEngine
}
