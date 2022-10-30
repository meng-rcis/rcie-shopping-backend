package repositories

import "database/sql"

type DBModel struct {
	SqlDB *sql.DB
}

type Models struct {
	DB DBModel
}

func InitModels(SqlDB *sql.DB) Models {
	return Models{
		DB: DBModel{
			SqlDB: SqlDB,
		},
	}
}
