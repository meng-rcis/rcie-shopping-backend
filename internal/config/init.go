package config

import (
	"flag"
	"fmt"

	"github.com/nuttchai/go-rest/internal/utils/env"
)

func InitAPIConfig(apiConfig *APIConfig) error {
	// Load Environment Variables
	appEnv := env.GetEnv("APP_ENV", "development")
	envDefaultDir, err := env.GetDefaultEnvDir(appEnv)
	if err != nil {
		return err
	}

	envDir := env.GetEnv("ENV_PATH", envDefaultDir)
	env.LoadEnv(envDir)

	dbType := env.GetEnv("DB_TYPE", "postgres")
	dbUser := env.GetEnv("APP_DB_USER", "postgres")
	dbPass := env.GetEnv("APP_DB_PASS", "postgres")
	dbHost := env.GetEnv("DB_HOST", "db")
	dbPort := env.GetEnv("DB_PORT", "5432")
	dbName := env.GetEnv("APP_DB_NAME", "shopping")
	dbDriver := env.GetEnv("DB_DRIVER", "postgres")
	port := env.GetEnv("APP_PORT", "8000")
	dbConnStr := fmt.Sprintf(
		"%s://%s:%s@%s:%s/%s?sslmode=disable",
		dbType,
		dbUser,
		dbPass,
		dbHost,
		dbPort,
		dbName,
	)

	// Store ENV Variables to ApiConfig
	flag.StringVar(&apiConfig.Env, "env", appEnv, "Application Environment")
	flag.StringVar(&apiConfig.Port, "port", port, "Server Listening Port")
	flag.StringVar(&apiConfig.Db.Dsn, "dsn", dbConnStr, "Data Source Name")
	flag.StringVar(&apiConfig.Db.Driver, "driver", dbDriver, "Database Driver")
	flag.Parse()

	return nil
}
