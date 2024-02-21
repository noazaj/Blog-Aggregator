package main

import "net/http"

func (cfg *apiConfig) getFeeds(w http.ResponseWriter, r *http.Request) {
	// Call the GetFeeds method query to form response
	feed, err := cfg.DB.GetFeeds(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "error getting feeds")
		return
	}

	// Generate the response for the feeds
	resp := feedResponse{
		ID:        feed.ID,
		CreatedAt: feed.CreatedAt,
		UpdatedAt: feed.UpdatedAt,
		Name:      feed.Name,
		Url:       feed.Url,
		UserID:    feed.UserID,
	}

	// Respond with JSON
	respondWithJson(w, http.StatusOK, resp)
}
