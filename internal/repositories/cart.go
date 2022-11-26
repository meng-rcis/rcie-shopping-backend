package repositories

import (
	"github.com/nuttchai/go-rest/internal/models"
	"github.com/nuttchai/go-rest/internal/shared/console"
	"github.com/nuttchai/go-rest/internal/utils/context"
)

func (m *DBModel) GetItems(userId string) ([]*models.CartItem, error) {
	ctx, cancel := context.WithTimeout(3)
	defer cancel()
	console.App.Log("GetItems: " + userId)
	query := "select * from cart where owner_id = $1"
	rows, err := m.DB.QueryContext(ctx, query, userId)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	cartItems := []*models.CartItem{}
	for rows.Next() {
		var cartItem models.CartItem
		err := rows.Scan(
			&cartItem.Id,
			&cartItem.OwnerId,
			&cartItem.ProductId,
			&cartItem.Quantity,
			&cartItem.TotalPrice,
			&cartItem.CreatedAt,
			&cartItem.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		cartItems = append(cartItems, &cartItem)
	}

	return cartItems, nil
}
