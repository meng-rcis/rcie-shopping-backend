package routers

import (
	"github.com/labstack/echo"
	"github.com/nuttchai/go-rest/internal/handlers"
	"github.com/nuttchai/go-rest/internal/shared/console"
	"github.com/nuttchai/go-rest/internal/utils/api"
)

func initSearchRouterPath(e *echo.Echo) *echo.Echo {
	e.GET(api.CreatePath("search"), handlers.SearchHandler.SearchProduct)
	// API NO.5 With keyword = "" and limit = "100000" with random offset (prevent cache) ~ 5 calls/min
	// API NO.6 With keyword = "" and limit = "1000" with random offset (prevent cache) ~ 500 calls/min

	console.App.Log("Search Router Initialized")
	return e
}
