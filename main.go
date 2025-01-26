package main

import (
	"VeterinaryClinic/config"
	"VeterinaryClinic/model"
	"VeterinaryClinic/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	db := config.ConnectDatabase()

	model.MigrateDB(db)

	router := mux.NewRouter()

	routes.RegisterPetRoutes(router)

	log.Println("Server running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", router))
}
