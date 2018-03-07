package helpers

import (
	"errors"
	"reflect"
)

// ValidateRequired check if that a field in a request data must be existed
func ValidateRequired(requestData map[string]interface{}, fieldName string) (bool, error) {
	var value interface{}
	var err error

	value = requestData[fieldName]

	// if there is no field value in the checked request
	if value == nil {
		err = errors.New(fieldName + " field is required.")
		return false, err
	}

	// check type of the value then
	// make the interface assertion type and assign to its value
	valueType := reflect.TypeOf(value)
	switch valueType.Kind() {
	case reflect.Bool:
		value = value.(bool)
		break
	case reflect.Float64:
		value = value.(float64)
		break
	case reflect.String:
		value = value.(string)
		break
	case reflect.Slice:
		break
	default:
		value = nil
	}

	// then validate the value (its weird that zero value of float64 is 0.0 not 0 :D)
	if value == false || value == 0.0 || value == "" || value == nil {
		err = errors.New(fieldName + " value is not valid.")
		return false, err
	} else if valueType.Kind() == reflect.Slice { // validate slice type field
		vs := []interface{}{}
		for _, v := range value.([]interface{}) {
			vs = append(vs, v)
		}

		if len(vs) == 0 {
			err = errors.New(fieldName + " value is emtpy.")
			return false, err
		}
	}

	return true, nil
}

// ValidateRequiredMany validate a slice of input
func ValidateRequiredMany(requestData map[string]interface{}, sliceInput []string) (bool, []error) {
	var errorBag []error

	for _, fieldName := range sliceInput {
		if valid, err := ValidateRequired(requestData, fieldName); !valid {
			errorBag = append(errorBag, err)
		}
	}

	if len(errorBag) > 0 {
		return false, errorBag
	}

	return true, []error{}
}
