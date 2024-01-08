package main

import "net/http"

func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	type payload struct {
		Status string `json:"status"`
	}
	respondWithJson(w, http.StatusOK, payload{"ok"})
}
