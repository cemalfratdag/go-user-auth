package dto

type ForgotPasswordRequest struct {
	Email string `json:"email" binding:"required"`
}

type ResetPasswordRequest struct {
	Email           string `json:"email" validate:"required,email"`
	ResetCode       string `json:"resetCode" validate:"required,len=10"`
	Password        string `json:"password" validate:"required,validPassword"`
	PasswordConfirm string `json:"passwordConfirm" validate:"required,validPassword,passwordsMatch"`
}

type ChangePasswordRequest struct {
	OldPassword        string `json:"oldPassword" validate:"required,validPassword"`
	NewPassword        string `json:"newPassword" validate:"required,validPassword"`
	NewPasswordConfirm string `json:"newPasswordConfirm" validate:"required,validPassword,passwordsMatch"`
}

type ViewProfileResponse struct {
	Name      string  `json:"name"`
	Email     string  `json:"email"`
	Birthdate string  `json:"birthdate"`
	Age       int     `json:"age"`
	Fee       float64 `json:"fee"`
}

type UpdateProfileRequest struct {
	Name      string `json:"name"`
	Email     string `json:"email" validate:"omitempty,email"`
	Birthdate string `json:"birthdate" validate:"omitempty,validBirthdate"`
}
