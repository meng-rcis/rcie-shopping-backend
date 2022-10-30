package handlers

import (
	"encoding/json"
	"strconv"

	"github.com/labstack/echo"
	"github.com/nuttchai/go-rest/internal/constants"
	"github.com/nuttchai/go-rest/internal/models"
	"github.com/nuttchai/go-rest/internal/services"
	"github.com/nuttchai/go-rest/internal/utils/api"
	jsonGen "github.com/nuttchai/go-rest/internal/utils/json"
)

type sampleHandler struct{}

type sampleHandlerInterface interface {
	GetSample(c echo.Context) error
	CreateSample(c echo.Context) error
	UpdateSample(c echo.Context) error
	DeleteSample(c echo.Context) error
	EmptySampleDesc(c echo.Context) error
}

var (
	SampleHandler sampleHandlerInterface
)

func init() {
	SampleHandler = &sampleHandler{}
}

func (h *sampleHandler) GetSample(c echo.Context) error {
	id := c.Param("id")
	sample, err := services.SampleService.GetSample(id)
	if err != nil {
		jsonErr := api.InternalServerError(err)
		return c.JSON(jsonErr.Status, jsonErr)
	}

	res := api.SuccessResponse(sample, constants.GetSampleSuccessMsg)
	return c.JSON(res.Status, res)
}

func (h *sampleHandler) CreateSample(c echo.Context) error {
	var sample *models.NewSample
	err := json.NewDecoder(c.Request().Body).Decode(&sample)
	if err != nil {
		jsonErr := api.CustomError(err, 500, constants.DecodingJSONError)
		return c.JSON(jsonErr.Status, jsonErr)
	}

	sampleId, err := services.SampleService.CreateSample(sample)
	if err != nil {
		jsonErr := api.InternalServerError(err)
		return c.JSON(jsonErr.Status, jsonErr)
	}

	json := &models.CreatedSample{
		Id: sampleId,
	}
	res := api.SuccessResponse(json, constants.CreateSampleSuccessMsg)
	return c.JSON(res.Status, res)
}

func (h *sampleHandler) UpdateSample(c echo.Context) error {
	var sample *models.Sample
	err := json.NewDecoder(c.Request().Body).Decode(&sample)
	if err != nil {
		jsonErr := api.CustomError(err, 500, constants.DecodingJSONError)
		return c.JSON(jsonErr.Status, jsonErr)
	}

	err = services.SampleService.UpdateSample(sample)
	if err != nil {
		jsonErr := jsonGen.GenerateNotFoundIfErrorMatched(err, constants.SampleNotFound)
		return c.JSON(jsonErr.Status, jsonErr)
	}

	json := &models.UpdatedSample{Id: sample.Id, Updated: true}
	res := api.SuccessResponse(json, constants.UpdateSampleSuccessMsg)
	return c.JSON(res.Status, res)
}

func (h *sampleHandler) DeleteSample(c echo.Context) error {
	id := c.Param("id")
	err := services.SampleService.DeleteSample(id)
	if err != nil {
		jsonErr := jsonGen.GenerateNotFoundIfErrorMatched(err, constants.SampleNotFound)
		return c.JSON(jsonErr.Status, jsonErr)
	}

	intId, err := strconv.Atoi(id)
	if err != nil {
		jsonErr := api.CustomError(err, 500, constants.ConvertStringToIntError)
		return c.JSON(jsonErr.Status, jsonErr)
	}

	json := &models.DeleteSample{Id: intId, Deleted: true}
	res := api.SuccessResponse(json, constants.DeleteSampleSuccessMsg)
	return c.JSON(res.Status, res)
}

func (h *sampleHandler) EmptySampleDesc(c echo.Context) error {
	id := c.Param("id")
	err := services.SampleService.EmptySampleDesc(id)
	if err != nil {
		jsonErr := jsonGen.GenerateNotFoundIfErrorMatched(err, constants.SampleNotFound)
		return c.JSON(jsonErr.Status, jsonErr)
	}

	intId, err := strconv.Atoi(id)
	if err != nil {
		jsonErr := api.CustomError(err, 500, constants.ConvertStringToIntError)
		return c.JSON(jsonErr.Status, jsonErr)
	}

	json := &models.UpdatedSample{Id: intId, Updated: true}
	res := api.SuccessResponse(json, constants.UpdateSampleSuccessMsg)
	return c.JSON(res.Status, res)
}
