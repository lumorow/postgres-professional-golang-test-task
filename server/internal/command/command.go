package command

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
	ID          int64  `json:"id" db:"id"`
	CommandBody string `json:"commandbody" db:"commandbody"`
	Description string `json:"description" db:"description"`
}

type Repository interface {
}

type Service interface {
}
