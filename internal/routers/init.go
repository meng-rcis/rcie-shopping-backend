package routers

import "github.com/labstack/echo"

func InitRouters(e *echo.Echo) *echo.Echo {
	initSampleRouterPath(e)
	initCartRouterPath(e)
	initProductRouterPath(e)

	return e
}
