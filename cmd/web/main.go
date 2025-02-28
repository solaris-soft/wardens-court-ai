package main

import (
	"fmt"
	"joshuamURD/wardens-court-summariser/config"
	"joshuamURD/wardens-court-summariser/internal/database"
	"joshuamURD/wardens-court-summariser/internal/handlers"
	"os"

	"github.com/joho/godotenv"
)

// main starts the server and registers the route to the Mux
func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		fmt.Printf("Error loading .env file: %v\n", err)
	}

	db, err := database.NewSQLITEDB("wardens-court-summariser.db", "uploads")
	if err != nil {
		fmt.Printf("Error creating database: %v\n", err)
	}

	r := []config.Route{
		{Path: "/", Handler: handlers.MakeRoute(handlers.HandleIndex)},
		{Path: "/upload", Handler: handlers.MakeRoute(handlers.UploadFile(db))},
		{Path: "/table", Handler: handlers.MakeRoute(handlers.HandleTable(db))},
	}

	address := config.WithAddr(os.Getenv("ADDRESS"))
	port := config.WithPort(os.Getenv("PORT"))
	routes := config.WithRoutes(&r)
	cfg := config.NewAppConfig(address, port, routes)

	fmt.Printf("Listening on %s\n", cfg.Addr)
	cfg.ListenAndServe()
}
