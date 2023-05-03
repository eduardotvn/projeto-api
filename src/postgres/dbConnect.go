package postgres

import (
	"database/sql"
	"os"

	_ "github.com/lib/pq"
)

func DBConnect() (*sql.DB, error) {
	psqlconn := os.Getenv("DB_STRING")

	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
