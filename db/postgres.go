package db

import (
	"database/sql"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func Connect() (*sql.DB, error) {
	dbUrl := "postgres://admin:secret@localhost:5432/commerce"
	db, err := sql.Open("pgx", dbUrl)

	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
