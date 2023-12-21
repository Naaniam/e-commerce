package utilities

import (
	"github.com/go-playground/validator/v10"
)

// ValidateStruct function is used to validate the models with required fields
func ValidateStruct(data any) error {
	var validate = validator.New()
	return validate.Struct(data)
}
