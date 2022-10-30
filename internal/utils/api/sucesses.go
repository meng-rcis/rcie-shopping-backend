package api

import "github.com/nuttchai/go-rest/internal/constants"

type APISuccess struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Result  interface{} `json:"result"`
}

func SuccessResponse(result interface{}, message ...string) *APISuccess {
	msg := constants.DefaultAPISuccessMsg
	if len(message) > 0 {
		msg = message[0]
	}

	return &APISuccess{
		Status:  200,
		Message: msg,
		Result:  result,
	}
}
