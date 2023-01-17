package services

import (
	"github.com/nuttchai/go-rest/internal/models"
	"github.com/nuttchai/go-rest/internal/types"
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
	searchFilter := []*types.QueryFilter{}
	if searchQuery.Keyword != "" {
		searchFilter = append(searchFilter, &types.QueryFilter{
			Field:    "p.name",
			Operator: "LIKE",
			Value:    "%" + searchQuery.Keyword + "%",
		})
	}

	if searchQuery.ShopId != "" {
		searchFilter = append(searchFilter, &types.QueryFilter{
			Field:    "p.shop_id",
			Operator: "=",
			Value:    searchQuery.ShopId,
		})
	}

	return s.repo.Models.DB.SearchProduct(
		searchQuery.Offset,
		searchQuery.Limit,
		searchFilter...,
	)
}
