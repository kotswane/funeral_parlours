package controllers

import (
	"funeral_parlour/services/user"
	"github.com/gin-gonic/gin"
)

type User struct {
	svc user.UserService
}

func NewUserController(svc user.UserService) User {
	return User{svc: svc}
}

func (u *User) AddUser(c *gin.Context) {

}

func (u *User) UpdateUser(c *gin.Context) {

}

func (u *User) FindAllUser(c *gin.Context) {

}

func (u *User) DeleteUser(c *gin.Context) {

}

func (u *User) FindUser(c *gin.Context) {

}
