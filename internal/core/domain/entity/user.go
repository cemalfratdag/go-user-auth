package entity

import (
	"cfd/myapp/internal/core/domain/enum"
	"time"
)

type User struct {
	ID                         int         `json:"id" gorm:"primaryKey"`
	Name                       string      `json:"name" gorm:"type:varchar(255)"`
	Email                      string      `json:"email" gorm:"type:varchar(255)"`
	Password                   string      `json:"password" gorm:"type:varchar(255)"`
	Birthdate                  string      `json:"birthdate" gorm:"type:date;default:null"`
	Age                        int         `json:"age" gorm:"type:int;default:0"`
	Fee                        float64     `json:"fee" gorm:"type:decimal;default:0"`
	VerificationCode           *string     `json:"verificationCode" gorm:"type:varchar(255);default:null"`
	VerificationCodeCreatedAt  *time.Time  `json:"codeCreatedAt;default:null"`
	PasswordResetCode          *string     `json:"passwordResetCode" gorm:"type:varchar(255);default:null"`
	PasswordResetCodeCreatedAt *time.Time  `json:"passwordResetCodeCreatedAt" gorm:"default:null"`
	Status                     enum.Status `json:"status" gorm:"type:smallint;default:0"`
	CreatedAt                  time.Time   `json:"createdAt"`
	UpdatedAt                  time.Time   `json:"updatedAt"`
	DeletedAt                  *time.Time  `json:"deletedAt" gorm:"index;default:null"`
	Roles                      []Role      `gorm:"many2many:user_roles;"`
}
