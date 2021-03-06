package controllers

import (
	"net/http"

	"github.com/fahribaharudin/petstore_restapi/app/helpers"
	"github.com/fahribaharudin/petstore_restapi/app/services"
)

// PetController contains some handler to the /pet endpoint
type PetController struct {
	PetService services.PetService
}

// Store is a handler to POST:/pet
func (c *PetController) Store(w http.ResponseWriter, r *http.Request) {
	// decode request data
	requestData, err := helpers.BodyParser().ParseJSON(r.Body)
	if err != nil {
		helpers.Response().WriteError(w, http.StatusBadRequest, "Invalid json request data.")
		return
	}

	// validate (required) request data
	inputToBeValidated := []string{"category_id", "name", "photoUrls", "tag_id", "status"}
	valid, errorBag := helpers.ValidateRequiredMany(requestData, inputToBeValidated)
	if !valid && len(errorBag) > 0 {
		helpers.Response().WriteError(w, http.StatusBadRequest, errorBag[0].Error())
		return
	}

	// store the pet
	err = c.PetService.StorePet(requestData)
	pet, err := c.PetService.GetLatestPet()
	if err != nil {
		helpers.Response().WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	// write success message to the response
	helpers.Response().WriteJSON(w, http.StatusCreated, struct {
		ID         uint   `json:"id"`
		Name       string `json:"name"`
		Status     string `json:"status"`
		CategoryID uint   `json:"category_id"`
		TagID      uint   `json:"tag_id"`
	}{
		ID:         pet.ID,
		Name:       pet.Name,
		Status:     pet.Status,
		CategoryID: pet.CategoryID,
		TagID:      pet.TagID,
	}, "The pet was successfully created")

	return
}
