package main

import "net/http"

func handlerError(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	respondWithError(w, http.StatusInternalServerError, "Internal Server Error")
}
