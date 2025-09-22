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

type Store struct {
	DB      *sql.DB
	Queries *sqlc.Queries
}

func LoadConfig() Config {
	return Config{
		DB_URL: os.Getenv("DB_LOCAL_URL"),
	}
}

func ConnectDB() (*Store, error) {
	config := LoadConfig()

	db_url := config.DB_URL

	db, err := sql.Open("postgres", db_url)
	if err != nil {
		return nil, fmt.Errorf("could not open db: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("could not connect to db: %w", err)
	}

	log.Println("Successfully connected to the database")
	return &Store{
		DB:      db,
		Queries: sqlc.New(db),
	}, nil
}

func (s *Store) Disconnect() {
	if s.DB != nil {
		s.DB.Close()
		log.Println("Database connection closed")
	}
}
