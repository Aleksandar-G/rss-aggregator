package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// Get `ResponseWriter`, `responseCode` and `payload` and marshal an json response to be send to the client
func respondWithJSON(w http.ResponseWriter, responseCode int, payload interface{}) {
	// Set response header to json
	w.Header().Set("Content-Type", "application/json")

	// Try to marshal payload to json
	jsonPayload, err := json.Marshal(payload)

	// Handle error in marshaling
	if err != nil {
		log.Printf("Failed to marshal the payload: %v with error message of %v\n", payload, err)
		w.WriteHeader(500)
		return
	}

	// Create the response
	w.WriteHeader(responseCode)
	w.Write(jsonPayload)
}

// Get `ResponseWriter`, `responseCode` and `errMsg` and create an error response to be send to the client
func respondWithError(w http.ResponseWriter, responseCode int, errMsg string) {

	// Create en errorResponse struct so that the it implements the error interface
	type errorResponse struct {
		Error string `json:"error"`
	}

	// Checks if it is a server error
	if responseCode > 499 {
		log.Printf("Responding with error code:%d and error message %v\n", responseCode, errMsg)
		return
	}

	// Create the Jsonresponse
	respondWithJSON(w, responseCode, errorResponse{
		Error: errMsg,
	})
}
