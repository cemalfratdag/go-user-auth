package validations

import (
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
	_ = validate.RegisterValidation("validBirthdate", ValidBirthdate)
	_ = validate.RegisterValidation("validPassword", ValidPassword)
	_ = validate.RegisterValidation("passwordsMatch", PasswordsMatch)
}

func Validate(data interface{}) map[string]string {
	err := validate.Struct(data)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		errorMap := make(map[string]string)
		for _, err := range validationErrors {
			field := err.StructField()
			message := err.Tag()
			errorMap[field] = message
		}
		return errorMap
	}
	return nil
}
