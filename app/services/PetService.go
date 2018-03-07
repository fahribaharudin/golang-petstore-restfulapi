package services

import (
	"github.com/fahribaharudin/petstore_restapi/app/models"
	"github.com/fahribaharudin/petstore_restapi/app/repositories"
)

// PetService contains some service for the pet resource
type PetService struct {
	PetRepository repositories.PetRepository
}

// StorePet service..
func (s *PetService) StorePet(requestData map[string]interface{}) error {
	success, err := s.PetRepository.Store(requestData)

	if success {
		return nil
	}

	return err
}

// GetLatestPet from repository
func (s *PetService) GetLatestPet() (models.Pet, error) {
	pet, err := s.PetRepository.GetLastRecord()
	if err == nil {
		return pet, nil
	}

	return pet, err
}
