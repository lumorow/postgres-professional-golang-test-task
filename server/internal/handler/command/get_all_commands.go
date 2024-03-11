package command

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetAllCommands @Summary Get all commands
// @Description get all commands from database
// @Tags all-commands
// @Produce json
// @Success 200 {object} string
// @Failure 500 {object} string
// @Router /all-commands [get]
func (h *Handler) GetAllCommands(c *gin.Context) {
	res, err := h.Service.GetAllCommands(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}
