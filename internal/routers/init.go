package routers

import "github.com/labstack/echo"

func InitRouters(e *echo.Echo) *echo.Echo {
	initSampleRouterPath(e)
	initAdminRouterPath(e)
	initCartRouterPath(e)
	initProductRouterPath(e)
	initOrderRouterPath(e)
	initSearchRouterPath(e)

	return e
}
