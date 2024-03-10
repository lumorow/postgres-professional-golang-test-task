package command

import (
	"context"
	"pstgrprof/server/internal/entity"
)

//go:generate mockgen -destination=mocks/service.go -package=mock -source=command_service.go
//go:generate touch mocks/.coverignore

type Repository interface {
	CreateCommand(ctx context.Context, command *entity.Command) (*entity.Command, error)
	GetCommandById(ctx context.Context, id int64) (*entity.Command, error)
	GetAllCommands(ctx context.Context) (*[]entity.Command, error)
	GetCommands(ctx context.Context, id []int64) (*[]entity.Command, error)
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
