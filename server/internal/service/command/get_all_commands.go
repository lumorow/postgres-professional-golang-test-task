package command

import (
	"context"
	"pstgrprof/server/internal/entity"
)

func (s *Service) GetAllCommands(ctx context.Context) (*[]entity.Command, error) {

	r, err := s.Repository.GetAllCommands(ctx)
	if err != nil {
		return nil, err
	}

	return r, nil
}
