package main

import (
	"database/sql"
	"net/http"

	"github.com/zajicekn/Blog-Aggregator/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func middlewareCors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func (cfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get the ApiKey
		apiKey, err := cfg.getAuthHeader(w, r)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, "Invalid ApiKey")
			return
		}

		// Validate the ApiKey and retrieve the user
		user, err := cfg.DB.GetUserApiKey(r.Context(), apiKey)
		if err != nil {
			if err == sql.ErrNoRows {
				respondWithError(w, http.StatusNotFound, "user not found")
			} else {
				respondWithError(w, http.StatusInternalServerError, "error retrieving user")
			}
			return
		}

		// Call the next handler with the user's information
		handler(w, r, user)
	}
}
