package repositories

import (
	"database/sql"

	"github.com/nuttchai/go-rest/internal/models"
	"github.com/nuttchai/go-rest/internal/utils/context"
)

func (m *DBModel) GetProduct(id string) (*models.Product, error) {
	ctx, cancel := context.WithTimeout(3)
	defer cancel()

	query := `
		select p.id, p.name, p.description, p.price, p.quantity, p.shop_id, ps.name, p.created_at, p.updated_at
		from product as p 
		join product_status as ps on p.status_id = ps.id 
		where p.id = $1
	`
	row := m.DB.QueryRowContext(ctx, query, id)

	var product models.Product
	err := row.Scan(
		&product.Id,
		&product.Name,
		&product.Description,
		&product.Price,
		&product.Quantity,
		&product.ShopId,
		&product.Status,
		&product.CreatedAt,
		&product.UpdatedAt,
	)

	return &product, err
}

func (m *DBModel) UpdateProduct(product *models.Product) (*models.Product, error) {
	ctx, cancel := context.WithTimeout(3)
	defer cancel()

	query := `
		update product
		set name = $1, description = $2, price = $3, quantity = $4, shop_id = $5, status_id = (
			select id from product_status where name = $6
		)
		where id = $7
		returning id, name, description, price, quantity, shop_id, $6, created_at, updated_at
	`

	row := m.DB.QueryRowContext(
		ctx,
		query,
		product.Name,
		product.Description,
		product.Price,
		product.Quantity,
		product.ShopId,
		product.Status,
		product.Id,
	)

	var updatedProduct models.Product
	err := row.Scan(
		&updatedProduct.Id,
		&updatedProduct.Name,
		&updatedProduct.Description,
		&updatedProduct.Price,
		&updatedProduct.Quantity,
		&updatedProduct.ShopId,
		&updatedProduct.Status,
		&updatedProduct.CreatedAt,
		&updatedProduct.UpdatedAt,
	)

	return &updatedProduct, err
}

func (m *DBModel) AddProductQuantity(id string, quantity int) (sql.Result, error) {
	ctx, cancel := context.WithTimeout(3)
	defer cancel()

	query := `
		update product
		set quantity = quantity + $1
		where id = $2
	`
	return m.DB.ExecContext(ctx, query, quantity, id)
}

func (m *DBModel) DeductProductQuantity(id string, quantity int) (sql.Result, error) {
	ctx, cancel := context.WithTimeout(3)
	defer cancel()

	query := `
		update product 
		set quantity = quantity - $1 
		where id = $2
	`
	return m.DB.ExecContext(ctx, query, quantity, id)
}
