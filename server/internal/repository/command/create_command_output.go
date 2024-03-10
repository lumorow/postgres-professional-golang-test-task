package command

import (
	"context"
)

func (r *Repository) CreateCommandOutput(ctx context.Context, id int64, output string) error {
	var lastInsertId int64
	query := "INSERT INTO commands_output(id_command, output, time) VALUES ($1,$2, NOW()) RETURNING id"
	err := r.db.QueryRowContext(ctx, query, id, output).Scan(&lastInsertId)
	if err != nil {
		return err
	}

	return nil
}
