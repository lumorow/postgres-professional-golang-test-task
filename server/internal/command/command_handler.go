package command

import (
	"github.com/gin-gonic/gin"
)

type Handler struct {
	Service
}

func NewHandler(s Service) *Handler {
	return &Handler{
		Service: s,
	}
}

func (h *Handler) CreateCommand(c *gin.Context) {

}

func (h *Handler) GetCommand(c *gin.Context) {

}

func (h *Handler) GetCommands(c *gin.Context) {

}
