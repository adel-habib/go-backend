package main

import (
	"database/sql"
	"log"

	"github.com/adel-habib/go-backend/api"
	db "github.com/adel-habib/go-backend/db/sqlc"
	"github.com/adel-habib/go-backend/util"
	_ "github.com/lib/pq"
)

func main() {

	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load configs")
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("Could not start server", err)
	}
}
