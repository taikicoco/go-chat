package repository

import (
	"errors"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func Open() (*sqlx.DB, error) {
	dsn := os.Getenv("DATASOURCE")
	if dsn == "" {
		return nil, errors.New("missing DATASOURCE environment variable")
	}

	db, err := sqlx.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
