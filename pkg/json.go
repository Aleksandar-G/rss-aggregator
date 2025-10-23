package pkg

import (
	"encoding/json"
	"log"
	"net/http"
)

// Get `ResponseWriter`, `responseCode` and `payload` and marshal an json response to be send to the client
func RespondWithJSON(w http.ResponseWriter, responseCode int, payload interface{}) {
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
func RespondWithError(w http.ResponseWriter, responseCode int, errMsg string) {

	// Create en errorResponse struct so that the it implements the error interface
	type errorResponse struct {
		Error string `json:"error"`
	}

	// Checks if it is a server error
	if responseCode > 499 {
		log.Printf("Internal error has occurred error message: %v\n", errMsg)
		RespondWithJSON(w, 500, errorResponse{
			Error: "Internal server error",
		})
		return
	}

	// Create the JSON response
	RespondWithJSON(w, responseCode, errorResponse{
		Error: errMsg,
	})
}

func DecodeRequestBody[V any](r *http.Request, parameters V) (*V, error) {
	// Decode the request body and convert to `parameters` struct
	jsonDecoder := json.NewDecoder(r.Body)

	err := jsonDecoder.Decode(&parameters)
	if err != nil {
		return nil, err
	}
	return &parameters, nil
}
