package router

import (
	"github.com/gin-gonic/gin"
	"pstgrprof/server/internal/command"
)

func InitRouter(r *gin.Engine, commandHandler *command.Handler) {
	r = gin.Default()

	r.POST("/signup", commandHandler.CreateCommand)
	r.POST("/login", commandHandler.GetCommand)
	r.GET("/logout", commandHandler.GetCommands)

}

func Start(r *gin.Engine, addr string) error {
	return r.Run(addr)
}
