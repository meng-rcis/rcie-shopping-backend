package routers

import (
	"github.com/labstack/echo"
	"github.com/nuttchai/go-rest/internal/handlers"
	"github.com/nuttchai/go-rest/internal/shared/console"
	"github.com/nuttchai/go-rest/internal/utils/api"
)

func initAdminRouterPath(e *echo.Echo) *echo.Echo {
	e.PUT(api.CreatePath("admin/order"), handlers.AdminHandler.UpdateOrderStatus)
	e.PUT(api.CreatePath("admin/product"), handlers.AdminHandler.AddProductQuantity)

	console.App.Log("Cart Router Initialized")
	return e
}
