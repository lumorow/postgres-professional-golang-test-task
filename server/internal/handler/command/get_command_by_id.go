package command

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) GetCommand(c *gin.Context) {
	id := c.Param("id")
	res, err := h.Service.GetCommandById(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}
