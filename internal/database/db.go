package internal

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	sqlc "nba-api/sql/database"

	_ "github.com/lib/pq"
)

type Config struct {
	DB_URL string
}

func LoadConfig() Config {
	return Config{
		DB_URL: os.Getenv("DB_LOCAL_URL"),
	}
}

// Global variables to hold the database and queries
var (
	DB      *sql.DB
	Queries *sqlc.Queries
)

func ConnectDB() error {
	config := LoadConfig()

	db_url := config.DB_URL

	db, err := sql.Open("postgres", db_url)
	if err != nil {
		return fmt.Errorf("could not open db: %w", err)
	}

	if err := db.Ping(); err != nil {
		return fmt.Errorf("could not connect to db: %w", err)
	}

	DB = db
	Queries = sqlc.New(DB) // Initialize sqlc Queries w/ the DB connection

	log.Println("Successfully connected to the database")
	return nil
}

func DisconnectDB() {
	if DB != nil {
		DB.Close()
		log.Println("Database connection closed")
	}
}
