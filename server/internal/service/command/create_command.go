package command

import (
	"context"
	"pstgrprof/server/internal/entity"
	"strconv"
)

func (s *Service) CreateCommand(ctx context.Context, req *entity.CreateCommandReq) (*entity.CreateCommandRes, error) {
	cd := &entity.Command{
		Script:      req.Script,
		Description: req.Description,
	}

	r, err := s.Repository.CreateCommand(ctx, cd)
	if err != nil {
		return nil, err
	}

	res := &entity.CreateCommandRes{
		ID:          strconv.Itoa(int(r.ID)),
		Script:      r.Script,
		Description: req.Description,
	}

	return res, nil
}
