package command

import (
	"context"
	"pstgrprof/server/internal/entity"
)

func (r *Repository) CreateCommand(ctx context.Context, command *entity.Command) (*entity.Command, error) {
	var lastInsertId int64
	query := "INSERT INTO commands(script, description) VALUES ($1,$2) RETURNING id"
	err := r.db.QueryRowContext(ctx, query, command.Script, command.Description).Scan(&lastInsertId)
	if err != nil {
		return nil, err
	}

	command.ID = lastInsertId
	return command, nil
}
