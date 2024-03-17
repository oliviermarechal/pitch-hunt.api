package main

import (
	"log"
	"net/http"

	"pitch-hunt/src/adapters/primary/router"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erreur lors du chargement du fichier .env")
	}

	router := router.SetupRouter()

	http.ListenAndServe(":8000", router)
}
