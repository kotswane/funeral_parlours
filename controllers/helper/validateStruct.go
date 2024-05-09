package helper

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

// Validator instance
var validate *validator.Validate

// initValidator initializes the validator
func initValidator() {
	if validate == nil {
		validate = validator.New()
	}
}

// ValidateStruct validates any struct passed to it and returns a slice of error strings if validation fails
func ValidateStruct(s interface{}) string {
	initValidator() // Ensure the validator is initialized

	// Validate the struct
	err := validate.Struct(s)
	if err == nil {
		return "" // No errors, return nil
	}

	// Collect all the errors
	var errorMessage string
	for _, err := range err.(validator.ValidationErrors) {
		errorMessage = fmt.Sprintf("Error in field '%s' of type '%s': the condition '%s' failed",
			err.Field(), err.Type().Name(), err.ActualTag())

	}
	return errorMessage
}
