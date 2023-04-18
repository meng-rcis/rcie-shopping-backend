package routers

import "github.com/labstack/echo"

func InitRouters(e *echo.Echo) *echo.Echo {
	initPrometheusRouterPath(e)

	initAdminRouterPath(e)
	initCartRouterPath(e)
	initProductRouterPath(e)
	initOrderRouterPath(e)
	initSearchRouterPath(e)

	return e
}
