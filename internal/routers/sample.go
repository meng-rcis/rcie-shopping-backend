package routers

import (
	"github.com/labstack/echo"
	"github.com/nuttchai/go-rest/internal/handlers"
	"github.com/nuttchai/go-rest/internal/utils/api"
)

func initSampleRouterPath(e *echo.Echo) *echo.Echo {
	e.GET(api.CreatePath("sample/:id"), handlers.SampleHandler.GetSample)

	e.POST(api.CreatePath("sample"), handlers.SampleHandler.CreateSample)

	e.PUT(api.CreatePath("sample/:id"), handlers.SampleHandler.UpdateSample)
	e.PUT(api.CreatePath("sample/:id/empty"), handlers.SampleHandler.EmptySampleDesc)

	e.DELETE(api.CreatePath("sample/:id"), handlers.SampleHandler.DeleteSample)

	return e
}
