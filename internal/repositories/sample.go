package repositories

import (
	"database/sql"

	"github.com/nuttchai/go-rest/internal/models"
	"github.com/nuttchai/go-rest/internal/utils/context"
)

func (m *DBModel) GetSample(id string) (*models.Sample, error) {
	ctx, cancel := context.WithTimeout(3)
	defer cancel()

	query := `select * from samples where id = $1`
	row := m.DB.QueryRowContext(ctx, query, id)

	var sample models.Sample
	err := row.Scan(
		&sample.Id,
		&sample.Name,
	)

	return &sample, err
}

func (m *DBModel) CreateSample(s *models.NewSample) (int, error) {
	ctx, cancel := context.WithTimeout(3)
	defer cancel()

	var id int
	query := `
		insert into samples (name, desc)
		values ($1, $2)
		returning id
	`
	err := m.DB.QueryRowContext(ctx, query, s.Name, s.Desc).Scan(&id)

	return id, err
}

func (m *DBModel) UpdateSample(s *models.Sample) (sql.Result, error) {
	ctx, cancel := context.WithTimeout(3)
	defer cancel()

	query := `
		update samples set name = $1, desc = $2
		where id = $3
	`
	result, err := m.DB.ExecContext(ctx, query, s.Name, s.Desc, s.Id)

	return result, err
}

func (m *DBModel) DeleteSample(id string) (sql.Result, error) {
	ctx, cancel := context.WithTimeout(3)
	defer cancel()

	query := `
		delete from samples 
		where id = $1
	`
	result, err := m.DB.ExecContext(ctx, query, id)

	return result, err
}

func (m *DBModel) EmptySampleDesc(id string) (sql.Result, error) {
	ctx, cancel := context.WithTimeout(3)
	defer cancel()

	query := `
		update samples set desc = '-'
		where id = $1
	`
	result, err := m.DB.ExecContext(ctx, query, id)

	return result, err
}
