package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/zajicekn/Blog-Aggregator/internal/database"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	// Load enviornment variables
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error: Couldn't load enviornment variables")
	}

	// Set port and database connection variable
	port := os.Getenv("PORT")
	conn := os.Getenv("CONN")

	// Load the database URL and open a connection to the database
	// and store it within the conig struct
	db, err := sql.Open("postgres", conn)
	if err != nil {
		log.Fatal("Error: Couldn't connect to database")
	}
	dbQueries := database.New(db)

	// Store the queries in the config struct
	config := apiConfig{
		DB: dbQueries,
	}

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
	v1.Post("/users", config.createUser)

	// Create a server with the designated port and router
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	// Start the server on its designated port
	fmt.Printf("Starting server on port :%s\n", port)
	log.Fatal(srv.ListenAndServe())

}
