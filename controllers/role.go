package controllers

import (
	"funeral_parlour/services/role"
	"github.com/gin-gonic/gin"
)

type Role struct {
	svc role.RoleService
}

func NewRoleController(svc role.RoleService) Role {
	return Role{svc: svc}
}

func (s *Role) AddRole(c *gin.Context) {

}

func (s *Role) UpdateRole(c *gin.Context) {

}

func (s *Role) FindAllRole(c *gin.Context) {

}

func (s *Role) DeleteRole(c *gin.Context) {

}

func (s *Role) FindRole(c *gin.Context) {

}
