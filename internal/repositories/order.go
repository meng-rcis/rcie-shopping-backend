package repositories

import (
	"github.com/nuttchai/go-rest/internal/models"
	"github.com/nuttchai/go-rest/internal/types"
	"github.com/nuttchai/go-rest/internal/utils/context"
	"github.com/nuttchai/go-rest/internal/utils/db"
)

func (m *DBModel) GetOrders(filters ...*types.QueryFilter) ([]*models.Order, error) {
	ctx, cancel := context.WithTimeout(3)
	defer cancel()

	// we need to pass table name as "order", not order with not "", because order is a reserved word in postgres
	baseQuery := `
		select o.id, o.owner_id, o.product_id, os.name, o.quantity, o.total_price, o.created_at, o.updated_at
		from "order" as o
		join order_status as os on o.status_id = os.id
	`
	baseArgs := []interface{}{}

	query, args := db.BuildQueryWithFilter(baseQuery, baseArgs, filters...)
	rows, err := m.DB.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	orders := []*models.Order{}
	for rows.Next() {
		var order models.Order
		if err := rows.Scan(
			&order.Id,
			&order.OwnerId,
			&order.ProductId,
			&order.Status,
			&order.Quantity,
			&order.TotalPrice,
			&order.CreatedAt,
			&order.UpdatedAt,
		); err != nil {
			return nil, err
		}

		orders = append(orders, &order)
	}

	return orders, nil
}
