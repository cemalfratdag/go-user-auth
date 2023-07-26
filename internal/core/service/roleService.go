package service

import (
	"cfd/myapp/internal/common"
	"cfd/myapp/internal/core/domain/dto"
	"cfd/myapp/internal/core/domain/entity"
	"cfd/myapp/internal/core/port"
)

type RoleService struct {
	roleRepository port.RoleRepository
}

func NewRoleService(roleRepository port.RoleRepository) *RoleService {
	return &RoleService{
		roleRepository: roleRepository,
	}
}

func (s *RoleService) CreateRole(roleRequest dto.RoleRequest) (*int, error) {
	_, err := s.roleRepository.GetRoleByCode(roleRequest.Code)
	if err == nil {
		return nil, common.ErrInvalidRoleCode
	}

	_, err = s.roleRepository.GetRoleByName(roleRequest.Name)
	if err == nil {
		return nil, common.ErrInvalidRoleName
	}

	var role entity.Role
	role.Name = roleRequest.Name
	role.Code = roleRequest.Code

	roleID, err := s.roleRepository.CreateRole(role)
	if err != nil {
		return nil, err
	}

	return roleID, nil
}

func (s *RoleService) GetRoleByID(ID int) (*dto.RoleResponse, error) {
	role, err := s.roleRepository.GetRoleByID(ID)
	if err != nil {
		return nil, common.ErrRoleNotFound
	}

	return &dto.RoleResponse{
		Name: role.Name,
		Code: role.Code}, nil
}

func (s *RoleService) UpdateRole(roleID int, roleRequest dto.RoleRequest) error {
	existingRole, err := s.roleRepository.GetRoleByID(roleID)
	if err != nil {
		return common.ErrRoleNotFound
	}

	if roleRequest.Name != existingRole.Name { // NAME IS BEING UPDATED
		_, err := s.roleRepository.GetRoleByName(roleRequest.Name)
		if err == nil {
			return common.ErrInvalidRoleName
		}
		existingRole.Name = roleRequest.Name
	}

	if roleRequest.Code != existingRole.Code { // CODE IS BEING UPDATED
		_, err := s.roleRepository.GetRoleByCode(roleRequest.Code)
		if err == nil {
			return common.ErrInvalidRoleCode
		}
		existingRole.Code = roleRequest.Code
	}

	err = s.roleRepository.UpdateRole(*existingRole)
	if err != nil {
		return err
	}
	return nil
}

func (s *RoleService) DeleteRole(ID int) error {
	_, err := s.roleRepository.GetRoleByID(ID)
	if err != nil {
		return common.ErrRoleNotFound
	}

	err = s.roleRepository.DeleteRole(ID)
	if err != nil {
		return err
	}

	return nil
}
