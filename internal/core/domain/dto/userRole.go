package dto

type UserRoleRequest struct {
	UserID int `json:"userID" validator:"required"`
	RoleID int `json:"roleID" validator:"required"`
}
