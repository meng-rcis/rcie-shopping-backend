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

	// we need to pass table name as "order", not order with no "", because order is a reserved word in postgres
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

func (m *DBModel) CreateOrder(order *models.Order) (*models.Order, error) {
	ctx, cancel := context.WithTimeout(3)
	defer cancel()

	query := `
		insert into "order" (owner_id, product_id, quantity, total_price)
		values ($1, $2, $3, $4)
		returning *
	`
	row := m.DB.QueryRowContext(
		ctx,
		query,
		order.OwnerId,
		order.ProductId,
		order.Quantity,
		order.TotalPrice,
	)

	var newOrder models.Order
	err := row.Scan(
		&newOrder.Id,
		&newOrder.OwnerId,
		&newOrder.ProductId,
		&newOrder.Status,
		&newOrder.Quantity,
		&newOrder.TotalPrice,
		&newOrder.CreatedAt,
		&newOrder.UpdatedAt,
	)

	return &newOrder, err
}

func (m *DBModel) UpdateOrder(order *models.Order) (*models.Order, error) {
	ctx, cancel := context.WithTimeout(3)
	defer cancel()

	query := `
		update "order"
		set status_id = (
			select id from order_status where name = $1
		)
		where id = $2
		returning *
	`
	row := m.DB.QueryRowContext(
		ctx,
		query,
		order.Status,
		order.Id,
	)

	var updatedOrder models.Order
	err := row.Scan(
		&updatedOrder.Id,
		&updatedOrder.OwnerId,
		&updatedOrder.ProductId,
		&updatedOrder.Status,
		&updatedOrder.Quantity,
		&updatedOrder.TotalPrice,
		&updatedOrder.CreatedAt,
		&updatedOrder.UpdatedAt,
	)

	return &updatedOrder, err
}
