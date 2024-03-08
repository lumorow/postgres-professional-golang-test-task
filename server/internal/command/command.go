package command

import "context"

type Commands struct {
	cmds *[]Command
}

type Command struct {
	ID          int64  `json:"id" db:"id"`
	CommandBody string `json:"commandbody" db:"commandbody"`
	Description string `json:"description" db:"description"`
}

type CreateCommandReq struct {
	CommandBody string `json:"commandbody" db:"commandbody"`
	Description string `json:"description" db:"description"`
}

type CreateCommandRes struct {
	ID          string `json:"id" db:"id"`
	CommandBody string `json:"commandbody" db:"commandbody"`
	Description string `json:"description" db:"description"`
}

type Repository interface {
	CreateCommand(ctx context.Context, command *Command) (*Command, error)
	GetCommandById(ctx context.Context, id int64) (*Command, error)
	GetCommands(ctx context.Context) (*[]Command, error)
}

type Service interface {
	CreateCommand(c context.Context, req *CreateCommandReq) (*CreateCommandRes, error)
	GetCommandById(c context.Context, id int64) (*Command, error)
	GetCommands(c context.Context) (*[]Command, error)
}
