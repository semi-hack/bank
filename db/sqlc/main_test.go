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

func TestMain(m *testing.M) {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cant connect to db", err)
	}

	testQueries = New(conn)

	os.Exit(m.Run())
}