package services

import (
	"github.com/nuttchai/go-rest/internal/config"
	"github.com/nuttchai/go-rest/internal/repositories"
)

var repo Repository

type Repository struct {
	Models *repositories.Models
}

func InitRepo(appConfig *config.AppConfig) Repository {
	return Repository{
		Models: &appConfig.Models,
	}
}

func InitServices(r Repository) {
	repo = r
}
