package service

import (
	"cfd/myapp/internal/common"
	"cfd/myapp/internal/core/domain/dto"
	"cfd/myapp/internal/core/domain/entity"
	"cfd/myapp/internal/core/domain/enum"
	"cfd/myapp/internal/core/port"
	"cfd/myapp/internal/helper"
	"cfd/myapp/pkg/jwt"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type AuthService struct {
	userRepository port.UserRepository
	jwtConfig      authjwt.JwtToken
}

func NewAuthService(userRepository port.UserRepository, jwtConfig authjwt.JwtToken) *AuthService {
	return &AuthService{
		userRepository: userRepository,
		jwtConfig:      jwtConfig,
	}
}

func (s *AuthService) Login(loginRequest dto.LoginRequest) (*dto.TokenResponse, error) {
	user, err := s.userRepository.GetUserByEmail(loginRequest.Email)
	if err != nil {
		return nil, common.ErrUserNotFound
	}

	if user.Status != enum.UserActive {
		return nil, common.ErrInactiveUser
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password))
	if err != nil {
		return nil, common.ErrLogin
	}

	tokenResponse, err := s.jwtConfig.GenerateAllTokens(user.ID)
	if err != nil {
		return nil, common.ErrCreateToken
	}

	return tokenResponse, nil
}

func (s *AuthService) Signup(userRequest dto.CreateUserRequest) (*int, error) {
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

func (s *AuthService) RefreshToken(refreshToken string) (*dto.RefreshTokenResponse, error) {
	token, err := s.jwtConfig.ValidateRefreshToken(refreshToken)
	if err != nil {
		return nil, err
	}

	user, err := s.userRepository.GetUserByID(token.ID)
	if err != nil {
		return nil, err
	}

	if user.Status != enum.UserActive {
		return nil, common.ErrInactiveUser
	}

	newAccessToken, err := s.jwtConfig.GenerateAccessToken(user.ID)
	if err != nil {
		return nil, err
	}

	return &dto.RefreshTokenResponse{
		AccessToken: *newAccessToken,
	}, nil
}

func (s *AuthService) VerifyUser(verifyRequest dto.VerifyRequest) error {
	user, err := s.userRepository.GetUserByEmail(verifyRequest.Email)
	if err != nil {
		return common.ErrUserNotFound
	}

	if user.Status != enum.UserInactive {
		return common.ErrInvalidCode
	}

	if verifyRequest.Code != *user.VerificationCode {
		return common.ErrInvalidCode
	}

	expirationTime := user.VerificationCodeCreatedAt.Add(5 * time.Minute)
	if time.Now().After(expirationTime) {
		return common.ErrInvalidCode
	}

	user.Status = 1
	user.VerificationCode = nil
	user.VerificationCodeCreatedAt = nil

	err = s.userRepository.UpdateUser(*user)
	if err != nil {
		return err
	}

	return nil
}

func (s *AuthService) RefreshVerificationCode(email string) (*string, error) {
	user, err := s.userRepository.GetUserByEmail(email)
	if err != nil {
		return nil, common.ErrUserNotFound
	}

	if user.Status != enum.UserInactive {
		return nil, common.ErrAlreadyVerified
	}

	code := helper.CreateVerificationCode()
	user.VerificationCode = &code

	now := time.Now()
	user.VerificationCodeCreatedAt = &now

	err = s.userRepository.UpdateUser(*user)
	if err != nil {
		return nil, err
	}

	return user.VerificationCode, nil
}
