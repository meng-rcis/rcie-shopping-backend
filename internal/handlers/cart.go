package handlers

import (
	"errors"
	"fmt"

	"github.com/labstack/echo"
	"github.com/nuttchai/go-rest/internal/constants"
	cartdto "github.com/nuttchai/go-rest/internal/dto/cart"
	shareddto "github.com/nuttchai/go-rest/internal/dto/shared"
	"github.com/nuttchai/go-rest/internal/services"
	"github.com/nuttchai/go-rest/internal/utils/api"
)

type cartHandler struct{}

type cartHandlerInterface interface {
	GetAllCartItems(c echo.Context) error
	AddCartItem(c echo.Context) error
	UpdateCartItem(c echo.Context) error
	RemoveCartItem(c echo.Context) error
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
	var reqBody cartdto.AddCartItemDTO
	err := api.DecodeDTO(c, &reqBody)
	if err != nil {
		jsonErr := api.BadRequestError(err)
		return c.JSON(jsonErr.Status, jsonErr)
	}

	item, err := services.CartService.AddCartItem(&reqBody)
	if err != nil {
		jsonErr := api.InternalServerError(err)
		return c.JSON(jsonErr.Status, jsonErr)
	}

	res := api.SuccessResponse(item, constants.AddCartItemSuccessMsg)
	return c.JSON(res.Status, res)
}

func (h *cartHandler) UpdateCartItem(c echo.Context) error {
	var reqBody cartdto.UpdateCartItemDTO
	err := api.DecodeDTO(c, &reqBody)
	if err != nil {
		jsonErr := api.BadRequestError(err)
		return c.JSON(jsonErr.Status, jsonErr)
	}

	item, err := services.CartService.UpdateCartItem(&reqBody)
	if err != nil {
		jsonErr := api.InternalServerError(err)
		return c.JSON(jsonErr.Status, jsonErr)
	}

	res := api.SuccessResponse(item, constants.UpdateCartItemSuccessMsg)
	return c.JSON(res.Status, res)
}

func (h *cartHandler) RemoveCartItem(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		jsonErr := api.BadRequestError(
			errors.New(
				fmt.Sprint(constants.MissingParamError, ": cart-id"),
			),
		)
		return c.JSON(jsonErr.Status, jsonErr)
	}

	err := services.CartService.RemoveCartItem(id)
	if err != nil {
		jsonErr := api.InternalServerError(err)
		return c.JSON(jsonErr.Status, jsonErr)
	}

	res := api.SuccessResponse(
		&shareddto.ValidatorResultDTO{
			IsSuccess: true,
			Action:    "RemoveCartItem",
		},
		constants.RemoveCartItemSuccessMsg)

	return c.JSON(res.Status, res)
}
