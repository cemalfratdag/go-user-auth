package dto

type RoleResponse struct {
	Name string `json:"name"`
	Code int    `json:"code"`
}

type RoleRequest struct {
	Name string `json:"name" validate:"required"`
	Code int    `json:"code" validate:"required"`
}
