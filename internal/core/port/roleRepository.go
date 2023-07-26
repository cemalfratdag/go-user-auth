package port

import "cfd/myapp/internal/core/domain/entity"

type RoleRepository interface {
	CreateRole(user entity.Role) (*int, error)
	GetRoleByID(ID int) (*entity.Role, error)
	GetRolesByIDs(roleIDs []int) ([]entity.Role, error)
	GetRoleByCode(code int) (*entity.Role, error)
	GetRoleByName(name string) (*entity.Role, error)
	UpdateRole(user entity.Role) error
	DeleteRole(ID int) error
}
