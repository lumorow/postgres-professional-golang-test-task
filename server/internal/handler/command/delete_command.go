package command

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// DeleteCommandById @Summary Delete command and command output from DB
// @Description get command id in path
// @Tags command
// @Produce json
// @Param id path int true "Command ID"
// @Success 200 {object} string
// @Failure 500 {object} string
// @Router /command/{id} [delete]
func (h *Handler) DeleteCommandById(c *gin.Context) {
	id := c.Param("id")
	err := h.Service.DeleteCommandById(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "command deleted"})
}
