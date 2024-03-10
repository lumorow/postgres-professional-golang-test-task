package command

import (
	"context"
	"pstgrprof/server/internal/entity"
)

func (r *Repository) GetCommandById(ctx context.Context, id int64) (*entity.Command, error) {
	c := entity.Command{}
	query := "SELECT id, script, description FROM commands WHERE id = $1"
	err := r.db.QueryRowContext(ctx, query, id).Scan(&c.ID, &c.Script, &c.Description)
	if err != nil {
		return nil, err
	}
	
	// Добавить к кеш значение
	// map[id] = entity.Command{}

	return &c, nil
}
