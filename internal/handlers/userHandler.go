package handler

import (
	"cfd/myapp/internal/common"
	"cfd/myapp/internal/core/domain/dto"
	"cfd/myapp/internal/core/service"
	"cfd/myapp/internal/helper"
	"cfd/myapp/pkg/chacha20"
	"cfd/myapp/pkg/validator"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (h *UserHandler) CreateUser(c *gin.Context) {
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

	userID, err := h.userService.CreateUser(userRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.ResponseError(err.Error()))
		return
	}

	c.JSON(http.StatusCreated, helper.ResponseSuccess(userID))
}

func (h *UserHandler) GetUserByID(c *gin.Context) {
	userIDEncoded := c.Param("id")
	userID, err := chacha20.Decrypt(userIDEncoded)
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.ResponseError(common.ErrBadRequest.Error()))
		return
	}

	userResponse, err := h.userService.GetUserByID(userID)
	if err != nil {
		response := helper.ResponseError(err.Error())
		c.JSON(http.StatusNotFound, response)
		return
	}

	response := helper.ResponseSuccess(userResponse)
	c.JSON(http.StatusOK, response)
}

func (h *UserHandler) UpdateUser(c *gin.Context) {
	userIDEncoded := c.Param("id")
	userID, err := chacha20.Decrypt(userIDEncoded)
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.ResponseError(common.ErrBadRequest.Error()))
		return
	}

	var userRequest dto.UpdateUserRequest
	err = c.ShouldBindJSON(&userRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.ResponseError(err.Error()))
		return
	}

	validationErrors := validations.Validate(userRequest)
	if validationErrors != nil {
		c.JSON(http.StatusBadRequest, helper.ResponseError(validationErrors))
		return
	}

	err = h.userService.UpdateUser(userID, userRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.ResponseError(err.Error()))
		return
	}

	c.JSON(http.StatusOK, helper.ResponseSuccess(userID))
}

func (h *UserHandler) DeleteUser(c *gin.Context) {
	userIDEncoded := c.Param("id")
	userID, err := chacha20.Decrypt(userIDEncoded)
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.ResponseError(common.ErrBadRequest.Error()))
		return
	}

	err = h.userService.DeleteUser(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.ResponseError(err.Error()))
		return
	}

	c.JSON(http.StatusOK, helper.ResponseSuccess(userID))
}
