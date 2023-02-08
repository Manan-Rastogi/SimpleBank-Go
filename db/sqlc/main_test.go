package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/Manan-Rastogi/simplebank/util"
	_ "github.com/lib/pq"
)



var (
	testQueries *Queries
	testDB      *sql.DB
)

func TestMain(m *testing.M) { // Main testing point
	var err error
	config, err := util.LoadConfig("../..")
	if err != nil{
		log.Fatal("cannot load configs: ", err.Error())
	}
	testDB, err = sql.Open(config.DbDriver, config.DbSource)
	if err != nil {
		log.Fatal("cannot establish connection to db : ", err.Error())
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}
