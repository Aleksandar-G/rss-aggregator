package handlers

import (
	"log"
	"net/http"

	"github.com/Aleksandar-G/rss-aggregator/pkg"
)

// Readiness response
func HandlerReadiness(w http.ResponseWriter, r *http.Request) {
	log.Println("Request on GET /v1/healthz")
	// Create a statusResponse struct
	type readinessResponse struct {
		Status string `json:"status"`
	}

	pkg.RespondWithJSON(w, http.StatusOK, readinessResponse{Status: "ok"})
}

// Test error handler
func HandlerErr(w http.ResponseWriter, r *http.Request) {
	log.Println("Request on GET /v1/err")
	pkg.RespondWithError(w, http.StatusInternalServerError, "Internal Server Error")
}
