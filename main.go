package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {
	// Load enviornment variables
	err := godotenv.Load(".env")
	if err != nil {
		log.Printf("Error: %v\nCouldn't load enviornment variables", err)
	}

	// Set port variable
	port := os.Getenv("PORT")

	// Create a router
	router := chi.NewRouter()

	// Implement the CORS middleware to the router
	router.Use(middlewareCors)

	// Create a sub router and mount it to the main router
	v1 := chi.NewRouter()
	router.Mount("/v1", v1)

	// Setup a file server for static files
	fs := http.FileServer(http.Dir("public"))
	router.Handle("/*", fs)

	// Setup handlers
	v1.Get("/readiness", handlerReadiness)
	v1.Get("/err", handlerError)

	// Create a server with the designated port and router
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	// Start the server on its designated port
	fmt.Printf("Starting server on port :%s\n", port)
	log.Fatal(srv.ListenAndServe())

}
