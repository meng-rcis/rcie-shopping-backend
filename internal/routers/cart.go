package routers

import (
	"github.com/labstack/echo"
	"github.com/nuttchai/go-rest/internal/handlers"
	"github.com/nuttchai/go-rest/internal/shared/console"
	"github.com/nuttchai/go-rest/internal/utils/api"
)

func initCartRouterPath(e *echo.Echo) *echo.Echo {
	e.GET(api.CreatePath("cart"), handlers.CartHandler.GetAllCartItems) // NO.1
	e.POST(api.CreatePath("cart"), handlers.CartHandler.AddCartItem)    // NO.1
	e.PUT(api.CreatePath("cart"), handlers.CartHandler.UpdateCartItem)
	e.DELETE(api.CreatePath("cart/:id"), handlers.CartHandler.RemoveCartItem)

	console.App.Log("Cart Router Initialized")
	return e
}
