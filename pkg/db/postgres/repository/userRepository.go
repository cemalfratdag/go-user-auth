package repository

import (
	"cfd/myapp/internal/core/domain/entity"
	"cfd/myapp/internal/core/domain/enum"
	"gorm.io/gorm"
	"time"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) CreateUser(user entity.User) (*int, error) {
	if err := r.db.Create(&user).Error; err != nil {
		return nil, err
	}
	return &user.ID, nil
}

func (r *UserRepository) GetUserByID(ID int) (*entity.User, error) {
	user := &entity.User{}
	if err := r.db.Where("id = ?", ID).First(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) GetUserByEmail(email string) (*entity.User, error) {
	user := &entity.User{}
	if err := r.db.Where("email = ?", email).First(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) UpdateUser(user entity.User) error {
	user.UpdatedAt = time.Now()

	//userMap := map[string]interface{}{
	//	"name":                           user.Name,
	//	"email":                          user.Email,
	//	"password":                       user.Password,
	//	"birthdate":                      user.Birthdate,
	//	"age":                            user.Age,
	//	"fee":                            user.Fee,
	//	"status":                         user.Status,
	//	"verification_code":              user.VerificationCode,
	//	"verification_code_created_at":   user.VerificationCodeCreatedAt,
	//	"password_reset_code":            user.PasswordResetCode,
	//	"password_reset_code_created_at": user.PasswordResetCodeCreatedAt,
	//	"updated_at":                     user.UpdatedAt,
	//	"roles":                          user.Roles,
	//}
	//
	//err := r.db.Model(&user).Where("id = ?", user.ID).Updates(userMap).Error
	//if err != nil {
	//	return err
	//}
	//return nil

	err := r.db.Model(&user).Where("id = ?", user.ID).Updates(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) DeleteUser(ID int) error {
	result := r.db.Model(&entity.User{}).Where("id = ?", ID).
		Updates(map[string]interface{}{
			"deleted_at": time.Now(),
			"status":     enum.UserDeleted,
		}).Omit("updated_at")

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *UserRepository) GetUserByIDWithRole(ID int) (*entity.User, error) {
	user := &entity.User{}

	err := r.db.Model(user).Preload("Roles").Where("id = ?", ID).First(user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}
