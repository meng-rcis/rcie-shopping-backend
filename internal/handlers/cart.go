package handlers

import "github.com/labstack/echo"

type cartHandler struct{}

type cartHandlerInterface interface {
	GetCartItem(c echo.Context) error
}

var (
	CartHandler cartHandlerInterface
)

func init() {
	CartHandler = &cartHandler{}
}

func (h *cartHandler) GetCartItem(c echo.Context) error {
	return nil
}
