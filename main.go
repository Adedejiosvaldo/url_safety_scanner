package main

import (
	"log"

	"github.com/adedejiosvaldo/safe_url/config"
	"github.com/adedejiosvaldo/safe_url/routes"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found; using environment variables")
	}

	port := config.GetPort()
	router := routes.SetupRouter()

	log.Fatal(router.Run(":" + port))

}
