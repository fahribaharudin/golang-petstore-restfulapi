package helpers

import (
	"encoding/json"
	"io"
)

// ResponseBodyHelper ..
type ResponseBodyHelper struct{}

// BodyParser create a response body helper object
func BodyParser() ResponseBodyHelper {
	return ResponseBodyHelper{}
}

// ParseJSON parsing json body data to a map containing request data
func (h ResponseBodyHelper) ParseJSON(requestBody io.Reader) (map[string]interface{}, error) {
	var requestData map[string]interface{}

	decoder := json.NewDecoder(requestBody)
	err := decoder.Decode(&requestData)
	if err != nil {
		return requestData, err
	}

	return requestData, nil
}
