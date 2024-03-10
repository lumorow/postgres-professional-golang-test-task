package command

import (
	"context"
	"strconv"
)

func (s *Service) DeleteCommandById(ctx context.Context, id string) error {
	i, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}

	err = s.Repository.DeleteCommandById(ctx, i)
	if err != nil {
		return err
	}

	return nil
}
