package api

import (
	"encoding/json"

	"github.com/labstack/echo"
	"github.com/nuttchai/go-rest/internal/utils/validators"
)

func DecodeDTO(c echo.Context, ptr any) error {
	decoder := json.NewDecoder(c.Request().Body)
	if err := decoder.Decode(ptr); err != nil {
		return err
	}

	return validators.ValidateStruct(ptr)
}
