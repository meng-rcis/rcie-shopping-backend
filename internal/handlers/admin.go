package handlers

import (
	"encoding/json"

	"github.com/labstack/echo"
	"github.com/nuttchai/go-rest/internal/constants"
	admindto "github.com/nuttchai/go-rest/internal/dto/admin"
	"github.com/nuttchai/go-rest/internal/services"
	"github.com/nuttchai/go-rest/internal/utils/api"
	"github.com/nuttchai/go-rest/internal/utils/validators"
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
	var reqBody *admindto.UpdateOrderStatusDTO
	err := json.NewDecoder(c.Request().Body).Decode(&reqBody)
	if err != nil {
		jsonErr := api.BadRequestError(err)
		return c.JSON(jsonErr.Status, jsonErr)
	}

	if err := validators.ValidateStruct(reqBody); err != nil {
		jsonErr := api.BadRequestError(err)
		return c.JSON(jsonErr.Status, jsonErr)
	}

	order, err := services.AdminService.UpdateOrderStatus(reqBody)
	if err != nil {
		jsonErr := api.InternalServerError(err)
		return c.JSON(jsonErr.Status, jsonErr)
	}

	res := api.SuccessResponse(order, constants.UpdateOrderStatusSuccessMsg)
	return c.JSON(res.Status, res)
}

func (h *adminHandler) AddProductQuantity(c echo.Context) error {
	var reqBody *admindto.AddProductQuantityDTO
	err := json.NewDecoder(c.Request().Body).Decode(&reqBody)
	if err != nil {
		jsonErr := api.BadRequestError(err)
		return c.JSON(jsonErr.Status, jsonErr)
	}

	if err := validators.ValidateStruct(reqBody); err != nil {
		jsonErr := api.BadRequestError(err)
		return c.JSON(jsonErr.Status, jsonErr)
	}

	err = services.AdminService.AddProductQuantity(reqBody)
	if err != nil {
		jsonErr := api.InternalServerError(err)
		return c.JSON(jsonErr.Status, jsonErr)
	}

	res := api.SuccessResponse(nil, constants.UpdateProductQuantitySuccessMsg)
	return c.JSON(res.Status, res)
}
