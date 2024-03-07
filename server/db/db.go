package db

import (
	"database/sql"
	_ "github.com/lib/pq"
)

type database struct {
	db *sql.DB
}

func NewDatabase() (*database, error) {
	db, err := sql.Open("postgres", "postgresql://postgres:password@localhost:5432/chat?sslmode=disable")
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &database{
		db: db,
	}, nil
}

func (d *database) Close() {
	d.db.Close()
}

func (d *database) GetDB() *sql.DB {
	return d.db
}
