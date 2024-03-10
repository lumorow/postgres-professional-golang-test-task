package command

import (
	"context"
	"database/sql"
	"github.com/jmoiron/sqlx"
	"pstgrprof/server/internal/entity"
)

func (r *Repository) GetCommands(ctx context.Context, ids []int64) (*[]entity.Command, error) {
	var cs []entity.Command
	query, args, err := sqlx.In("SELECT id, script, description FROM commands WHERE id IN(?)", ids)
	if err != nil {
		return nil, err
	}

	query = sqlx.Rebind(sqlx.DOLLAR, query)
	rows, err := r.db.Query(query, args...)
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

	if len(cs) == 0 {
		return nil, sql.ErrNoRows
	}
	return &cs, nil
}
