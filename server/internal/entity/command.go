package entity

type Command struct {
	ID          int64  `json:"id" db:"id"`
	Script      string `json:"script" db:"script"`
	Description string `json:"description" db:"description"`
}

type CreateCommandReq struct {
	Script      string `json:"script" db:"script"`
	Description string `json:"description" db:"description"`
}

type CreateCommandRes struct {
	ID          string `json:"id" db:"id"`
	Script      string `json:"script" db:"script"`
	Description string `json:"description" db:"description"`
}
