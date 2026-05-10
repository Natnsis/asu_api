package repository

import (
	"database/sql"
	"github.com/lib/pq"
)

func NewNeonDB (connStr string) (*sql.DB, error) {
	db, err := sql.open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	//test connection
	if err := db.Ping(); err != nil {
		return nil, error
	}
	return db, nil
}
