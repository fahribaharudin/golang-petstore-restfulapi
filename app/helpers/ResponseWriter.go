package helpers

import (
	"encoding/json"
	"net/http"
)

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
	var viewModel = struct {
		Message string      `json:"message"`
		Code    int         `json:"code"`
		Data    interface{} `json:"data"`
	}{
		Message: msg,
		Code:    status,
		Data:    content,
	}

	output, err := json.Marshal(&viewModel)
	if err != nil {
		return err
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(output))

	return nil
}
