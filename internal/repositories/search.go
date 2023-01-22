package repositories

import (
	"fmt"

	"github.com/nuttchai/go-rest/internal/models"
	"github.com/nuttchai/go-rest/internal/types"
	"github.com/nuttchai/go-rest/internal/utils/context"
	"github.com/nuttchai/go-rest/internal/utils/db"
)

func (m *DBModel) SearchProduct(
	offset string,
	limit string,
	isHiddenRequired bool,
	filters ...*types.QueryFilter,
) ([]*models.Product, error) {
	ctx, cancel := context.WithTimeout(3)
	defer cancel()

	status := "('Shown')"
	if isHiddenRequired {
		status = "('Shown', 'Hidden')"
	}

	baseArgs := []interface{}{status}
	baseQuery := `
		select p.id, p.name, p.description, p.price, p.quantity, p.shop_id, ps.name, p.created_at, p.updated_at
		from product as p 
		join product_status as ps on p.status_id = ps.id
		where ps.name = $1
	`

	query, args := db.BuildQueryWithFilter(baseQuery, baseArgs, filters...)
	query += " order by p.name asc"

	if offset != "" && limit != "" {
		limitIndex, offsetIndex := len(args)+1, len(args)+2
		query += fmt.Sprintf(" limit $%d offset $%d", limitIndex, offsetIndex)
		args = append(args, limit, offset)
	}

	rows, err := m.DB.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	products := []*models.Product{}
	for rows.Next() {
		var product models.Product
		if err := rows.Scan(
			&product.Id,
			&product.Name,
			&product.Description,
			&product.Price,
			&product.Quantity,
			&product.ShopId,
			&product.Status,
			&product.CreatedAt,
			&product.UpdatedAt,
		); err != nil {
			return nil, err
		}

		products = append(products, &product)
	}

	return products, nil
}
