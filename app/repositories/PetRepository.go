package repositories

import (
	"github.com/fahribaharudin/petstore_restapi/app/models"
	"github.com/jinzhu/gorm"
)

// PetRepository is a repo..
type PetRepository struct {
	DbHandler *gorm.DB
}

// Store the pet data to db
func (repo PetRepository) Store(data map[string]interface{}) (bool, error) {
	pet := models.Pet{
		Name:       data["name"].(string),
		Status:     data["status"].(string),
		CategoryID: uint(data["categoryID"].(float64)),
		TagID:      uint(data["tagID"].(float64)),
	}

	dbc := repo.DbHandler.Create(&pet)
	if dbc.Error != nil {
		return false, dbc.Error
	}

	return true, nil
}

// GetLastRecord from the database
func (repo *PetRepository) GetLastRecord() (models.Pet, error) {
	pet := models.Pet{}
	dbc := repo.DbHandler.Last(&pet)
	if dbc.Error != nil {
		return pet, dbc.Error
	}

	return pet, nil
}
