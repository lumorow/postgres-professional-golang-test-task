package command

import (
	"context"
	"pstgrprof/server/internal/entity"
)

//go:generate mockgen -destination=mocks/handler.go -package=mock -source=command_handler.go
//go:generate touch mocks/.coverignore

type Service interface {
	CreateCommand(ctx context.Context, req *entity.CreateCommandReq) (*entity.CreateCommandRes, error)
	GetCommandById(ctx context.Context, id string) (*entity.Command, error)
	GetAllCommands(ctx context.Context) (*[]entity.Command, error)
	GetCommands(ctx context.Context, ids []string) (*[]entity.Command, error)
	DeleteCommandById(ctx context.Context, id string) error
	StopCommandById(id string) error
}

type Handler struct {
	Service
}

func NewHandler(s Service) *Handler {
	return &Handler{
		Service: s,
	}
}
