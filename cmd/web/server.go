package main

import (
	"fmt"
	"joshuamURD/wardens-court-summariser/config"
	"joshuamURD/wardens-court-summariser/handlers"
	"joshuamURD/wardens-court-summariser/scrape"
	"net/http"

	"github.com/joho/godotenv"
)

// main starts the server and registers the route to the Mux
func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		fmt.Printf("Error loading .env file: %v\n", err)
	}
	scrape.ScrapeDecisions()

	cfg := config.NewAppConfig()

	// Add static file server to serve the public files
	// This is used to serve the CSS, JS, and other static files
	fs := http.FileServer(http.Dir("public"))

	// Create a new router to handle the routes
	router := http.NewServeMux()

	// Add the static file server to the router
	router.Handle("/public/", http.StripPrefix("/public/", fs))

	// Home
	router.HandleFunc("/", handlers.MakeRoute(handlers.HandleIndex))

	fmt.Printf("Listening on %q\n", cfg.GetServeAddr())
	http.ListenAndServe(cfg.GetServeAddr(), router)
}
