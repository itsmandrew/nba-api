package main

import (
	"log"

	db "nba-api/internal/database"
	"nba-api/internal/server"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {

	store, err := db.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	defer store.Disconnect()

	srv := server.New(store)
	if err := srv.Start(); err != nil {
		log.Fatal(err)
	}
}
