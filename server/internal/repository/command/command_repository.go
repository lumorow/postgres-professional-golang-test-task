package command

import (
	"context"
	"database/sql"
	"pstgrprof/server/internal/command"
)

type DBTX interface {
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

type Repository struct {
	db DBTX
}

func NewRepository(db DBTX) *Repository {
	return &Repository{db: db}
}

func (r *Repository) CreateCommand(ctx context.Context, command *command.Command) (*command.Command, error) {
	var lastInsertId int64
	query := "INSERT INTO commands(script, description) VALUES ($1,$2) RETURNING id"
	err := r.db.QueryRowContext(ctx, query, command.Script, command.Description).Scan(&lastInsertId)
	if err != nil {
		return nil, err
	}

	command.ID = lastInsertId
	return command, nil
}

func (r *Repository) GetCommandById(ctx context.Context, id int64) (*command.Command, error) {
	c := command.Command{}
	query := "SELECT id, script, description FROM commands WHERE id = $1"
	err := r.db.QueryRowContext(ctx, query, id).Scan(&c.ID, &c.Script, &c.Description)
	if err != nil {
		return nil, err
	}

	return &c, nil
}

func (r *Repository) GetAllCommands(ctx context.Context) (*[]command.Command, error) {
	var cs []command.Command
	query := "SELECT id, script, description FROM commands"
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		c := command.Command{}
		err = rows.Scan(&c.ID, &c.Script, &c.Description)
		if err != nil {
			return nil, err
		}
		cs = append(cs, c)
	}
	return &cs, nil
}
