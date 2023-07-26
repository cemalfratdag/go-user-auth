package dto

type CreateUserRequest struct {
	Name      string `json:"name" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,validPassword"`
	Birthdate string `json:"birthdate" validate:"required,validBirthdate"`
}

type UpdateUserRequest struct {
	Name      string `json:"name" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Birthdate string `json:"birthdate" validate:"required,validBirthdate"`
	RoleList  []int  `json:"roleList" validate:"required,dive"`
}

type UserResponse struct {
	ID        int     `json:"id"`
	Name      string  `json:"name"`
	Email     string  `json:"email"`
	Password  string  `json:"password"`
	Birthdate string  `json:"birthdate"`
	Age       int     `json:"age"`
	Fee       float64 `json:"fee"`
}
