package api

import "github.com/nuttchai/go-rest/internal/constants"

type APIError struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Error   string `json:"error"`
}

func CustomError(err error, status int, message ...string) *APIError {
	msg := constants.DefaultAPIErrorMsg
	if len(message) > 0 {
		msg = message[0]
	}

	return &APIError{
		Status:  status,
		Message: msg,
		Error:   err.Error(),
	}
}

func InternalServerError(err error) *APIError {
	return &APIError{
		Status:  500,
		Message: "InternalServerError",
		Error:   err.Error(),
	}
}

func NotFoundError(err error) *APIError {
	return &APIError{
		Status:  404,
		Message: "NotFoundError",
		Error:   err.Error(),
	}
}
