package services

import (
	"VeterinaryClinic/config"
	"VeterinaryClinic/model"
	"errors"
	"fmt"
)

func GetAllPets() (*[]model.Pet, error) {
	var pets []model.Pet
	result := config.DB.Find(&pets)
	if result.Error != nil {
		return nil, result.Error
	}

	return &pets, nil
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
	fmt.Printf("Recebido para criação: %+v\n", pet)
	result := config.DB.Create(pet)
	return result.Error
}

func UpdatePet(id uint, updatedPet *model.Pet) (*model.Pet, error) {
	var pet model.Pet
	result := config.DB.First(&pet, id)
	if result.Error != nil {
		return nil, errors.New("Pet not found")
	}
	updates := make(map[string]interface{})

	if updatedPet.Name != "" {
		updates["Name"] = updatedPet.Name
	}
	if updatedPet.Age != 0 {
		updates["Age"] = updatedPet.Age
	}
	if updatedPet.Species != "" {
		updates["Species"] = updatedPet.Species
	}
	if updatedPet.Owner != "" {
		updates["Owner"] = updatedPet.Owner
	}
	if len(updates) > 0 {
		result = config.DB.Model(&pet).Updates(updates)
		if result.Error != nil {
			return nil, result.Error
		}
	}
	return &pet, nil
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
