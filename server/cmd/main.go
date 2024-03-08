package main

import (
	"github.com/spf13/viper"
	"log"
	"pstgrprof/server/db"
	"pstgrprof/server/internal/command"
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

	commandRep := command.NewRepository(dbConn.GetDB())
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
