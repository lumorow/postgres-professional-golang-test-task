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
	DeleteCommandById(ctx context.Context, id int64) error
	CreateCommandOutput(ctx context.Context, id int64, output string) error
}

type Cache interface {
	Set(key int64, value any) error
	Get(key int64) (any, error)
	GetAllKeys() ([]int64, error)
	Delete(key int64) error
}

type Service struct {
	Repository
	ScriptsCache Cache
	ExecCmdCache Cache
}

func NewService(repository Repository, scriptsCache, execCmdCache Cache) *Service {
	s := &Service{
		repository,
		scriptsCache,
		execCmdCache,
	}

	go s.Runner()
	return s
}
