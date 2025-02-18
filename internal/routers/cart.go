package routers

import (
	"github.com/labstack/echo"
	"github.com/nuttchai/go-rest/internal/handlers"
	"github.com/nuttchai/go-rest/internal/shared/console"
	"github.com/nuttchai/go-rest/internal/utils/api"
)

func initCartRouterPath(e *echo.Echo) {
	e.GET(api.CreatePath("cart"), handlers.CartHandler.GetAllCartItems)
	e.POST(api.CreatePath("cart"), handlers.CartHandler.AddCartItem)
	e.PUT(api.CreatePath("cart"), handlers.CartHandler.UpdateCartItem)
	e.DELETE(api.CreatePath("cart/:id"), handlers.CartHandler.RemoveCartItem)

	console.App.Log("Cart Router Initialized")
}
