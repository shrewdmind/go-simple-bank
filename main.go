package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/shrewdmind/simplebank/api"
	db "github.com/shrewdmind/simplebank/db/sqlc"
	"github.com/shrewdmind/simplebank/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	conn, err := sql.Open(config.DbDriver, config.DbSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", config.DbDriver)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}