package internal

import (
	"fmt"

	"github.com/labstack/echo"
	_ "github.com/lib/pq"
	"github.com/nuttchai/go-rest/internal/config"
	"github.com/nuttchai/go-rest/internal/middleware"
	"github.com/nuttchai/go-rest/internal/repositories"
	"github.com/nuttchai/go-rest/internal/routers"
	"github.com/nuttchai/go-rest/internal/services"
	"github.com/nuttchai/go-rest/internal/shared/console"
	"github.com/prometheus/client_golang/prometheus"
)

var appConfig *config.AppConfig
var apiConfig config.APIConfig

func init() {
	// Register Prometheus Metrics
	prometheus.MustRegister(config.ResponseTimeHistogram)
}

func Client() {
	// Add the Configuration into ApiConfig
	console.App.Log("Loading App Configuration...")
	err := config.InitAPIConfig(&apiConfig)
	if err != nil {
		console.App.Fatalf("Error Loading Root Directory (Error: %s)", err.Error())
	}

	// Establish Database Connection
	console.App.Log("Connecting Database...")
	db, err := config.InitSqlDB(&apiConfig)
	if err != nil {
		console.App.Fatalf("Database Connection Failed (Error: %s)", err.Error())
	}
	console.App.Log("Connected Database Successfully")
	defer db.Close()

	// Add the Configuration into AppConfig
	appConfig = &config.AppConfig{
		APIConfig: apiConfig,
		Models:    repositories.InitModels(db),
	}

	// Initialize Services
	console.App.Logf("Initializing Services...")
	repo := services.InitRepo(appConfig)
	services.InitServices(repo)

	// Initialize Routers
	console.App.Logf("Initializing Routers...")
	e := echo.New()
	middleware.EnableCORS(e)
	middleware.EnableResponseTimeMiddleware(e)
	routers.InitRouters(e)

	// Start Server
	console.App.Logf("Starting Server...")
	serverPort := fmt.Sprintf(":%s", apiConfig.Port)
	if err := e.Start(serverPort); err != nil {
		console.App.Fatalf("Server Start Failed (Error: %s)", err.Error())
	}
}
