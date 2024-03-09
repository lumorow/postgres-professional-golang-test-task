package command

import (
	"context"
	"errors"
	"fmt"
	"pstgrprof/server/internal/entity"
	"strconv"
	"time"
)

//go:generate mockgen -destination=mocks/service.go -package=mock -source=command_service.go
//go:generate touch mocks/.coverignore

type Repository interface {
	CreateCommand(ctx context.Context, command *entity.Command) (*entity.Command, error)
	GetCommandById(ctx context.Context, id int64) (*entity.Command, error)
	GetAllCommands(ctx context.Context) (*[]entity.Command, error)
}

type Service struct {
	Repository
	timeout time.Duration
}

func NewService(repository Repository) *Service {
	return &Service{
		repository,
		time.Duration(2) * time.Second,
	}
}

func (s *Service) CreateCommand(ctx context.Context, req *entity.CreateCommandReq) (*entity.CreateCommandRes, error) {

	if len(req.Script) == 0 {
		return nil, errors.New("script must not be empty")
	}

	if len(req.Description) == 0 {
		return nil, errors.New("description for script must not be empty")
	}

	cd := &entity.Command{
		Script:      req.Script,
		Description: req.Description,
	}

	r, err := s.Repository.CreateCommand(ctx, cd)
	if err != nil {
		return nil, err
	}

	res := &entity.CreateCommandRes{
		//ID:          strconv.Itoa(int(r.ID)),
		ID:          "1",
		Script:      r.Script,
		Description: req.Description,
	}

	fmt.Println(res)

	return res, nil
}

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

func (s *Service) GetAllCommands(ctx context.Context) (*[]entity.Command, error) {

	r, err := s.Repository.GetAllCommands(ctx)
	if err != nil {
		return nil, err
	}

	return r, nil
}
