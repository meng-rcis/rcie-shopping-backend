package routers

import (
	"github.com/labstack/echo"
	"github.com/nuttchai/go-rest/internal/handlers"
	"github.com/nuttchai/go-rest/internal/utils/api"
)

func initCartRouterPath(e *echo.Echo) *echo.Echo {
	e.GET(api.CreatePath("cart"), handlers.CartHandler.GetItems)

	return e
}
