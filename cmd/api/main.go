package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/mcctrix/ctrix-social-go-backend/pkg/database"
)

func main() {

	// mainRouter := chi.NewRouter()

	loadEnvironment()

	port := os.Getenv("PORT")

	if port == "" {
		port = "4000"
	}
	database.CreateInitialDBStructure()

	// http.ListenAndServe(":"+port, mainRouter)
}

func loadEnvironment() {
	// Load the .env file in the current directory
	godotenv.Load()

	// db.ResetDB()
	// db.CreateInitialDBStructure()
}
