package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/Aleksandar-G/rss-aggregator/internal/database"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"

	_ "modernc.org/sqlite"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	// Load environment variables
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT is not set")
	}

	dbUrl := os.Getenv("DB_URL")
	if dbUrl == "" {
		log.Fatal("DB_URL is not set")
	}

	// Open a connection to the database
	conn, err := sql.Open("sqlite", dbUrl)
	if err != nil {
		log.Fatal("Cannot open a connection to the database")
	}

	apiCfg := apiConfig{
		DB: database.New(conn),
	}

	// Main base router
	mainRouter := chi.NewRouter()

	// CORS settings
	mainRouter.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	// V1 router
	v1Router := chi.NewRouter()

	// Mount the V1 Router to the mainRouter on the `/v1` path
	mainRouter.Mount("/v1", v1Router)

	// Endpoints for the V1 router
	v1Router.Get("/healthz", handlerReadiness)
	v1Router.Get("/err", handlerErr)
	v1Router.Post("/users", apiCfg.handlerCreateUser)

	// Create a server
	server := &http.Server{
		Addr:    ":" + port,
		Handler: mainRouter,
	}

	// Server the server
	log.Printf("Serving server on port: %s\n", port)
	err = server.ListenAndServe()
	log.Fatalf("The server has stopped with an error: %v", err)
}
