package service

import (
	"cfd/myapp/internal/common"
	"cfd/myapp/internal/core/domain/dto"
	"cfd/myapp/internal/core/domain/entity"
	"cfd/myapp/internal/core/domain/enum"
	"cfd/myapp/internal/core/port"
	"cfd/myapp/internal/helper"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type UserService struct {
	userRepository port.UserRepository
	roleRepository port.RoleRepository
}

func NewUserService(userRepository port.UserRepository, roleRepository port.RoleRepository) *UserService {
	return &UserService{
		userRepository: userRepository,
		roleRepository: roleRepository,
	}
}

func (s *UserService) CreateUser(userRequest dto.CreateUserRequest) (*int, error) {
	_, err := s.userRepository.GetUserByEmail(userRequest.Email)
	if err == nil {
		return nil, common.ErrInvalidEmail
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(userRequest.Password), 10)
	if err != nil {
		return nil, err
	}

	var user entity.User
	user.Email = userRequest.Email
	user.Password = string(hash)
	user.Name = userRequest.Name
	user.Birthdate = userRequest.Birthdate

	user.Age = helper.CalculateAge(user.Birthdate)
	user.Fee = helper.CalculateFee(user.Age)

	code := helper.CreateVerificationCode()
	user.VerificationCode = &code

	now := time.Now()
	user.VerificationCodeCreatedAt = &now

	userID, err := s.userRepository.CreateUser(user)
	if err != nil {
		return nil, err
	}
	return userID, nil
}

func (s *UserService) GetUserByID(ID int) (*dto.UserResponse, error) {
	user, err := s.userRepository.GetUserByID(ID)
	if err != nil {
		return nil, common.ErrUserNotFound
	}

	if user.Status != enum.UserActive {
		return nil, common.ErrInactiveUser
	}

	return &dto.UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Password:  user.Password,
		Birthdate: user.Birthdate,
		Age:       user.Age,
		Fee:       user.Fee,
	}, nil
}

func (s *UserService) UpdateUser(userID int, userRequest dto.UpdateUserRequest) error {
	existingUser, err := s.userRepository.GetUserByIDWithRole(userID)

	if err != nil {
		return common.ErrUserNotFound //user does not exist
	}

	if existingUser.Status != enum.UserActive {
		return common.ErrInactiveUser
	}

	if userRequest.Email != existingUser.Email { // EMAIL IS BEING UPDATED
		_, err := s.userRepository.GetUserByEmail(userRequest.Email)
		if err == nil {
			return common.ErrInvalidEmail
		}
		existingUser.Email = userRequest.Email
	}

	if userRequest.Name != existingUser.Name { // NAME IS BEING UPDATED
		existingUser.Name = userRequest.Name
	}

	if userRequest.Birthdate != existingUser.Birthdate { // BIRTHDATE IS BEING UPDATED
		existingUser.Birthdate = userRequest.Birthdate
		existingUser.Age = helper.CalculateAge(userRequest.Birthdate)
		existingUser.Fee = helper.CalculateFee(existingUser.Age)
	}

	existingRoleIDList := make([]int, len(existingUser.Roles))
	for i, role := range existingUser.Roles {
		existingRoleIDList[i] = role.ID
	}

	if !helper.CompareIntSlice(existingRoleIDList, userRequest.RoleList) { // ROLES BEING UPDATED
		roles, err := s.roleRepository.GetRolesByIDs(userRequest.RoleList)
		if err != nil {
			return err
		}
		existingUser.Roles = roles
	}

	err = s.userRepository.UpdateUser(*existingUser)
	if err != nil {
		return err
	}
	return nil
}

func (s *UserService) DeleteUser(ID int) error {
	user, err := s.userRepository.GetUserByID(ID)
	if err != nil {
		return common.ErrUserNotFound
	}

	if user.Status != enum.UserActive {
		return common.ErrInactiveUser
	}

	err = s.userRepository.DeleteUser(ID)
	if err != nil {
		return err
	}

	return nil
}

func (s *UserService) GetUserRoleByID(ID int) ([]entity.Role, error) {
	user, err := s.userRepository.GetUserByIDWithRole(ID)
	if err != nil {
		return nil, common.ErrUserNotFound
	}

	if user.Status != enum.UserActive {
		return nil, common.ErrInactiveUser
	}

	return user.Roles, nil
}
