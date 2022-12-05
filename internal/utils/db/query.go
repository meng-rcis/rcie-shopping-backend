package db

import (
	"strconv"

	"github.com/nuttchai/go-rest/internal/types"
)

func BuildQueryWithFilter(
	query string,
	args []interface{},
	filters ...*types.QueryFilter,
) (string, []interface{}) {
	if len(filters) > 0 {
		for index, filter := range filters {
			query += " and " + filter.Key + " = $" + strconv.Itoa(index+2)
			args = append(args, filter.Value)
		}
	}

	return query, args
}
