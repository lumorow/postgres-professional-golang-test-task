package command

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) GetCommands(c *gin.Context) {
	ids := c.QueryArray("id")

	res, err := h.Service.GetCommands(c.Request.Context(), ids)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}
