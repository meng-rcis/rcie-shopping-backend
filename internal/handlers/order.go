package handlers

import (
	"github.com/labstack/echo"
	"github.com/nuttchai/go-rest/internal/constants"
	"github.com/nuttchai/go-rest/internal/services"
	"github.com/nuttchai/go-rest/internal/types"
	"github.com/nuttchai/go-rest/internal/utils/api"
)

type orderHandler struct{}

type orderHandlerInterface interface {
	GetOrders(c echo.Context) error
}

var (
	OrderHandler orderHandlerInterface
)

func init() {
	OrderHandler = &orderHandler{}
}

func (h *orderHandler) GetOrders(c echo.Context) error {
	orderQuery := &types.OrderQuery{
		UserId: c.QueryParam("userId"),
		Status: c.QueryParam("status"),
	}

	orders, err := services.OrderService.GetOrders(orderQuery)
	if err != nil {
		jsonErr := api.InternalServerError(err)
		return c.JSON(jsonErr.Status, jsonErr)
	}

	res := api.SuccessResponse(orders, constants.GetOrdersSuccessMsg)
	return c.JSON(res.Status, res)
}
