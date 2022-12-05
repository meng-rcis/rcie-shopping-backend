package routers

import (
	"github.com/labstack/echo"
	"github.com/nuttchai/go-rest/internal/handlers"
	"github.com/nuttchai/go-rest/internal/shared/console"
	"github.com/nuttchai/go-rest/internal/utils/api"
)

func initProductRouterPath(e *echo.Echo) *echo.Echo {
	e.GET(api.CreatePath("product/:id"), handlers.ProductHandler.GetProduct)

	console.App.Log("Cart Router Initialized")
	return e
}
