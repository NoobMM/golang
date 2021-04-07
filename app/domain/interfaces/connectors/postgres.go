package connectors

import (
	//"database/sql"
	"fmt"
	
	"github.com/NoobMM/golang/app/environments"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
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
	var logLevel logger.LogLevel
	if dbLogMode {
		logLevel = logger.Info
	 } else {
		logLevel = logger.Silent
	 }
	if postgresDBInstance == nil {
		postgresDBInstance, err = gorm.Open(postgres.Open(connection), &gorm.Config{
		   Logger: logger.Default.LogMode(logLevel),
		})
		if err != nil {
		   panic(err.Error())
		}
	 }
	return postgresDBInstance
}
