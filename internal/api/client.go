package internal

import (
	"flag"
	"fmt"

	"github.com/labstack/echo"
	_ "github.com/lib/pq"
	"github.com/nuttchai/go-rest/internal/config"
	"github.com/nuttchai/go-rest/internal/middleware"
	"github.com/nuttchai/go-rest/internal/repositories"
	"github.com/nuttchai/go-rest/internal/routers"
	"github.com/nuttchai/go-rest/internal/services"
	"github.com/nuttchai/go-rest/internal/utils/db"
	"github.com/nuttchai/go-rest/internal/utils/env"
)

var appConfig *config.AppConfig
var apiConfig config.APIConfig

func Client() {
	// Load Environment Variables
	config.App.Log("Loading ENV...")
	appEnv := env.GetEnv("APP_ENV", "development")
	envDefaultDir, err := env.GetDefaultEnvDir(appEnv)
	if err != nil {
		config.App.Fatalf("Error Loading Root Directory (Error: %s)", err.Error())
	}

	envDir := env.GetEnv("ENV_PATH", envDefaultDir)
	env.LoadEnv(envDir)
	dbType := env.GetEnv("DB_TYPE", "postgres")
	dbUser := env.GetEnv("APP_DB_USER", "root")
	dbPass := env.GetEnv("APP_DB_PASS", "password")
	dbHost := env.GetEnv("DB_HOST", "localhost")
	dbPort := env.GetEnv("DB_PORT", "27017")
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
	flag.Parse()

	// Establish Database Connection
	config.App.Logf("Connecting database: %s", dbConnStr)
	db, err := db.OpenSqlDB(dbDriver, apiConfig)
	if err != nil {
		config.App.Fatalf("Database Connection Failed (Error: %s)", err.Error())
	}
	defer db.Close()

	// Add the Configuration into AppConfig
	appConfig = &config.AppConfig{
		APIConfig: apiConfig,
		Models:    repositories.InitModels(db),
	}

	// Initialize Services
	config.App.Logf("Initializing Services...")
	repo := services.InitRepo(appConfig)
	services.InitServices(repo)

	// Initialize Routers
	config.App.Logf("Initializing Routers...")
	e := echo.New()
	middleware.EnableCORS(e)
	routers.InitRouters(e)

	// Start Server
	config.App.Logf("Starting Server...")
	serverPort := fmt.Sprintf(":%s", apiConfig.Port)

	if err := e.Start(serverPort); err != nil {
		config.App.Fatalf("Server Start Failed (Error: %s)", err.Error())
	}
}
