package routers

import (
	"github.com/labstack/echo"
	"github.com/nuttchai/go-rest/internal/handlers"
	"github.com/nuttchai/go-rest/internal/shared/console"
	"github.com/nuttchai/go-rest/internal/utils/api"
)

func initAdminRouterPath(e *echo.Echo) *echo.Echo {
	e.PUT(api.CreatePath("admin/order"), handlers.AdminHandler.UpdateOrderStatus)
	e.PUT(api.CreatePath("admin/product/quantity"), handlers.AdminHandler.AddProductQuantity)

	e.PUT(api.CreatePath("admin/product/bulk"), handlers.AdminHandler.BulkUpdateProduct)
	// NO.13 Search + Loop to recalculate the product internally with keyword = "" and limit = "100000" with random offset (prevent cache) ~ 5 calls/min + randomly choose and save 10~20 products into DB
	// NO.14 Search + Loop to recalculate the product internally with keyword = "" and limit = "1000" with random offset (prevent cache) ~ 500 calls/min + randomly choose and save 10~20 products into DB

	// NOTE: NO.15 and NO.16 use "admin/product/bulk" as same as NO.13 and NO.14 but require query parameter "isPreview" = "true"
	// NO.15 Search + Loop to recalculate the product internally with keyword = "" and limit = "10000" with random offset (prevent cache) ~ 5 calls/min + choose and save all products into DB
	// NO.16 Search + Loop to recalculate the product internally with keyword = "" and limit = "1000" with random offset (prevent cache) ~ 500 calls/min + choose and save all products into DB

	console.App.Log("Admin Router Initialized")
	return e
}
