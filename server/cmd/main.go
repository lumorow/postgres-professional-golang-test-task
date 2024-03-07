package main

import (
	"log"
	"pstgrprof/server/db"
	"pstgrprof/server/internal/command"
)

func main() {
	dbConn, err := db.NewDatabase()
	defer dbConn.Close()
	if err != nil {
		log.Fatalf("could not initialize database connection: %s", err)
	}

	commandRep := command.NewRepository(dbConn.GetDB())
}
