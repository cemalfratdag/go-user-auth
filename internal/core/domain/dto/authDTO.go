package dto

import (
	"github.com/dgrijalva/jwt-go"
)

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,validPassword"`
}

type Claims struct {
	TokenUser
	jwt.StandardClaims
}

// TokenUser todo: add user info
type TokenUser struct {
	ID int `json:"id"`
}

type TokenResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type RefreshTokenResponse struct {
	AccessToken string `json:"accessToken"`
}

type VerifyRequest struct {
	Email string `json:"email" validate:"required,email"`
	Code  string `json:"code" validate:"required,len=10"`
}

type RefreshVerificationCodeRequest struct {
	Email string `json:"email" validate:"required,email"`
}
