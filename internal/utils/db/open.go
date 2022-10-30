package db

import (
	"database/sql"

	"github.com/nuttchai/go-rest/internal/config"
	"github.com/nuttchai/go-rest/internal/utils/context"
)

func OpenSqlDB(driver string, cfg config.APIConfig) (*sql.DB, error) {
	db, err := sql.Open(driver, cfg.Db.Dsn)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(5)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return db, nil
}
