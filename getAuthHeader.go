package main

import (
	"net/http"
	"strings"
)

func (cfg *apiConfig) getAuthHeader(w http.ResponseWriter, r *http.Request) (string, error) {
	// create the authHeader
	authHeader := r.Header.Get("Authorization")

	// Split the header into two parts
	headerParts := strings.Split(authHeader, " ")
	if len(headerParts) != 2 || headerParts[0] != "ApiKey" {
		// Handle error: either the header is not in the correct format
		// or the 'ApiKey' keyword was missing
		respondWithError(w, http.StatusBadRequest, "Invalid ApiKey or format")
		return "", nil
	}

	// Get the apiKey and return it
	return headerParts[1], nil
}
