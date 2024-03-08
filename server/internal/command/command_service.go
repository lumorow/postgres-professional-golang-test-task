package command

import (
	"context"
	"errors"
	"strconv"
	"time"
)

type service struct {
	Repository
	timeout time.Duration
}

func NewService(repository Repository) Service {
	return &service{
		repository,
		time.Duration(2) * time.Second,
	}
}

func (s *service) CreateCommand(c context.Context, req *CreateCommandReq) (*CreateCommandRes, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	if len(req.CommandBody) == 0 {
		return nil, errors.New("script must not be empty")
	}

	if len(req.Description) == 0 {
		return nil, errors.New("description for script must not be empty")
	}

	cd := &Command{
		CommandBody: req.CommandBody,
		Description: req.Description,
	}

	r, err := s.Repository.CreateCommand(ctx, cd)
	if err != nil {
		return nil, err
	}

	res := &CreateCommandRes{
		ID:          strconv.Itoa(int(r.ID)),
		CommandBody: r.CommandBody,
		Description: req.Description,
	}

	return res, nil
}

func (s *service) GetCommandById(c context.Context, id int64) (*Command, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	r, err := s.Repository.GetCommandById(ctx, id)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (s *service) GetCommands(c context.Context) (*[]Command, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	r, err := s.Repository.GetAllCommands(ctx)
	if err != nil {
		return nil, err
	}

	return r, nil
}
