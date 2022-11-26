package handlers

import (
	"errors"
	"fmt"

	"github.com/labstack/echo"
	"github.com/nuttchai/go-rest/internal/constants"
	"github.com/nuttchai/go-rest/internal/services"
	"github.com/nuttchai/go-rest/internal/utils/api"
)

type cartHandler struct{}

type cartHandlerInterface interface {
	GetItems(c echo.Context) error
}

var (
	CartHandler cartHandlerInterface
)

func init() {
	CartHandler = &cartHandler{}
}

func (h *cartHandler) GetItems(c echo.Context) error {
	userId := c.QueryParam("userId")
	if userId == "" {
		jsonErr := api.BadRequestError(
			errors.New(
				fmt.Sprint(constants.MissingParamError, ": userId"),
			),
		)
		return c.JSON(jsonErr.Status, jsonErr)
	}

	items, err := services.CartService.GetItems(userId)
	if err != nil {
		jsonErr := api.InternalServerError(err)
		return c.JSON(jsonErr.Status, jsonErr)
	}

	res := api.SuccessResponse(items, constants.GetCartItemsSuccessMsg)
	return c.JSON(res.Status, res)
}
