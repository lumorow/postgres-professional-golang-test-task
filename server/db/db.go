package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type database struct {
	db *sql.DB
}

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewDatabase(cfg Config) (*database, error) {
	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))
	//db, err := sql.Open("postgres", "postgresql://postgres:password@localhost:5432/chat?sslmode=disable")
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
