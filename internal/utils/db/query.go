package db

import (
	"fmt"

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
		joinQuery := "and"
		if argsIndex == 1 {
			joinQuery = "where"
		}

		query += fmt.Sprintf(" %s %s %s $%d", joinQuery, filter.Field, filter.Operator, argsIndex)
		args = append(args, filter.Value)
		argsIndex++
	}

	return query, args
}
