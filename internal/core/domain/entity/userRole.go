package entity

type UserRole struct {
	UserID int `json:"userID" gorm:"primaryKey"`
	RoleID int `json:"roleID" gorm:"primaryKey"`
}
