package routes

import (
	"VeterinaryClinic/controllers"

	"github.com/gorilla/mux"
)

func RegisterPetRoutes(router *mux.Router) {
	router.HandleFunc("/pets", controllers.GetPets).Methods("GET")
	router.HandleFunc("/pets/{id:[0-9]+}", controllers.GetPetById).Methods("GET")
	router.HandleFunc("/pets", controllers.CreatePet).Methods("POST")
	router.HandleFunc("/pets/{id:[0-9]+}", controllers.UpdatePet).Methods("PUT")
	router.HandleFunc("/pets/{id:[0-9]+}", controllers.DeletePet).Methods("DELETE")
}
