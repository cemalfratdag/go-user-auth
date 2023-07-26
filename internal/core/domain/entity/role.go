package entity

type Role struct {
	ID   int    `json:"id" gorm:"primaryKey"`
	Name string `json:"roleName" gorm:"type:varchar(255)"`
	Code int    `json:"roleCode" gorm:"type:int"`
}
