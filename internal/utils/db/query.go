package db

import (
	"strconv"

	"github.com/nuttchai/go-rest/internal/types"
)

// query is a base SQL query
// args is a list of existing arguments that are passed in base SQL query
func BuildQueryWithFilter(
	query string,
	args []interface{},
	filters ...*types.QueryFilter,
) (string, []interface{}) {
	argsIndex := len(args) + 1

	for _, filter := range filters {
		joinQuery := " and "
		if argsIndex == 1 {
			joinQuery = " where "
		}

		query += joinQuery + filter.Field + " = $" + strconv.Itoa(argsIndex)
		args = append(args, filter.Value)
		argsIndex++
	}

	return query, args
}
