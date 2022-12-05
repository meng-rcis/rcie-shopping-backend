package handlers

import (
	"errors"
	"fmt"

	"github.com/labstack/echo"
	"github.com/nuttchai/go-rest/internal/constants"
	"github.com/nuttchai/go-rest/internal/services"
	"github.com/nuttchai/go-rest/internal/utils/api"
)

type productHandler struct{}

type productHandlerInterface interface {
	GetProductById(c echo.Context) error
}

var (
	ProductHandler productHandlerInterface
)

func init() {
	ProductHandler = &productHandler{}
}

func (h *productHandler) GetProductById(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		jsonErr := api.BadRequestError(
			errors.New(
				fmt.Sprint(constants.MissingParamError, ": product-id"),
			),
		)
		return c.JSON(jsonErr.Status, jsonErr)
	}

	product, err := services.ProductService.GetProductById(id)
	if err != nil {
		jsonErr := api.InternalServerError(err)
		return c.JSON(jsonErr.Status, jsonErr)
	}

	res := api.SuccessResponse(product, constants.GetProductSuccessMsg)
	return c.JSON(res.Status, res)
}
