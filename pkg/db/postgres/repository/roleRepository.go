package repository

import (
	"cfd/myapp/internal/core/domain/entity"
	"gorm.io/gorm"
)

type RoleRepository struct {
	db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) *RoleRepository {
	return &RoleRepository{
		db: db,
	}
}

func (r *RoleRepository) CreateRole(role entity.Role) (*int, error) {
	if err := r.db.Create(&role).Error; err != nil {
		return nil, err
	}

	return &role.ID, nil
}

func (r *RoleRepository) GetRoleByID(ID int) (*entity.Role, error) {
	role := &entity.Role{}
	if err := r.db.Where("id = ?", ID).First(role).Error; err != nil {
		return nil, err
	}

	return role, nil
}

func (r *RoleRepository) GetRolesByIDs(roleIDs []int) ([]entity.Role, error) {
	var roles []entity.Role

	err := r.db.Where("id IN (?)", roleIDs).Find(&roles).Error
	if err != nil {
		return nil, err
	}

	return roles, nil
}

func (r *RoleRepository) GetRoleByCode(code int) (*entity.Role, error) {
	role := &entity.Role{}
	if err := r.db.Where("code = ?", code).First(role).Error; err != nil {
		return nil, err
	}

	return role, nil
}

func (r *RoleRepository) GetRoleByName(name string) (*entity.Role, error) {
	role := &entity.Role{}
	if err := r.db.Where("name = ?", name).First(role).Error; err != nil {
		return nil, err
	}

	return role, nil
}

func (r *RoleRepository) UpdateRole(role entity.Role) error {

	roleMap := map[string]interface{}{
		"name": role.Name,
		"code": role.Code,
	}

	if err := r.db.Model(&role).Where("id = ?", role.ID).Updates(roleMap).Error; err != nil {
		return err
	}

	return nil
}

func (r *RoleRepository) DeleteRole(ID int) error {
	result := r.db.Where("id = ?", ID).Delete(&entity.Role{})
	if result.Error != nil {
		return result.Error
	}

	return nil
}
