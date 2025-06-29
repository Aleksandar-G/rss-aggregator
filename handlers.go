package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/Aleksandar-G/rss-aggregator/internal/database"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
)

// Readiness response
func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	log.Println("Request on GET /v1/healthz")
	// Create a statusResponse struct
	type readinessResponse struct {
		Status string `json:"status"`
	}

	respondWithJSON(w, http.StatusOK, readinessResponse{Status: "ok"})
}

// Test error handler
func handlerErr(w http.ResponseWriter, r *http.Request) {
	log.Println("Request on GET /v1/err")
	respondWithError(w, http.StatusInternalServerError, "Internal Server Error")
}

// User handler
func (apiCfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	// Create request body struct
	type parameters struct {
		Name string `json:"name"`
	}

	// Decode the request body and convert to `parameters` struct
	jsonDecoder := json.NewDecoder(r.Body)
	params := parameters{}

	err := jsonDecoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Not a valid body")
		return
	}

	// Create the user
	dbUser, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})
	if err != nil {
		respondWithError(w, 500, err.Error())
		return
	}

	// Convert database user object to models user
	user := databaseUserToUser(dbUser)

	// Return a successful response
	respondWithJSON(w, http.StatusCreated, user)
}

func (apiCfg *apiConfig) handlerDeleteUser(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "id")
	if userId == "" {
		respondWithError(w, http.StatusBadRequest, "Id is not passed")
		return
	}

	err := apiCfg.DB.DeleteUser(r.Context(), userId)
	if err != nil {
		respondWithError(w, 500, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, struct {
		Status string `json:"status"`
	}{
		Status: "resource deleted",
	})
}

func (apiCfg *apiConfig) handlerGetUser(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "id")
	if userId == "" {
		respondWithError(w, http.StatusBadRequest, "Id is not passed")
		return
	}

	// Fetch the user from the database
	dbUser, err := apiCfg.DB.GetUserById(r.Context(), userId)
	if err != nil {
		respondWithError(w, 500, err.Error())
		return
	}
	// Convert database user to models user
	user := databaseUserToUser(dbUser)

	// Return response with user in body
	respondWithJSON(w, http.StatusOK, user)
}

func (apiCfg *apiConfig) handlerListUsers(w http.ResponseWriter, r *http.Request) {

	// Fetch the user from the database
	dbUsers, err := apiCfg.DB.ListUsers(r.Context())
	if err != nil {
		respondWithError(w, 500, err.Error())
		return
	}
	// Convert database user to models user
	var users []User

	for _, dbUser := range dbUsers {
		users = append(users, databaseUserToUser(dbUser))
	}

	// Return response with user in body
	respondWithJSON(w, http.StatusOK, users)
}
