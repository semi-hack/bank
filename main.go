package main

import (
	"log"
	db "bank/database"
	//"bank/routes"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Printf("No .env file Found")
	}
}

func main() {
	db.ConnectDB()
	//routes.Initialize()
}