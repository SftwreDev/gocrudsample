package main

import (
	"github.com/joho/godotenv"
	"gocrudsample/models"
	"gocrudsample/routers"
)

func main() {
	// Initialize env
	err := godotenv.Load(".env")
	if err != nil {
		return
	}

	// Initialize models
	models.ConnectDatabase()

	// Initialize routers
	routers.InitializeRouter()

}
