package router

import (
	"github.com/gin-gonic/gin"
	"pstgrprof/server/internal/command"
)

func InitRouter(r *gin.Engine, commandHandler *command.Handler) {
	r = gin.Default()

	r.POST("/command", commandHandler.CreateCommand)
	r.GET("/command/:id", commandHandler.GetCommand)
	r.GET("/commands", commandHandler.GetAllCommands)

}

func Start(r *gin.Engine, addr string) error {
	return r.Run(addr)
}
