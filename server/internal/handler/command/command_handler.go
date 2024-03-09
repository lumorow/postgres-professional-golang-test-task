package command

import (
	"context"
	"pstgrprof/server/internal/entity"
)

//go:generate mockgen -destination=mocks/handler.go -package=mock -source=command_handler.go
//go:generate touch mocks/.coverignore

type Service interface {
	CreateCommand(c context.Context, req *entity.CreateCommandReq) (*entity.CreateCommandRes, error)
	GetCommandById(c context.Context, id string) (*entity.Command, error)
	GetAllCommands(c context.Context) (*[]entity.Command, error)
}

type Handler struct {
	Service
}

func NewHandler(s Service) *Handler {
	return &Handler{
		Service: s,
	}
}