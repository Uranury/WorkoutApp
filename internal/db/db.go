package db

import (
	"github.com/jmoiron/sqlx"
)

func InitDB(driverName, dsn string) (*sqlx.DB, error) {
	db, err := sqlx.Connect(driverName, dsn)
	if err != nil {
		return nil, err
	}
	return db, nil
}
