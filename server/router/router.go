package router

import (
	"github.com/gin-gonic/gin"
	"pstgrprof/server/internal/handler/command"
)

func InitRouter(commandHandler *command.Handler) *gin.Engine {
	r := gin.New()

	r.POST("/command", commandHandler.CreateCommand)
	r.GET("/command/:id", commandHandler.GetCommand)
	r.GET("/commands", commandHandler.GetAllCommands)
	r.DELETE("/command/:id", commandHandler.GetCommand)
	return r
}

func Start(r *gin.Engine, addr string) error {
	return r.Run(addr)
}
