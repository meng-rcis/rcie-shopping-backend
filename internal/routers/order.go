package routers

import (
	"github.com/labstack/echo"
	"github.com/nuttchai/go-rest/internal/handlers"
	"github.com/nuttchai/go-rest/internal/shared/console"
	"github.com/nuttchai/go-rest/internal/utils/api"
)

func initOrderRouterPath(e *echo.Echo) *echo.Echo {
	e.GET(api.CreatePath("order"), handlers.OrderHandler.GetOrders)
	e.POST(api.CreatePath("order"), handlers.OrderHandler.CreateOrder)

	console.App.Log("Order Router Initialized")
	return e
}
