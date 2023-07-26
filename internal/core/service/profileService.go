package service

import (
	"cfd/myapp/internal/common"
	"cfd/myapp/internal/core/domain/dto"
	"cfd/myapp/internal/core/domain/enum"
	"cfd/myapp/internal/core/port"
	"cfd/myapp/internal/helper"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type ProfileService struct {
	userRepository port.UserRepository
}

func NewProfileService(userRepository port.UserRepository) *ProfileService {
	return &ProfileService{
		userRepository: userRepository,
	}
}

func (s *ProfileService) ForgotPassword(email string) (*string, error) {
	user, err := s.userRepository.GetUserByEmail(email)
	if err != nil {
		return nil, common.ErrInterval
	}

	if user.Status != enum.UserActive {
		return nil, common.ErrInactiveUser
	}

	resetCode := helper.CreateVerificationCode()
	user.PasswordResetCode = &resetCode

	now := time.Now()
	user.PasswordResetCodeCreatedAt = &now

	err = s.userRepository.UpdateUser(*user)
	if err != nil {
		return nil, common.ErrResetPassword
	}
	return &resetCode, nil
}

func (s *ProfileService) ResetPassword(resetPasswordRequest dto.ResetPasswordRequest) error {
	user, err := s.userRepository.GetUserByEmail(resetPasswordRequest.Email)
	if err != nil {
		return common.ErrFailedPasswordReset
	}

	if user.Status != enum.UserActive {
		return common.ErrInactiveUser
	}

	if resetPasswordRequest.ResetCode != *user.PasswordResetCode {
		return common.ErrFailedPasswordReset
	}

	expirationTime := user.PasswordResetCodeCreatedAt.Add(5 * time.Minute)
	if time.Now().After(expirationTime) {
		return common.ErrPasswordResetExpired
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(resetPasswordRequest.Password), 10)
	if err != nil {
		return common.ErrInterval
	}

	user.Password = string(hash)
	user.PasswordResetCode = nil
	user.PasswordResetCodeCreatedAt = nil

	err = s.userRepository.UpdateUser(*user)
	if err != nil {
		return common.ErrInterval
	}
	return nil
}

func (s *ProfileService) ChangePassword(changePasswordRequest dto.ChangePasswordRequest, userID int) error {
	user, err := s.userRepository.GetUserByID(userID)
	if err != nil {
		return common.ErrFailedPasswordReset
	}

	if user.Status != enum.UserActive {
		return common.ErrInactiveUser
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(changePasswordRequest.OldPassword)); err != nil {
		return common.ErrOldPassword
	}

	newHash, err := bcrypt.GenerateFromPassword([]byte(changePasswordRequest.NewPassword), 10)
	if err != nil {
		return err
	}

	user.Password = string(newHash)
	if err := s.userRepository.UpdateUser(*user); err != nil {
		return common.ErrUpdatePassword
	}

	return nil
}

func (s *ProfileService) ViewProfile(userID int) (*dto.ViewProfileResponse, error) {
	user, err := s.userRepository.GetUserByID(userID)
	if err != nil {
		return nil, err
	}

	if user.Status != enum.UserActive {
		return nil, common.ErrInactiveUser
	}

	return &dto.ViewProfileResponse{
		Name:      user.Name,
		Email:     user.Email,
		Birthdate: user.Birthdate,
		Age:       user.Age,
		Fee:       user.Fee,
	}, nil
}

func (s *ProfileService) UpdateProfile(userID int, userRequest dto.UpdateProfileRequest) error {
	existingUser, err := s.userRepository.GetUserByID(userID)
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

	if userRequest.Birthdate != existingUser.Birthdate {
		existingUser.Birthdate = userRequest.Birthdate
		existingUser.Age = helper.CalculateAge(userRequest.Birthdate)
		existingUser.Fee = helper.CalculateFee(existingUser.Age)
	}

	err = s.userRepository.UpdateUser(*existingUser)
	if err != nil {
		return err
	}
	return nil
}

func (s *ProfileService) DeleteProfile(ID int) error {
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
