package command

import (
	"context"
	"database/sql"
)

type DBTX interface {
	NamedExecContext(ctx context.Context, query string, arg interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
	Query(query string, args ...interface{}) (*sql.Rows, error)
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
	Begin() (*sql.Tx, error)
}

type Repository struct {
	db DBTX
}

func NewRepository(db DBTX) *Repository {
	return &Repository{db: db}
}
