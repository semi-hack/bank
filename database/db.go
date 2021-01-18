package database

import (
  "database/sql"
  "fmt"
  "log"
  // "os"

  _ "github.com/lib/pq"
)

var ( 
  db *sql.DB
)

var err error


// ConnectDB ...
func init() {
  // dbname := os.Getenv("DBNAME")
  // dbsource := os.Getenv("DBSOURCE")
  db, err = sql.Open("postgres", "postgres://postgres:user123@localhost:5432/bank?sslmode=disable")
  // db, err = sql.Open(dbname, dbsource)

  if err != nil {
    fmt.Println(err.Error())
  }

  if err = db.Ping(); err != nil {
    fmt.Println(err.Error())
  }

  log.Println("connected to database")
}

// ConnectDB ...
func ConnectDB() *sql.DB {
  return db
}