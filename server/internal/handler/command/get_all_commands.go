package command

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) GetAllCommands(c *gin.Context) {
	res, err := h.Service.GetAllCommands(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}
