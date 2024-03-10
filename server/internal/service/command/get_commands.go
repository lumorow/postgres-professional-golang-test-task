package command

import (
	"context"
	"pstgrprof/server/internal/entity"
	"strconv"
)

func (s *Service) GetCommands(ctx context.Context, ids []string) (*[]entity.Command, error) {
	iids := make([]int64, 0, len(ids))
	for _, id := range ids {
		i, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			return nil, err
		}
		iids = append(iids, i)
	}

	r, err := s.Repository.GetCommands(ctx, iids)

	if err != nil {
		return nil, err
	}

	return r, nil
}
