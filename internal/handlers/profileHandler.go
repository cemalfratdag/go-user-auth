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

type ProfileHandler struct {
	profileService service.ProfileService
}

func NewProfileHandler(profileService service.ProfileService) ProfileHandler {
	return ProfileHandler{
		profileService: profileService,
	}
}

func (h *ProfileHandler) ViewProfile(c *gin.Context) {
	rawUserID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusInternalServerError, helper.ResponseError(common.ErrMissingContext))
		return
	}

	userID, ok := rawUserID.(int)
	if !ok {
		c.JSON(http.StatusInternalServerError, helper.ResponseError(common.ErrInterval))
		return
	}

	userResponse, err := h.profileService.ViewProfile(userID)
	if err != nil {
		response := helper.ResponseError(err.Error())
		c.JSON(http.StatusNotFound, response)
		return
	}

	response := helper.ResponseSuccess(userResponse)
	c.JSON(http.StatusOK, response)
}

func (h *ProfileHandler) UpdateProfile(c *gin.Context) {
	rawUserID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusInternalServerError, helper.ResponseError(common.ErrMissingContext))
		return
	}

	userID, ok := rawUserID.(int)
	if !ok {
		c.JSON(http.StatusInternalServerError, helper.ResponseError(common.ErrInterval))
		return
	}

	var userRequest dto.UpdateProfileRequest
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

	err = h.profileService.UpdateProfile(userID, userRequest)
	if err != nil {
		response := helper.ResponseError(err.Error())
		c.JSON(http.StatusNotFound, response)
		return
	}

	c.JSON(http.StatusOK, helper.ResponseSuccess(common.ResponseSuccess))
}

func (h *ProfileHandler) DeleteProfile(c *gin.Context) {
	rawUserID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusInternalServerError, helper.ResponseError(common.ErrMissingContext))
		return
	}

	userID, ok := rawUserID.(int)
	if !ok {
		c.JSON(http.StatusInternalServerError, helper.ResponseError(common.ErrInterval))
		return
	}

	err := h.profileService.DeleteProfile(userID)
	if err != nil {
		response := helper.ResponseError(err.Error())
		c.JSON(http.StatusNotFound, response)
		return
	}

	c.JSON(http.StatusOK, helper.ResponseSuccess(common.ResponseSuccess))
}

func (h *ProfileHandler) ForgotPassword(c *gin.Context) {
	var forgotPasswordRequest dto.ForgotPasswordRequest
	err := c.ShouldBindJSON(&forgotPasswordRequest)
	if err != nil {
		c.JSON(http.StatusOK, helper.ResponseSuccess(common.ResponseForgotPassword))
		return
	}

	validationErrors := validations.Validate(forgotPasswordRequest)
	if validationErrors != nil {
		c.JSON(http.StatusOK, helper.ResponseSuccess(validationErrors))
		return
	}

	code, err := h.profileService.ForgotPassword(forgotPasswordRequest.Email)
	if err != nil {
		c.JSON(http.StatusOK, helper.ResponseSuccess(common.ResponseForgotPassword))
		return
	}

	c.JSON(http.StatusOK, helper.ResponseSuccess(code))
}

func (h *ProfileHandler) ResetPassword(c *gin.Context) {
	var resetPasswordRequest dto.ResetPasswordRequest
	err := c.ShouldBindJSON(&resetPasswordRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.ResponseError(err.Error()))
		return
	}

	validationErrors := validations.Validate(resetPasswordRequest)
	if validationErrors != nil {
		c.JSON(http.StatusBadRequest, helper.ResponseError(validationErrors))
		return
	}

	err = h.profileService.ResetPassword(resetPasswordRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.ResponseError(err.Error()))
		return
	}

	c.JSON(http.StatusOK, helper.ResponseSuccess(common.ResponseSuccess))
}

func (h *ProfileHandler) ChangePassword(c *gin.Context) {
	rawUserID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusInternalServerError, helper.ResponseError(common.ErrMissingContext))
		return
	}

	userID, ok := rawUserID.(int)
	if !ok {
		c.JSON(http.StatusInternalServerError, helper.ResponseError(common.ErrInterval))
		return
	}

	var changePasswordRequest dto.ChangePasswordRequest
	err := c.ShouldBindJSON(&changePasswordRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.ResponseError(err.Error()))
		return
	}

	validationErrors := validations.Validate(changePasswordRequest)
	if validationErrors != nil {
		c.JSON(http.StatusBadRequest, helper.ResponseError(validationErrors))
		return
	}

	err = h.profileService.ChangePassword(changePasswordRequest, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.ResponseError(err.Error()))
		return
	}

	c.JSON(http.StatusOK, helper.ResponseSuccess(userID))
}
