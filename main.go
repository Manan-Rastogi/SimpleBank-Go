package main

import (
	"database/sql"
	"log"

	"github.com/Manan-Rastogi/simplebank/api"
	db "github.com/Manan-Rastogi/simplebank/db/sqlc"
	"github.com/Manan-Rastogi/simplebank/util"
	_ "github.com/lib/pq"
)



func main() {
	config, err := util.LoadConfig(".")
	if err != nil{
		log.Fatal("cannot load config: ", err.Error())
	}

	conn, err := sql.Open(config.DbDriver, config.DbSource)
	if err != nil {
		log.Fatal("cannot establish connection to db : ", err.Error())
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil{
		log.Fatal("cannot start server: ", err.Error())
	}
}