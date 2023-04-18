package routers

import (
	"github.com/labstack/echo"
	"github.com/nuttchai/go-rest/internal/handlers"
	"github.com/nuttchai/go-rest/internal/shared/console"
	"github.com/nuttchai/go-rest/internal/utils/api"
)

func initOrderRouterPath(e *echo.Echo) {
	e.GET(api.CreatePath("order"), handlers.OrderHandler.GetOrders)    // NO.1
	e.POST(api.CreatePath("order"), handlers.OrderHandler.CreateOrder) // NO.1

	console.App.Log("Order Router Initialized")
}
