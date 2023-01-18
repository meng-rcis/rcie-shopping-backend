package handlers

import (
	"github.com/labstack/echo"
	"github.com/nuttchai/go-rest/internal/constants"
	admindto "github.com/nuttchai/go-rest/internal/dto/admin"
	"github.com/nuttchai/go-rest/internal/services"
	"github.com/nuttchai/go-rest/internal/types"
	"github.com/nuttchai/go-rest/internal/utils/api"
)

type adminHandler struct{}

type adminHandlerInterface interface {
	UpdateOrderStatus(c echo.Context) error
	AddProductQuantity(c echo.Context) error
	BulkUpdateProduct(c echo.Context) error
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

func (h *adminHandler) BulkUpdateProduct(c echo.Context) error {
	var reqBody admindto.BulkUpdateProductDTO
	if err := api.DecodeDTO(c, &reqBody); err != nil {
		jsonErr := api.BadRequestError(err)
		return c.JSON(jsonErr.Status, jsonErr)
	}

	searchQuery := types.SearchQuery{
		Keyword:   c.QueryParam("keyword"),
		ProductId: c.QueryParam("productId"),
		ShopId:    c.QueryParam("shopId"),
		Offset:    c.QueryParam("offset"),
		Limit:     c.QueryParam("limit"),
	}

	products, err := services.SearchService.SearchProduct(&searchQuery)
	if err != nil {
		jsonErr := api.InternalServerError(err)
		return c.JSON(jsonErr.Status, jsonErr)
	}

	for _, product := range products {
		product.Price = reqBody.Price
		product.Status = reqBody.Status
		product.Quantity = reqBody.Quantity

		if _, err := services.ProductService.UpdateProduct(product); err != nil {
			jsonErr := api.InternalServerError(err)
			return c.JSON(jsonErr.Status, jsonErr)
		}
	}

	res := api.SuccessResponse(nil, constants.BulkUpdateProductSuccessMsg)
	return c.JSON(res.Status, res)
}
