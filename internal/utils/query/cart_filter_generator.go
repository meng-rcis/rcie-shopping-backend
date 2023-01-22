package query

import "github.com/nuttchai/go-rest/internal/types"

func GenerateCartFilter(cartQuery *types.CartQuery) []*types.QueryFilter {
	filter := []*types.QueryFilter{}

	if cartQuery.UserId != "" {
		filter = append(filter, &types.QueryFilter{
			Field:    "owner_id",
			Operator: "=",
			Value:    cartQuery.UserId,
		})
	}

	return filter
}
