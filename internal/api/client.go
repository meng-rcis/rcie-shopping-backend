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
	"github.com/nuttchai/go-rest/internal/utils/db"
)

var appConfig *config.AppConfig
var apiConfig config.APIConfig

func Client() {
	// Add the Configuration into ApiConfig
	config.App.Log("Loading App Configuration...")
	err := config.InitAPIConfig(&apiConfig)
	if err != nil {
		config.App.Fatalf("Error Loading Root Directory (Error: %s)", err.Error())
	}

	// Establish Database Connection
	config.App.Log("Connecting database...")
	db, err := db.OpenSqlDB(&apiConfig)
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
