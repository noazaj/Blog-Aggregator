package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/zajicekn/Blog-Aggregator/internal/database"
)

// New feed_follow request
type feedFollowRequest struct {
	FeedID uuid.UUID `json:"feed_id"`
}

// feed_follow response
type feedFollowResponse struct {
	ID        uuid.UUID `json:"id"`
	FeedID    uuid.UUID `json"feed_id"`
	UserID    uuid.UUID `josn"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (cfg *apiConfig) createFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	var req feedFollowRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		respondWithError(w, http.StatusExpectationFailed, "error generating decoder")
	}

	newUUID := uuid.New()

	currentTime := time.Now()

	// Prepare params
	params := database.CreateFeedFollowsParams{
		ID:        newUUID,
		FeedID:    req.FeedID,
		UserID:    user.ID,
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
	}

	// Call the CreateFeedFollow method
	feedFollow, err := cfg.DB.CreateFeedFollows(r.Context(), params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "error creating feed follow")
	}

	// Prepare response body
	resp := feedFollowResponse{
		ID:        feedFollow.ID,
		FeedID:    feedFollow.FeedID,
		UserID:    user.ID,
		CreatedAt: feedFollow.CreatedAt,
		UpdatedAt: feedFollow.UpdatedAt,
	}

	respondWithJson(w, http.StatusOK, resp)
}
