package repositories

import (
	"github.com/nuttchai/go-rest/internal/models"
	"github.com/nuttchai/go-rest/internal/utils/context"
)

func (m *DBModel) GetAllCartProducts(userId string) ([]*models.CartItem, error) {
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

func (m *DBModel) AddCartProduct(userId string, productId string, quantity int, total_price float64) (*models.CartItem, error) {
	ctx, cancel := context.WithTimeout(3)
	defer cancel()

	query := "insert into cart (owner_id, product_id, quantity, total_price) values ($1, $2, $3, $4) returning *"
	row := m.DB.QueryRowContext(ctx, query, userId, productId, quantity, total_price)

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
