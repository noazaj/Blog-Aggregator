package main

import (
	"net/http"

	"github.com/zajicekn/Blog-Aggregator/internal/database"
)

func (cfg *apiConfig) getUser(w http.ResponseWriter, r *http.Request, user database.User) {
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
