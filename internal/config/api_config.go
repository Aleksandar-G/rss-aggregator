package config

import (
	"database/sql"
	"log"
	"os"

	"github.com/Aleksandar-G/rss-aggregator/internal/database"
)

type APIConfig struct {
	DB *database.Queries
}

// Return a new instance of the APIConfig
func NewAPIConfig() *APIConfig {

	dbUrl := os.Getenv("DB_URL")
	if dbUrl == "" {
		log.Fatal("DB_URL is not set")
	}

	// Open a connection to the database
	conn, err := sql.Open("sqlite", dbUrl)
	if err != nil {
		log.Fatal("Cannot open a connection to the database")
	}

	return &APIConfig{DB: database.New(conn)}
}
