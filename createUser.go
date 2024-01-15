package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/zajicekn/Blog-Aggregator/internal/database"
)

// New user name to be created
type userRequest struct {
	Name string `json:"name"`
}

// New user response with name attached to new data
type userResponse struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
}

// createUser is utilized to handle the endpoint of v1/users
func (cfg *apiConfig) createUser(w http.ResponseWriter, r *http.Request) {
	var req userRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		respondWithError(w, http.StatusExpectationFailed, "error generating decoder")
	}

	// Generate a new UUID for the user
	newUUID := uuid.New()

	// Get the current time for created and updated
	currentTime := time.Now()

	// Prepare parameters for CreateUser function for the DB
	params := database.CreateUserParams{
		ID:        newUUID,
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
		Name:      req.Name,
	}

	// Call the CreateUser method using the params
	user, err := cfg.DB.CreateUser(r.Context(), params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "error creating user")
	}

	// Prepare the response
	resp := userResponse{
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Name:      user.Name,
	}

	// Respond with JSON
	respondWithJson(w, http.StatusOK, resp)
}
