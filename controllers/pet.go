package controllers

import (
	"VeterinaryClinic/model"
	"VeterinaryClinic/services"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func writeJSONError(w http.ResponseWriter, statusCode int, message string) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	errResponse := map[string]string{"error": message}
	json.NewEncoder(w).Encode(errResponse)

}
func GetPets(w http.ResponseWriter, r *http.Request) {
	pets, err := services.GetAllPets()
	if err != nil {
		http.Error(w, "Error fetching pets", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(pets)
}

func GetPetById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		writeJSONError(w, http.StatusInternalServerError, "Error fetching pets")
		return
	}

	pet, err := services.GetPetByID(uint(id))
	if err != nil {
		writeJSONError(w, http.StatusNotFound, err.Error())
		return
	}
	json.NewEncoder(w).Encode(pet)
}

func CreatePet(w http.ResponseWriter, r *http.Request) {
	var pet model.Pet
	err := json.NewDecoder(r.Body).Decode(&pet)
	if err != nil {
		writeJSONError(w, http.StatusBadRequest, "Invalid data")
		return
	}

	err = services.CreatePet(&pet)
	if err != nil {
		writeJSONError(w, http.StatusInternalServerError, "Error registering the pet")
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(pet)
}

func UpdatePet(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		writeJSONError(w, http.StatusBadRequest, "Invalid id")
		return
	}

	var updatedPet model.Pet
	err = json.NewDecoder(r.Body).Decode(&updatedPet)
	if err != nil {
		writeJSONError(w, http.StatusBadRequest, "Invalid data")
		return
	}

	err = services.UpdatePet(uint(id), &updatedPet)
	if err != nil {
		writeJSONError(w, http.StatusInternalServerError, err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedPet)
}

func DeletePet(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		writeJSONError(w, http.StatusBadRequest, "Invalid id")
		return
	}

	err = services.DeletePet(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		writeJSONError(w, http.StatusNotFound, err.Error())
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
