package utilities

import (
	"github.com/go-playground/validator/v10"
)

func ValidateStruct(data any) error {
	var validate = validator.New()
	return validate.Struct(data)
}
