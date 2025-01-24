package main

import (
	"VeterinaryClinic/config"
	"VeterinaryClinic/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	config.ConnectDatabase()

	router := mux.NewRouter()

	routes.RegisterPetRoutes(router)

	log.Println("Server running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", router))
}
