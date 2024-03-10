package command

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) DeleteCommandById(c *gin.Context) {
	id := c.Param("id")
	err := h.Service.DeleteCommandById(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "command deleted"})
}
