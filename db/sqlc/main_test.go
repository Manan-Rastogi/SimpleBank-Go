package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:password@localhost:5432/simple_bank?sslmode=disable"
)

var (
	testQueries *Queries
	testDB *sql.DB
)

func TestMain(m *testing.M) { // Main testing point
	var err error
	testDB, err = sql.Open(dbDriver, dbSource)
	if err != nil{
		log.Fatal("cannot establish connection to db : ", err.Error())
	}

	testQueries = New(testDB) 

	os.Exit(m.Run())
}