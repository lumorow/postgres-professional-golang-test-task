package command

type Command struct {
	ID          int64  `json:"id" db:"id"`
	CommandBody string `json:"commandbody" db:"commandbody"`
}

type CreateCommandReq struct {
	CommandBody string `json:"commandbody" db:"commandbody"`
}

type CreateCommandRes struct {
	ID          int64  `json:"id" db:"id"`
	CommandBody string `json:"commandbody" db:"commandbody"`
}

type Repository interface {
}

type Service interface {
}
