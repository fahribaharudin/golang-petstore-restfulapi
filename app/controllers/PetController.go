package controllers

import (
	"encoding/json"
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
	requestData, err := helpers.ParseJSON(r.Body)
	if err != nil {
		helpers.WriteErrorResponse(w, http.StatusBadRequest, "Invalid json request data.")
		return
	}

	// validate (required) request data
	inputToBeValidated := []string{"category_id", "name", "photoUrls", "tag_id", "status"}
	valid, errorBag := helpers.ValidateRequiredMany(requestData, inputToBeValidated)
	if !valid && len(errorBag) > 0 {
		helpers.WriteErrorResponse(w, http.StatusBadRequest, errorBag[0].Error())
		return
	}

	// wrap the request data that has been validated
	validatedRequestData := map[string]interface{}{
		"categoryID": requestData["category_id"],
		"name":       requestData["name"],
		"photoUrls":  requestData["photoUrls"],
		"tagID":      requestData["tag_id"],
		"status":     requestData["status"],
	}

	// store the pet
	err = c.PetService.StorePet(validatedRequestData)
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		w.Write([]byte(err.Error()))

		return
	}
	output, _ := json.Marshal(validatedRequestData)

	// write the data to response
	w.Header().Set("content-type", "application/json")
	w.Write(output)
}
