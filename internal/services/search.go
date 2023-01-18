package services

import (
	"github.com/nuttchai/go-rest/internal/models"
	"github.com/nuttchai/go-rest/internal/types"
	"github.com/nuttchai/go-rest/internal/utils/query"
)

type searchService struct {
	repo *Repository
}

type searchServiceInterface interface {
	SearchProduct(searchQuery *types.SearchQuery) ([]*models.Product, error)
}

var (
	SearchService searchServiceInterface
)

func init() {
	SearchService = &searchService{
		repo: &repo,
	}
}

func (s *searchService) SearchProduct(searchQuery *types.SearchQuery) ([]*models.Product, error) {
	filter := query.GenerateProductFilter(searchQuery)

	return s.repo.Models.DB.SearchProduct(
		searchQuery.Offset,
		searchQuery.Limit,
		filter...,
	)
}
