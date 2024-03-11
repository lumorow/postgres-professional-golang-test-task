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

	//	Создание команды
	r.POST("/command", commandHandler.CreateCommand)

	//	Запуск команды
	r.GET("/command/:id", commandHandler.GetCommandById)

	//	Запуск нескольких команд
	r.GET("/commands", commandHandler.GetCommands)

	//	Получение списка всех команд
	r.GET("/all-commands", commandHandler.GetAllCommands)

	//	Удаление команды
	r.DELETE("/command/:id", commandHandler.DeleteCommandById)

	//	Остановка команды
	r.POST("/command/:id", commandHandler.StopCommandById)

	// Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}

func Start(r *gin.Engine, addr string) error {
	return r.Run(addr)
}
