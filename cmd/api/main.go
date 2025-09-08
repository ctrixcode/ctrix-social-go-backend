package main

import (
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {

	mainRouter := chi.NewRouter()

	loadEnvironment()

	port := os.Getenv("PORT")

	if port == "" {
		port = "4000"
	}

	mainRouter.Route("/api", func(r chi.Router) {
		r.Get("/healthz", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Ctrix Social Backend!"))
		})
	})

	http.ListenAndServe(":"+port, mainRouter)
}

func loadEnvironment() {
	// Load the .env file in the current directory
	godotenv.Load()

	// db.ResetDB()
	// db.CreateInitialDBStructure()
}
