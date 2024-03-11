package command

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetCommandById @Summary Start command
// @Description get and start command by id
// @Tags command
// @Produce json
// @Param        id   path      int  true  "Command ID"
// @Success 200 {object} string
// @Failure 500 {object} string
// @Router /command/{id} [get]
func (h *Handler) GetCommandById(c *gin.Context) {
	id := c.Param("id")
	res, err := h.Service.GetCommandById(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}
