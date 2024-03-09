package command

import (
	"context"
	"pstgrprof/server/internal/entity"
	"strconv"
)

//go:generate mockgen -destination=mocks/service.go -package=mock -source=command_service.go
//go:generate touch mocks/.coverignore

type Repository interface {
	CreateCommand(ctx context.Context, command *entity.Command) (*entity.Command, error)
	GetCommandById(ctx context.Context, id int64) (*entity.Command, error)
	GetAllCommands(ctx context.Context) (*[]entity.Command, error)
}

type Cache interface {
	Set(key, value string) error
	Delete(key string) error
	GetAll() ([]int, error)
	CheckKey(key int) error
}

type Service struct {
	Repository
	Cache
}

func NewService(repository Repository, cache Cache) *Service {
	s := &Service{
		repository,
		cache,
	}

	go s.Runner()
	return s
}

func (s *Service) CreateCommand(ctx context.Context, req *entity.CreateCommandReq) (*entity.CreateCommandRes, error) {
	cd := &entity.Command{
		Script:      req.Script,
		Description: req.Description,
	}

	r, err := s.Repository.CreateCommand(ctx, cd)
	if err != nil {
		return nil, err
	}

	// Добавить к кеш значение
	// map[id] = entity.Command{}

	res := &entity.CreateCommandRes{
		ID:          strconv.Itoa(int(r.ID)),
		Script:      r.Script,
		Description: req.Description,
	}

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
