package repositories

import "database/sql"

type DBModel struct {
	DB *sql.DB
}

type Models struct {
	DB DBModel
}

func InitModels(db *sql.DB) Models {
	return Models{
		DB: DBModel{
			DB: db,
		},
	}
}
