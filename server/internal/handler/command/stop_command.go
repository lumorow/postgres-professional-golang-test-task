package command

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// StopCommandById @Summary Stop command
// @Description stop execution command by id
// @Tags command
// @Produce json
// @Param    id   path	int  true  "Command ID"
// @Success 200 {object} string
// @Failure 500 {object} string
// @Router /command/{id} [post]
func (h *Handler) StopCommandById(c *gin.Context) {
	id := c.Param("id")

	err := h.Service.StopCommandById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "command stopped successfully"})
}
