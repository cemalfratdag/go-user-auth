package authjwt

import (
	"cfd/myapp/config"
	"cfd/myapp/internal/core/domain/dto"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type JwtToken struct {
	accessSecret    string
	refreshSecret   string
	accessInterval  int
	refreshInterval int
}

func NewJwtToken(jwtConfig config.JwtConfig) JwtToken {
	return JwtToken{
		accessSecret:    jwtConfig.AccessSecret,
		refreshSecret:   jwtConfig.RefreshSecret,
		accessInterval:  jwtConfig.AccessInterval,
		refreshInterval: jwtConfig.RefreshInterval,
	}
}

func (t *JwtToken) ValidateAccessToken(token string) (*dto.TokenUser, error) {
	return t.validateToken(token, t.accessSecret)
}

func (t *JwtToken) ValidateRefreshToken(token string) (*dto.TokenUser, error) {
	return t.validateToken(token, t.refreshSecret)
}

func (t *JwtToken) validateToken(tokenString string, secret string) (*dto.TokenUser, error) {
	claims := &dto.Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}

	return &claims.TokenUser, nil
}

func (t *JwtToken) GenerateAccessToken(id int) (*string, error) {
	expirationTime := time.Now().Local().Add(time.Duration(t.accessInterval) * time.Minute).Unix()

	claims := &dto.Claims{
		TokenUser: dto.TokenUser{ID: id},
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessToken, err := token.SignedString([]byte(t.accessSecret))
	if err != nil {
		return nil, err
	}

	return &accessToken, nil
}

func (t *JwtToken) GenerateRefreshToken(id int) (*string, error) {
	expirationTime := time.Now().Local().Add(time.Duration(t.refreshInterval) * time.Minute).Unix()

	refreshClaims := &dto.Claims{
		TokenUser: dto.TokenUser{ID: id},
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshToken, err := token.SignedString([]byte(t.refreshSecret))
	if err != nil {
		return nil, err
	}

	return &refreshToken, err
}

func (t *JwtToken) GenerateAllTokens(id int) (tokenDTO *dto.TokenResponse, err error) {
	accessTokenString, err := t.GenerateAccessToken(id)
	if err != nil {
		return nil, err
	}

	refreshTokenString, err := t.GenerateRefreshToken(id)
	if err != nil {
		return nil, err
	}

	return &dto.TokenResponse{
		AccessToken:  *accessTokenString,
		RefreshToken: *refreshTokenString,
	}, err
}
