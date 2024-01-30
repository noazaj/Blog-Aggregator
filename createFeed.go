package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/zajicekn/Blog-Aggregator/internal/database"
)

// New feeds request with name and url body
type feedRequest struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

// Feeds response
type feedResponse struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Url       string    `json:"url"`
	UserID    uuid.UUID `json:"user_id"`
}

func (cfg *apiConfig) createFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	var req feedRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		respondWithError(w, http.StatusExpectationFailed, "error generating decoder")
	}

	// Create a new UUID for the feed
	newUUID := uuid.New()

	// Get the current time for the feed
	currentTime := time.Now()

	// Prepare params for CreateFeed function for DB
	params := database.CreateFeedParams{
		ID:        newUUID,
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
		Name:      req.Name,
		Url:       req.URL,
		UserID:    user.ID,
	}

	// Call the CreateFeed method using the params
	feed, err := cfg.DB.CreateFeed(r.Context(), params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "error creating feed")
	}

	// Prepare the response
	resp := feedResponse{
		ID:        feed.ID,
		CreatedAt: feed.CreatedAt,
		UpdatedAt: feed.UpdatedAt,
		Name:      feed.Name,
		Url:       feed.Url,
		UserID:    user.ID,
	}

	// Respond with JSON
	respondWithJson(w, http.StatusOK, resp)
}
