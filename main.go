package main

import (
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"github.com/mcctrix/ctrix-social-go-backend/utils"
)

func main() {

	mainRouter := chi.NewRouter()

	loadEnvironment()

	port := os.Getenv("PORT")

	if port == "" {
		port = "4000"
	}

	http.ListenAndServe(":"+port, mainRouter)
}

func loadEnvironment() {
	// Load the .env file in the current directory
	godotenv.Load()

	if _, err := os.Stat("./ecdsa_private_key.pem"); err == nil {
		// Pem File exist so do nothing
	} else {
		utils.GenerateEcdsaPrivateKey()
	}
	// db.ResetDB()
	// db.CreateInitialDBStructure()
}
