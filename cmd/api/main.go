package main

import (
	internal "github.com/nuttchai/go-rest/internal/api"
	"github.com/nuttchai/go-rest/internal/config"
)

func main() {
	config.App.Log("Start client...")
	internal.Client()
}
