package main

import (
	"errors"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"log"
	"os"
	"pstgrprof/server/db"
	"pstgrprof/server/internal/handler/command"
	command_cache "pstgrprof/server/internal/repository/cache"
	command_repo "pstgrprof/server/internal/repository/command"
	"pstgrprof/server/internal/router"
	command_service "pstgrprof/server/internal/service/command"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/github"
)

// @title Launch-command
// @version 1.0
// description API Server for Launch-command

// @host localhost:8000
// @BasePath /
func main() {
	log.Printf("Init config...")
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	log.Printf("Starting DB...")
	dbConn, err := db.NewDatabase(db.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Username: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
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

	log.Printf("Starting Server...")

	if err = router.Start(r, fmt.Sprintf("%s:%s", os.Getenv("SERVER_HOST"), os.Getenv("SERVER_PORT"))); err != nil {
		log.Fatalf("could not start server: %s", err.Error())
	}
}

func initConfig() error {
	if err := godotenv.Load("./server/configs/.env"); err != nil {
		return fmt.Errorf("error loading .env file: %w", err)
	}

	viper.AutomaticEnv()

	return nil
}
