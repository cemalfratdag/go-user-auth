package handler

import (
	"cfd/myapp/internal/common"
	"cfd/myapp/internal/core/domain/dto"
	"cfd/myapp/internal/core/service"
	"cfd/myapp/internal/helper"
	validations "cfd/myapp/pkg/validator"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthHandler struct {
	authService service.AuthService
}

func NewAuthHandler(authService service.AuthService) AuthHandler {
	return AuthHandler{
		authService: authService,
	}
}

func (h *AuthHandler) Signup(c *gin.Context) {
	var userRequest dto.CreateUserRequest
	err := c.ShouldBindJSON(&userRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.ResponseError(err.Error()))
		return
	}

	validationErrors := validations.Validate(userRequest)
	if validationErrors != nil {
		c.JSON(http.StatusBadRequest, helper.ResponseError(validationErrors))
		return
	}

	userID, err := h.authService.Signup(userRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.ResponseError(err.Error()))
		return
	}

	c.JSON(http.StatusCreated, helper.ResponseSuccess(userID))
}

func (h *AuthHandler) Login(c *gin.Context) {
	var loginRequest dto.LoginRequest
	err := c.ShouldBindJSON(&loginRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.ResponseError(err.Error()))
		return
	}

	validationErrors := validations.Validate(loginRequest)
	if validationErrors != nil {
		c.JSON(http.StatusBadRequest, helper.ResponseError(validationErrors))
		return
	}

	token, err := h.authService.Login(loginRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.ResponseError(err.Error()))
		return
	}

	c.JSON(http.StatusOK, helper.ResponseSuccess(token))
}

func (h *AuthHandler) RefreshToken(c *gin.Context) {
	refreshToken := c.GetHeader("Refresh-Token")
	if refreshToken == "" {
		c.JSON(http.StatusBadRequest, helper.ResponseError(common.ErrTokenMissing))
		return
	}

	newToken, err := h.authService.RefreshToken(refreshToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, helper.ResponseError(err.Error()))
		return
	}

	c.JSON(http.StatusOK, helper.ResponseSuccess(newToken))
}

func (h *AuthHandler) VerifyUser(c *gin.Context) {
	var verifyRequest dto.VerifyRequest
	err := c.ShouldBindJSON(&verifyRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.ResponseError(err.Error()))
		return
	}

	validationErrors := validations.Validate(verifyRequest)
	if validationErrors != nil {
		c.JSON(http.StatusBadRequest, helper.ResponseError(validationErrors))
		return
	}

	err = h.authService.VerifyUser(verifyRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.ResponseError(err.Error()))
		return
	}

	c.JSON(http.StatusOK, helper.ResponseSuccess(common.ResponseSuccess))
}

func (h *AuthHandler) RefreshVerificationCode(c *gin.Context) {
	var refreshCodeRequest dto.RefreshVerificationCodeRequest
	err := c.ShouldBindJSON(&refreshCodeRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.ResponseError(err.Error()))
		return
	}

	validationErrors := validations.Validate(refreshCodeRequest)
	if validationErrors != nil {
		c.JSON(http.StatusBadRequest, helper.ResponseError(validationErrors))
		return
	}

	code, err := h.authService.RefreshVerificationCode(refreshCodeRequest.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.ResponseError(err.Error()))
		return
	}

	c.JSON(http.StatusOK, helper.ResponseSuccess(code))
}
