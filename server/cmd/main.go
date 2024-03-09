package main

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"pstgrprof/server/db"
	"pstgrprof/server/internal/handler/command"
	command_repo "pstgrprof/server/internal/repository/command"
	command_service "pstgrprof/server/internal/service/command"
	"pstgrprof/server/router"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	dbConn, err := db.NewDatabase(db.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: viper.GetString("db.password"),
	})
	if err != nil {
		log.Fatalf("could not initialize database connection: %s", err.Error())
	}
	defer dbConn.Close()

	commandRep := command_repo.NewRepository(dbConn.GetDB())
	commandSvc := command_service.NewService(commandRep)
	CommandHandler := command.NewHandler(commandSvc)

	r := router.InitRouter(CommandHandler)
	if err = router.Start(r, fmt.Sprintf("%s:%s", viper.GetString("server.host"), viper.GetString("server.port"))); err != nil {
		log.Fatalf("could not start server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("server/configs")
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	return viper.ReadInConfig()
}
