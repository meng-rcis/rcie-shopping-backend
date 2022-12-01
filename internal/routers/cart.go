package routers

import (
	"github.com/labstack/echo"
	"github.com/nuttchai/go-rest/internal/handlers"
	"github.com/nuttchai/go-rest/internal/shared/console"
	"github.com/nuttchai/go-rest/internal/utils/api"
)

func initCartRouterPath(e *echo.Echo) *echo.Echo {
	// userId in Query Param is required
	e.GET(api.CreatePath("cart"), handlers.CartHandler.GetAllCartProducts)

	e.POST(api.CreatePath("cart"), handlers.CartHandler.AddCartProduct)

	console.App.Log("Cart Router Initialized")
	return e
}
