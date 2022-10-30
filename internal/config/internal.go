package config

import "github.com/nuttchai/go-rest/internal/repositories"

type APIConfig struct {
	Port string
	Env  string
	Db   struct {
		Dsn string
	}
}

type AppConfig struct {
	APIConfig APIConfig
	Models    repositories.Models
}
