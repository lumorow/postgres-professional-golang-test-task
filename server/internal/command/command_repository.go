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

func (r *repository) CreateCommand(ctx context.Context, command *Command) (*Command, error) {
	var lastInsertId int64
	query := "INSERT INTO commands(command, decription) VALUES ($1,$2) RETURNING id"
	err := r.db.QueryRowContext(ctx, query, command.CommandBody, command.Description).Scan(&lastInsertId)
	if err != nil {
		return nil, err
	}

	command.ID = lastInsertId
	return command, nil
}

func (r *repository) GetCommandById(ctx context.Context, id int64) (*Command, error) {
	c := Command{}
	query := "SELECT id, command, description FROM commands WHERE id = $1"
	err := r.db.QueryRowContext(ctx, query, id).Scan(&c.ID, &c.CommandBody, &c.Description)
	if err != nil {
		return nil, err
	}

	return &c, nil
}

func (r *repository) GetCommands(ctx context.Context) (*[]Command, error) {
	var cs []Command
	query := "SELECT id, command, description FROM commands"
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		c := Command{}
		err = rows.Scan(&c)
		if err != nil {
			return nil, err
		}
		cs = append(cs, c)
	}
	return &cs, nil
}
