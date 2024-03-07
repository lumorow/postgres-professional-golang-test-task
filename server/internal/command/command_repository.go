package command

import (
	"context"
	"database/sql"
)

type DBTX interface {
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

type repository struct {
	db DBTX
}

func NewRepository(db DBTX) Repository {
	return &repository{db: db}
}

func (r *repository) CreateCommand(ctx context.Context, user *Command) (*Command, error) {
	//var lastInsertId int64
	return user, nil
}

func (r *repository) GetCommandById(ctx context.Context, id int64) (*Command, error) {
	c := Command{}
	return &c, nil
}

func (r *repository) GetAllCommands(ctx context.Context) ([]*Command, error) {
	//c := Command{}
	return nil, nil
}
