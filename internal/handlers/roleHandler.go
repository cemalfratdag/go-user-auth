package handler

import (
	"cfd/myapp/internal/common"
	"cfd/myapp/internal/core/domain/dto"
	"cfd/myapp/internal/core/service"
	"cfd/myapp/internal/helper"
	"cfd/myapp/pkg/chacha20"
	validations "cfd/myapp/pkg/validator"
	"github.com/gin-gonic/gin"
	"net/http"
)

type RoleHandler struct {
	roleService service.RoleService
}

func NewRoleHandler(roleService service.RoleService) *RoleHandler {
	return &RoleHandler{
		roleService: roleService,
	}
}

func (h *RoleHandler) CreateRole(c *gin.Context) {
	var roleRequest dto.RoleRequest
	err := c.ShouldBindJSON(&roleRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.ResponseError(err.Error()))
		return
	}

	validationErrors := validations.Validate(roleRequest)
	if validationErrors != nil {
		c.JSON(http.StatusBadRequest, helper.ResponseError(validationErrors))
		return
	}

	roleID, err := h.roleService.CreateRole(roleRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.ResponseError(err.Error()))
		return
	}

	c.JSON(http.StatusCreated, helper.ResponseSuccess(roleID))
}

func (h *RoleHandler) GetRoleByID(c *gin.Context) {
	roleIDEncoded := c.Param("id")
	roleID, err := chacha20.Decrypt(roleIDEncoded)
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.ResponseError(common.ErrBadRequest.Error()))
		return
	}

	roleResponse, err := h.roleService.GetRoleByID(roleID)
	if err != nil {
		response := helper.ResponseError(err.Error())
		c.JSON(http.StatusNotFound, response)
		return
	}

	response := helper.ResponseSuccess(roleResponse)
	c.JSON(http.StatusOK, response)
}

func (h *RoleHandler) UpdateRole(c *gin.Context) {
	roleIDEncoded := c.Param("id")
	roleID, err := chacha20.Decrypt(roleIDEncoded)
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.ResponseError(common.ErrBadRequest.Error()))
		return
	}

	var roleRequest dto.RoleRequest
	err = c.ShouldBindJSON(&roleRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.ResponseError(err.Error()))
		return
	}

	validationErrors := validations.Validate(roleRequest)
	if validationErrors != nil {
		c.JSON(http.StatusBadRequest, helper.ResponseError(validationErrors))
		return
	}

	err = h.roleService.UpdateRole(roleID, roleRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.ResponseError(err.Error()))
		return
	}

	c.JSON(http.StatusOK, helper.ResponseSuccess(roleID))
}

func (h *RoleHandler) DeleteRole(c *gin.Context) {
	roleIDEncoded := c.Param("id")
	roleID, err := chacha20.Decrypt(roleIDEncoded)
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.ResponseError(common.ErrBadRequest.Error()))
		return
	}

	err = h.roleService.DeleteRole(roleID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.ResponseError(err.Error()))
		return
	}

	c.JSON(http.StatusOK, helper.ResponseSuccess(roleID))
}
