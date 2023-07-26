package port

import "cfd/myapp/internal/core/domain/entity"

type UserRepository interface {
	CreateUser(user entity.User) (*int, error)
	GetUserByID(ID int) (*entity.User, error)
	UpdateUser(user entity.User) error
	DeleteUser(ID int) error
	GetUserByEmail(email string) (*entity.User, error)
	GetUserByIDWithRole(ID int) (*entity.User, error)
}
