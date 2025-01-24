package services

import (
	"VeterinaryClinic/config"
	"VeterinaryClinic/model"
	"errors"
)

func GetAllPets() ([]model.Pet, error) {
	var pets []model.Pet
	result := config.DB.Find(&pets)
	if result.Error != nil {
		return nil, result.Error
	}
	return pets, nil
}

func GetPetByID(id uint) (*model.Pet, error) {
	var pet model.Pet
	result := config.DB.First(&pet, id)
	if result.Error != nil {
		return nil, errors.New("pet não encontrado")
	}
	return &pet, nil
}
