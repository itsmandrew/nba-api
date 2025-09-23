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
	DBUrl string
}

type Store struct {
	DB      *sql.DB // Connection pool
	Queries *sqlc.Queries
}

func LoadConfig() Config {
	return Config{
		DBUrl: os.Getenv("DB_LOCAL_URL"),
	}
}

func ConnectDB() (*Store, error) {
	config := LoadConfig()

	dbURL := config.DBUrl

	db, err := sql.Open("postgres", dbURL)
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
