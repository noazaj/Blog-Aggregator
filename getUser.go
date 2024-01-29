package main

import (
	"database/sql"
	"log"
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

func (cfg *apiConfig) getUser(w http.ResponseWriter, r *http.Request) {
	// Get the key from the helper function
	apiKey, err := cfg.getAuthHeader(w, r)
	if err != nil {
		log.Fatal(err)
	}

	// Use the key to create the response
	user, err := cfg.DB.GetUserApiKey(r.Context(), apiKey)
	if err != nil {
		if err == sql.ErrNoRows {
			respondWithError(w, http.StatusNotFound, "user not found")
		} else {
			respondWithError(w, http.StatusInternalServerError, "error retrieving user")
		}
		return
	}

	// Generate the response for the user
	resp := userResponse{
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Name:      user.Name,
		ApiKey:    user.ApiKey,
	}

	// Respond with JSON
	respondWithJson(w, http.StatusOK, resp)
}
