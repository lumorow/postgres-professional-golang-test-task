package command

import (
	"context"
)

func (r *Repository) DeleteCommandById(ctx context.Context, id int64) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	deleteOutputsCommandQuery := "DELETE FROM commands_output WHERE id = $1"
	_, err = tx.ExecContext(ctx, deleteOutputsCommandQuery, id)

	if err != nil {
		tx.Rollback()
		return err
	}

	deleteCommandQuery := "DELETE FROM commands WHERE id = $1"
	_, err = tx.ExecContext(ctx, deleteCommandQuery, id)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}
