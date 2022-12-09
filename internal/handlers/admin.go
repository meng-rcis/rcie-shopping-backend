package handlers

import (
	"github.com/labstack/echo"
	"github.com/nuttchai/go-rest/internal/constants"
	admindto "github.com/nuttchai/go-rest/internal/dto/admin"
	"github.com/nuttchai/go-rest/internal/services"
	"github.com/nuttchai/go-rest/internal/utils/api"
)

type adminHandler struct{}

type adminHandlerInterface interface {
	UpdateOrderStatus(c echo.Context) error
	AddProductQuantity(c echo.Context) error
}

var (
	AdminHandler adminHandlerInterface
)

func init() {
	AdminHandler = &adminHandler{}
}

func (h *adminHandler) UpdateOrderStatus(c echo.Context) error {
	var reqBody admindto.UpdateOrderStatusDTO
	if err := api.DecodeDTO(c, &reqBody); err != nil {
		jsonErr := api.BadRequestError(err)
		return c.JSON(jsonErr.Status, jsonErr)
	}

	order, err := services.AdminService.UpdateOrderStatus(&reqBody)
	if err != nil {
		jsonErr := api.InternalServerError(err)
		return c.JSON(jsonErr.Status, jsonErr)
	}

	res := api.SuccessResponse(order, constants.UpdateOrderStatusSuccessMsg)
	return c.JSON(res.Status, res)
}

func (h *adminHandler) AddProductQuantity(c echo.Context) error {
	var reqBody admindto.AddProductQuantityDTO
	if err := api.DecodeDTO(c, &reqBody); err != nil {
		jsonErr := api.BadRequestError(err)
		return c.JSON(jsonErr.Status, jsonErr)
	}

	err := services.AdminService.AddProductQuantity(&reqBody)
	if err != nil {
		jsonErr := api.InternalServerError(err)
		return c.JSON(jsonErr.Status, jsonErr)
	}

	res := api.SuccessResponse(nil, constants.UpdateProductQuantitySuccessMsg)
	return c.JSON(res.Status, res)
}
