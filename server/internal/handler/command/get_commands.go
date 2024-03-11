package command

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetCommands @Summary Start some commands
// @Description get and start some command by ids
// @Tags commands
// @Produce json
// @Param    ids   query	[]int  true  "Command IDs"
// @Success 200 {object} string
// @Failure 500 {object} string
// @Router /commands [get]
func (h *Handler) GetCommands(c *gin.Context) {
	ids := c.QueryArray("id")

	res, err := h.Service.GetCommands(c.Request.Context(), ids)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}
