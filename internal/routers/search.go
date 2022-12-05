package routers

import (
	"github.com/labstack/echo"
	"github.com/nuttchai/go-rest/internal/handlers"
	"github.com/nuttchai/go-rest/internal/shared/console"
	"github.com/nuttchai/go-rest/internal/utils/api"
)

func initSearchRouterPath(e *echo.Echo) *echo.Echo {
	e.GET(api.CreatePath("search"), handlers.SearchHandler.SearchProduct)

	console.App.Log("Search Router Initialized")
	return e
}
