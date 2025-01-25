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
		return nil, errors.New("Pet not found")
	}
	return &pet, nil
}

func CreatePet(pet *model.Pet) error {
	result := config.DB.Create(pet)
	return result.Error
}

func UpdatePet(id uint, updatedPet *model.Pet) error {
	var pet model.Pet
	result := config.DB.First(&pet, id)
	if result.Error != nil {
		return errors.New("Pet not found")
	}
	pet.Name = updatedPet.Name
	pet.Age = updatedPet.Age
	pet.Species = updatedPet.Species
	pet.Owner = updatedPet.Owner

	result = config.DB.Save(&pet)
	return result.Error
}

func DeletePet(id uint) error {
	var pet model.Pet
	result := config.DB.First(&pet, id)
	if result.Error != nil {
		return errors.New("Pet not found")
	}

	result = config.DB.Delete(&pet)
	return result.Error
}
