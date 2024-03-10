package command

import (
	"context"
	"pstgrprof/server/internal/entity"
	"strconv"
)

func (s *Service) GetCommandById(ctx context.Context, id string) (*entity.Command, error) {
	i, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, err
	}

	r, err := s.Repository.GetCommandById(ctx, i)
	if err != nil {
		return nil, err
	}

	return r, nil
}
