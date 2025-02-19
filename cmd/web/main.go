package main

import (
	"fmt"
	"joshuamURD/wardens-court-summariser/config"
	"joshuamURD/wardens-court-summariser/internal/handlers"
	"joshuamURD/wardens-court-summariser/scrape"
	"os"

	"github.com/joho/godotenv"
)

// main starts the server and registers the route to the Mux
func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		fmt.Printf("Error loading .env file: %v\n", err)
	}
	scrape.ScrapeDecisions()

	r := []config.Route{
		{Path: "/", Handler: handlers.MakeRoute(handlers.HandleIndex)},
	}

	address := config.WithAddr(os.Getenv("ADDRESS"))
	port := config.WithPort(os.Getenv("PORT"))
	routes := config.WithRoutes(r)
	cfg := config.NewAppConfig(address, port, routes)

	fmt.Printf("Listening on %s\n", cfg.Addr)
	cfg.ListenAndServe()
}
