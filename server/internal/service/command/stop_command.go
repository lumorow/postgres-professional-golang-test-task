package command

import (
	"context"
	"strconv"
)

func (s *Service) StopCommandById(ctx context.Context, id string) error {
	i, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}

	err = s.Cache.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
