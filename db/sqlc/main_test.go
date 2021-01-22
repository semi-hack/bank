package db

import (
	"log"
	"database/sql"
	"testing"
	"os"
	_ "github.com/lib/pq"

)

const (
	dbDriver = "postgres"
	dbSource = "postgres://postgres:user123@localhost:5432/bank?sslmode=disable"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	var err error

	testDB, err = sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cant connect to db", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}