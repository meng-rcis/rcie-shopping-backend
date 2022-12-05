package repositories

import (
	"database/sql"

	"github.com/nuttchai/go-rest/internal/models"
	"github.com/nuttchai/go-rest/internal/types"
	"github.com/nuttchai/go-rest/internal/utils/context"
	"github.com/nuttchai/go-rest/internal/utils/db"
)

func (m *DBModel) GetCartItem(id string, filters ...*types.QueryFilter) (*models.CartItem, error) {
	ctx, cancel := context.WithTimeout(3)
	defer cancel()

	baseQuery := "select * from cart where id = $1"
	baseArgs := []interface{}{id}

	query, args := db.BuildQueryWithFilter(baseQuery, baseArgs, filters...)
	row := m.DB.QueryRowContext(ctx, query, args...)

	var cartItem models.CartItem
	err := row.Scan(
		&cartItem.Id,
		&cartItem.OwnerId,
		&cartItem.ProductId,
		&cartItem.Quantity,
		&cartItem.TotalPrice,
		&cartItem.CreatedAt,
		&cartItem.UpdatedAt,
	)

	return &cartItem, err
}

func (m *DBModel) GetAllCartItems(userId string) ([]*models.CartItem, error) {
	ctx, cancel := context.WithTimeout(3)
	defer cancel()

	query := "select * from cart where owner_id = $1"
	rows, err := m.DB.QueryContext(ctx, query, userId)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	cartItems := []*models.CartItem{}
	for rows.Next() {
		var cartItem models.CartItem
		if err := rows.Scan(
			&cartItem.Id,
			&cartItem.OwnerId,
			&cartItem.ProductId,
			&cartItem.Quantity,
			&cartItem.TotalPrice,
			&cartItem.CreatedAt,
			&cartItem.UpdatedAt,
		); err != nil {
			return nil, err
		}

		cartItems = append(cartItems, &cartItem)
	}

	return cartItems, nil
}

func (m *DBModel) AddCartItem(userId string, productId string, quantity int, totalPrice float64) (*models.CartItem, error) {
	ctx, cancel := context.WithTimeout(3)
	defer cancel()

	query := "insert into cart (owner_id, product_id, quantity, total_price) values ($1, $2, $3, $4) returning *"
	row := m.DB.QueryRowContext(ctx, query, userId, productId, quantity, totalPrice)

	var cartItem models.CartItem
	if err := row.Scan(
		&cartItem.Id,
		&cartItem.OwnerId,
		&cartItem.ProductId,
		&cartItem.Quantity,
		&cartItem.TotalPrice,
		&cartItem.CreatedAt,
		&cartItem.UpdatedAt,
	); err != nil {
		return nil, err
	}

	return &cartItem, nil
}

func (m *DBModel) UpdateCartItem(id string, quantity int, totalPrice float64) (*models.CartItem, error) {
	ctx, cancel := context.WithTimeout(3)
	defer cancel()

	query := "update cart set quantity = $1, total_price = $2 where id = $3 returning *"
	row := m.DB.QueryRowContext(ctx, query, quantity, totalPrice, id)

	var cartItem models.CartItem
	err := row.Scan(
		&cartItem.Id,
		&cartItem.OwnerId,
		&cartItem.ProductId,
		&cartItem.Quantity,
		&cartItem.TotalPrice,
		&cartItem.CreatedAt,
		&cartItem.UpdatedAt,
	)

	return &cartItem, err
}

func (m *DBModel) RemoveCartItem(id string) (sql.Result, error) {
	ctx, cancel := context.WithTimeout(3)
	defer cancel()

	query := `
		delete from cart 
		where id = $1
	`
	result, err := m.DB.ExecContext(ctx, query, id)

	return result, err
}
