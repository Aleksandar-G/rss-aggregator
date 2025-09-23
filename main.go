package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Aleksandar-G/rss-aggregator/internal/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"

	_ "modernc.org/sqlite"
)

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

	userHandler := handlers.NewUserHandler()
	feedHandler := handlers.NewFeedHandler()
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
	v1Router.Get("/healthz", handlers.HandlerReadiness)
	v1Router.Get("/err", handlers.HandlerErr)

	// Mount the V1 Router to the mainRouter on the `/v1` path
	mainRouter.Mount("/v1", v1Router)

	// Users router
	userRouter := chi.NewRouter()
	userRouter.Get("/{id}", userHandler.HandlerGetUser)
	userRouter.Get("/", userHandler.HandlerListUsers)
	userRouter.Post("/", userHandler.HandlerCreateUser)
	userRouter.Delete("/{id}", userHandler.HandlerDeleteUser)

	// Feeds router
	feedRouter := chi.NewRouter()
	feedRouter.Get("/{id}", feedHandler.HandlerGetFeed)
	feedRouter.Get("/", feedHandler.HandlerListFeeds)
	feedRouter.Post("/", feedHandler.HandlerCreateFeed)
	feedRouter.Delete("/{id}", feedHandler.HandlerDeleteFeed)

	// Mount the user Router to the v1Router
	v1Router.Mount("/users", userRouter)
	v1Router.Mount("/feeds", feedRouter)

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
