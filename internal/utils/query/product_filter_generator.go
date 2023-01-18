package query

import "github.com/nuttchai/go-rest/internal/types"

func GenerateProductFilter(searchQuery *types.SearchQuery) []*types.QueryFilter {
	filter := []*types.QueryFilter{}

	if searchQuery.Keyword != "" {
		filter = append(filter, &types.QueryFilter{
			Field:    "p.name",
			Operator: "LIKE",
			Value:    "%" + searchQuery.Keyword + "%",
		})
	}

	if searchQuery.ProductId != "" {
		filter = append(filter, &types.QueryFilter{
			Field:    "p.id",
			Operator: "=",
			Value:    searchQuery.ShopId,
		})
	}

	if searchQuery.ShopId != "" {
		filter = append(filter, &types.QueryFilter{
			Field:    "p.shop_id",
			Operator: "=",
			Value:    searchQuery.ShopId,
		})
	}

	return filter
}
