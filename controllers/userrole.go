package controllers

import (
	"funeral_parlour/services/userrole"
	"github.com/gin-gonic/gin"
)

type UserRole struct {
	svc userrole.UserRoleService
}

func NewUserRoleController(svc userrole.UserRoleService) UserRole {
	return UserRole{svc: svc}
}

func (u *UserRole) AddUserRole(c *gin.Context) {

}

func (u *UserRole) UpdateUserRole(c *gin.Context) {

}

func (u *UserRole) FindAllUserRole(c *gin.Context) {

}

func (u *UserRole) DeleteUserRole(c *gin.Context) {

}

func (u *UserRole) FindUserRole(c *gin.Context) {

}
