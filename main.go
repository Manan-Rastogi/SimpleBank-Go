package main

import (
	"database/sql"
	"log"

	"github.com/Manan-Rastogi/simplebank/api"
	db "github.com/Manan-Rastogi/simplebank/db/sqlc"
	_ "github.com/lib/pq"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://root:password@localhost:5432/simple_bank?sslmode=disable"
	serverAddress = "127.0.0.1:8080"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot establish connection to db : ", err.Error())
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddress)
	if err != nil{
		log.Fatal("cannot start server: ", err.Error())
	}
}