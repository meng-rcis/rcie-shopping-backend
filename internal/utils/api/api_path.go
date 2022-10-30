package api

import "github.com/nuttchai/go-rest/internal/constants"

func CreatePath(path string) string {
	return constants.BasePath + path
}
