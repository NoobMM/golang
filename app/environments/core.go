package environments

import (
	"fmt"
	"os"
	"strings"
	"syscall"
)

// A list of env vars
var (
	Environment      string
	DevMode          bool
	BaseURL          string
	PostgresHost     string
	PostgresPort     string
	PostgresUser     string
	PostgresPassword string
	PostgresDB       string
)

// Init Initialize env variables
func Init() {
	Environment = requireEnv("ENVIRONMENT")
	DevMode = strings.ToLower(os.Getenv("DEV_MODE")) == "true"
	BaseURL = requireEnv("BASE_URL")
	PostgresHost = requireEnv("POSTGRES_HOST")
	PostgresPort = requireEnv("POSTGRES_PORT")
	PostgresUser = requireEnv("POSTGRES_USER")
	PostgresPassword = requireEnv("POSTGRES_PASSWORD")
	PostgresDB = requireEnv("POSTGRES_DB")
}

func requireEnv(envName string) string {
	value, found := syscall.Getenv(envName)

	if !found {
		panic(fmt.Sprintf("%s env is required", envName))
	}

	return value
}

func getEnv(envName string) string {
	return os.Getenv(envName)
}
