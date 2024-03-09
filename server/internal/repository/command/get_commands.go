package command

import (
	"context"
	"pstgrprof/server/internal/entity"
)

func (r *Repository) GetAllCommands(ctx context.Context) (*[]entity.Command, error) {
	var cs []entity.Command
	query := "SELECT id, script, description FROM commands"
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		c := entity.Command{}
		err = rows.Scan(&c.ID, &c.Script, &c.Description)
		if err != nil {
			return nil, err
		}
		cs = append(cs, c)
	}
	return &cs, nil
}
