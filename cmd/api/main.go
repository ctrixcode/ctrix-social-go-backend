package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/mcctrix/ctrix-social-go-backend/pkg/security"
)

func main() {

	// mainRouter := chi.NewRouter()

	loadEnvironment()

	port := os.Getenv("PORT")

	if port == "" {
		port = "4000"
	}

	// http.ListenAndServe(":"+port, mainRouter)
}

func loadEnvironment() {
	// Load the .env file in the current directory
	godotenv.Load()

	if _, err := os.Stat("./ecdsa_private_key.pem"); err == nil {
		// Pem File exist so do nothing
	} else {
		security.GenerateEcdsaPrivateKey()
	}
	// db.ResetDB()
	// db.CreateInitialDBStructure()
}
