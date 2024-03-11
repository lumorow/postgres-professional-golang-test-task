package router

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	_ "pstgrprof/server/docs"
	"pstgrprof/server/internal/handler/command"
)

func InitRouter(commandHandler *command.Handler) *gin.Engine {
	r := gin.New()

	//	Create command
	r.POST("/command", commandHandler.CreateCommand)

	//	Start command
	r.GET("/command/:id", commandHandler.GetCommandById)

	//	Start some commans
	r.GET("/commands", commandHandler.GetCommands)

	//	Get list command (without start)
	r.GET("/all-commands", commandHandler.GetAllCommands)

	//	Delete command
	r.DELETE("/command/:id", commandHandler.DeleteCommandById)

	//	Stop command
	r.POST("/command/:id", commandHandler.StopCommandById)

	// Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}

func Start(r *gin.Engine, addr string) error {
	return r.Run(addr)
}
