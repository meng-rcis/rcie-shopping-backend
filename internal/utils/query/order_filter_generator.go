package query

import "github.com/nuttchai/go-rest/internal/types"

func GenerateOrderFilter(orderQuery *types.OrderQuery) []*types.QueryFilter {
	filter := []*types.QueryFilter{}

	if orderQuery.UserId != "" {
		filter = append(filter, &types.QueryFilter{
			Field:    "o.owner_id",
			Operator: "=",
			Value:    orderQuery.UserId,
		})
	}

	if orderQuery.Status != "" {
		filter = append(filter, &types.QueryFilter{
			Field:    "os.name",
			Operator: "=",
			Value:    orderQuery.Status,
		})
	}

	return filter
}
