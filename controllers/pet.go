package controllers

import (
	"VeterinaryClinic/model"
	"VeterinaryClinic/services"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetPets(w http.ResponseWriter, r *http.Request) {
	pets, err := services.GetAllPets()
	if err != nil {
		http.Error(w, "Erro ao buscar os pets", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(pets)
}

func GetPetById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	pet, err := services.GetPetByID(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(pet)
}

func CreatePet(w http.ResponseWriter, r *http.Request) {
	var pet model.Pet
	err := json.NewDecoder(r.Body).Decode(&pet)
	if err != nil {
		http.Error(w, "Dados inválidos", http.StatusBadRequest)
		return
	}

	err = services.CreatePet(&pet)
	if err != nil {
		http.Error(w, "Erro ao criar o pet", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(pet)
}

func UpdatePet(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	var updatedPet model.Pet
	err = json.NewDecoder(r.Body).Decode(&updatedPet)
	if err != nil {
		http.Error(w, "Dados inválidos", http.StatusBadRequest)
		return
	}

	err = services.UpdatePet(uint(id), &updatedPet)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedPet)
}

func DeletePet(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	err = services.DeletePet(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
