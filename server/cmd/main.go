package main

import (
	"errors"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"log"
	"pstgrprof/server/db"
	"pstgrprof/server/internal/handler/command"
	command_cache "pstgrprof/server/internal/repository/cache"
	command_repo "pstgrprof/server/internal/repository/command"
	"pstgrprof/server/internal/router"
	command_service "pstgrprof/server/internal/service/command"

	"github.com/spf13/viper"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/github"
)

// @title Launch-command
// @version 1.0
// description API Server for Launch-command

// @host localhost:8000
// @BasePath /
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

	if err = dbConn.Migrate(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.Fatalf("could not run migration: %s", err.Error())
	}

	commandRep := command_repo.NewRepository(dbConn.GetDB())
	scriptsCache := command_cache.NewCache()
	execCmdCache := command_cache.NewCache()
	commandSvc := command_service.NewService(commandRep, scriptsCache, execCmdCache)
	defer func() { commandSvc.StopRunner() }()
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
