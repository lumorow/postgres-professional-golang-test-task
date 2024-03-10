package router

import (
	"github.com/gin-gonic/gin"
	"pstgrprof/server/internal/handler/command"
)

func InitRouter(commandHandler *command.Handler) *gin.Engine {
	r := gin.New()

	r.POST("/command", commandHandler.CreateCommand)
	r.GET("/command/:id", commandHandler.GetCommandById)
	r.GET("/commands", commandHandler.GetCommands)
	r.GET("/all-commands", commandHandler.GetAllCommands)
	r.DELETE("/command/:id", commandHandler.DeleteCommandById)
	r.POST("/command/:id", commandHandler.StopCommandById)
	return r
}

func Start(r *gin.Engine, addr string) error {
	return r.Run(addr)
}
