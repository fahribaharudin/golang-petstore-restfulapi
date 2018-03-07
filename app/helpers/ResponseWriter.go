package helpers

import (
	"encoding/json"
	"net/http"
)

// ViewModel is the default model to represent json response
type ViewModel struct {
	Message string      `json:"message"`
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
}

// ResponseHelper object
type ResponseHelper struct {
}

// Response is the ResponseHelper object factory
func Response() ResponseHelper {
	return ResponseHelper{}
}

// WriteError response with http status and message as json format
func (h ResponseHelper) WriteError(w http.ResponseWriter, status int, message string) error {
	var msg = map[string]interface{}{
		"code":    status,
		"message": message,
	}

	output, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(output))

	return nil
}

// WriteJSON response with http status and json content from map that will being marshaled
func (h ResponseHelper) WriteJSON(w http.ResponseWriter, status int, content interface{}, msg string) error {
	viewModel := ViewModel{}
	output, err := json.Marshal(&viewModel)
	if err != nil {
		return err
	}

	w.Header().Set("content-type", "application/json")
	w.Write([]byte(output))
	w.WriteHeader(status)

	return nil
}
