package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")

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

	// Endpoints for the V1 router
	v1Router.Get("/healthz", handlerReadiness)
	v1Router.Get("/err", handlerErr)

	// Mount the V1 Router to the mainRouter on the `/v1` path
	mainRouter.Mount("/v1", v1Router)

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
