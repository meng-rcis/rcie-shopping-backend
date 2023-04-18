package routers

import (
	"github.com/labstack/echo"
	"github.com/nuttchai/go-rest/internal/handlers"
	"github.com/nuttchai/go-rest/internal/shared/console"
	"github.com/nuttchai/go-rest/internal/utils/api"
)

func initProductRouterPath(e *echo.Echo) {
	e.GET(api.CreatePath("product/:id"), handlers.ProductHandler.GetProduct) // API NO.1

	console.App.Log("Product Router Initialized")
}
