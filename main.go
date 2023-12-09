package main

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/simplebankapp/api"
	db "github.com/simplebankapp/db/sqlc"
	"github.com/simplebankapp/db/util"
)

func main() {

	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config: ", err)
	}

	dbConnPool, err := pgxpool.New(context.Background(), config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to DB:", err)
	}
	defer dbConnPool.Close()

	store := db.NewStore(dbConnPool)
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server object: ", err)
	}
	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}