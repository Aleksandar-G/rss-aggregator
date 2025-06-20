package main

import (
	"log"
	"net/http"
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
