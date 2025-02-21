package config

import (
	"log"
	"os"
)

// GetPort retrieves the server port from environment variables or defaults to 8080
func GetPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Server running on port %s", port)
	return port
}
