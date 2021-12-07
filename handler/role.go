package handler

import (
	"backend-simple-pos/helper"
	"backend-simple-pos/role"
	"github.com/gin-gonic/gin"
	"net/http"
)

type RoleHandler struct {
	roleService role.Service
}

func NewRoleHandler(roleService role.Service) *RoleHandler {
	return &RoleHandler{roleService}
}

func (h *RoleHandler) GetRoles(c *gin.Context){
	roles,err := h.roleService.GetAllRoles()
	if err != nil {
		response := helper.APIResponse("Register account gagal", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Account has been registered", http.StatusOK, "success", roles)
	c.JSON(http.StatusOK, response)
}