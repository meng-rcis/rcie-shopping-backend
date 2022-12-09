package handlers

import (
	"github.com/labstack/echo"
	"github.com/nuttchai/go-rest/internal/constants"
	orderdto "github.com/nuttchai/go-rest/internal/dto/order"
	"github.com/nuttchai/go-rest/internal/services"
	"github.com/nuttchai/go-rest/internal/types"
	"github.com/nuttchai/go-rest/internal/utils/api"
)

type orderHandler struct{}

type orderHandlerInterface interface {
	GetOrders(c echo.Context) error
	CreateOrder(c echo.Context) error
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

func (h *orderHandler) CreateOrder(c echo.Context) error {
	var reqBody orderdto.CreateOrderDTO
	if err := api.DecodeDTO(c, &reqBody); err != nil {
		jsonErr := api.BadRequestError(err)
		return c.JSON(jsonErr.Status, jsonErr)
	}

	order, err := services.OrderService.CreateOrder(&reqBody)
	if err != nil {
		jsonErr := api.InternalServerError(err)
		return c.JSON(jsonErr.Status, jsonErr)
	}

	res := api.SuccessResponse(order, constants.CreateOrderSuccessMsg)
	return c.JSON(res.Status, res)
}
