package handlers

import (
	"github.com/labstack/echo"
	"github.com/nuttchai/go-rest/internal/constants"
	"github.com/nuttchai/go-rest/internal/services"
	"github.com/nuttchai/go-rest/internal/types"
	"github.com/nuttchai/go-rest/internal/utils/api"
)

type searchHandler struct{}

type searchHandlerInterface interface {
	SearchProduct(c echo.Context) error
}

var (
	SearchHandler searchHandlerInterface
)

func init() {
	SearchHandler = &searchHandler{}
}

func (h *searchHandler) SearchProduct(c echo.Context) error {
	searchQuery := types.SearchQuery{
		Keyword: c.QueryParam("keyword"),
		ShopId:  c.QueryParam("shopId"),
		Offset:  c.QueryParam("offset"),
	}

	limit := c.QueryParam("limit")
	if limit != "" {
		searchQuery.Limit = limit
	}

	searchResult, err := services.SearchService.SearchProduct(&searchQuery)
	if err != nil {
		jsonErr := api.InternalServerError(err)
		return c.JSON(jsonErr.Status, jsonErr)
	}

	res := api.SuccessResponse(searchResult, constants.SearchProductSuccessMsg)
	return c.JSON(res.Status, res)
}
