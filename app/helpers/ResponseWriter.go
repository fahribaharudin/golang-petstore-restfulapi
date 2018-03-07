package helpers

import (
	"encoding/json"
	"net/http"
)

// WriteErrorResponse with http status and message as json format
func WriteErrorResponse(w http.ResponseWriter, status int, message string) error {
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

// WriteJSONResponse with http status and json conten
func WriteJSONResponse(w http.ResponseWriter, status int, content map[string]interface{}) error {
	output, err := json.Marshal(content)
	if err != nil {
		return err
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(output))

	return nil
}
