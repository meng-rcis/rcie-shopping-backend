package main

import (
	internal "github.com/nuttchai/go-rest/internal/api"
	"github.com/nuttchai/go-rest/internal/shared/console"
)

func main() {
	console.App.Log("Start client...")
	internal.Client()
}
