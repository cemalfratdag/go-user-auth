package validations

import (
	"cfd/myapp/internal/core/domain/dto"
	"github.com/go-playground/validator/v10"
	"regexp"
	"time"
)

func ValidBirthdate(fl validator.FieldLevel) bool {
	date := fl.Field().String()

	//check if the date is valid
	parsedDate, err := time.Parse("2006-01-02", date)
	if err != nil {
		return false
	}

	// Check if the date is bigger than 1900-01-01
	minDate, _ := time.Parse("2006-01-02", "1900-01-01")
	if parsedDate.Before(minDate) {
		return false
	}

	//check 18 or older
	age := time.Since(parsedDate).Hours() / 24 / 365

	return age >= 18
}

func ValidPassword(fl validator.FieldLevel) bool {
	password := fl.Field().String()

	flag := true
	tests := []string{".{8,}", "[a-z]", "[A-Z]", "[0-9]"}
	for _, test := range tests {
		match, _ := regexp.MatchString(test, password)
		if !match {
			flag = false
			break
		}
	}
	return flag
}

func PasswordsMatch(fl validator.FieldLevel) bool {
	parentStruct := fl.Parent().Interface()

	switch req := parentStruct.(type) {
	case dto.ResetPasswordRequest:
		return req.Password == req.PasswordConfirm
	case dto.ChangePasswordRequest:
		return req.NewPassword == req.NewPasswordConfirm
	default:
		// If the struct type is not supported, return false
		return false
	}
}
