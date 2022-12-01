package handlers

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/labstack/echo"
	"github.com/nuttchai/go-rest/internal/constants"
	"github.com/nuttchai/go-rest/internal/dto/cart_dto"
	"github.com/nuttchai/go-rest/internal/services"
	"github.com/nuttchai/go-rest/internal/utils/api"
)

type cartHandler struct{}

type cartHandlerInterface interface {
	GetAllCartItems(c echo.Context) error
	AddCartItem(c echo.Context) error
	UpdateCartItem(c echo.Context) error
}

var (
	CartHandler cartHandlerInterface
)

func init() {
	CartHandler = &cartHandler{}
}

func (h *cartHandler) GetAllCartItems(c echo.Context) error {
	userId := c.QueryParam("userId")
	if userId == "" {
		jsonErr := api.BadRequestError(
			errors.New(
				fmt.Sprint(constants.MissingParamError, ": userId"),
			),
		)
		return c.JSON(jsonErr.Status, jsonErr)
	}

	items, err := services.CartService.GetAllCartItems(userId)
	if err != nil {
		jsonErr := api.InternalServerError(err)
		return c.JSON(jsonErr.Status, jsonErr)
	}

	res := api.SuccessResponse(items, constants.GetCartItemsSuccessMsg)
	return c.JSON(res.Status, res)
}

func (h *cartHandler) AddCartItem(c echo.Context) error {
	var reqBody *cart_dto.AddCartItemDTO
	err := json.NewDecoder(c.Request().Body).Decode(&reqBody)
	if err != nil {
		jsonErr := api.BadRequestError(err)
		return c.JSON(jsonErr.Status, jsonErr)
	}

	item, err := services.CartService.AddCartItem(reqBody)
	if err != nil {
		jsonErr := api.InternalServerError(err)
		return c.JSON(jsonErr.Status, jsonErr)
	}

	res := api.SuccessResponse(item, constants.AddCartItemSuccessMsg)
	return c.JSON(res.Status, res)
}

func (h *cartHandler) UpdateCartItem(c echo.Context) error {
	var reqBody *cart_dto.UpdateCartItemDTO
	err := json.NewDecoder(c.Request().Body).Decode(&reqBody)
	if err != nil {
		jsonErr := api.BadRequestError(err)
		return c.JSON(jsonErr.Status, jsonErr)
	}

	item, err := services.CartService.UpdateCartItem(reqBody)
	if err != nil {
		jsonErr := api.InternalServerError(err)
		return c.JSON(jsonErr.Status, jsonErr)
	}

	res := api.SuccessResponse(item, constants.UpdateCartItemSuccessMsg)
	return c.JSON(res.Status, res)
}
