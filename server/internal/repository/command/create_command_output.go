package command

import "context"

func (r *Repository) CreateCommandOutput(ctx context.Context, id int64, output string) error {
	query := "INSERT INTO commands_output(id_command, output, time) VALUES ($1,$2, NOW())"
	err := r.db.QueryRowContext(ctx, query, id, output).Scan()
	if err != nil {
		return err
	}

	return nil
}
