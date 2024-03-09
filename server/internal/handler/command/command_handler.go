package command

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
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
	scripts map[int]struct{}
	Service
}

func NewHandler(s Service) *Handler {
	return &Handler{
		scripts: make(map[int]struct{}, 50),
		Service: s,
	}
}

func (h *Handler) CreateCommand(c *gin.Context) {
	var cd entity.CreateCommandReq
	if err := c.ShouldBindJSON(&cd); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.Service.CreateCommand(c.Request.Context(), &cd)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h *Handler) GetCommand(c *gin.Context) {
	id := c.Param("id")
	res, err := h.Service.GetCommandById(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h *Handler) GetAllCommands(c *gin.Context) {
	res, err := h.Service.GetAllCommands(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}
