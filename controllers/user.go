package controllers

import (
	"funeral_parlour/controllers/helper"
	"funeral_parlour/models/dto"
	"funeral_parlour/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	svc services.UserService
}

func NewUserController(svc services.UserService) User {
	return User{svc: svc}
}

// AddUser godoc
// @Summary      This API Creates a new User
// @Description  This API Creates a new User
// @Tags         User
// @Accept       json
// @Produce      json
// @Param user body dto.FPUser true "New User"
// @Success      200  {object}  dto.FPUser
// @Failure      400  {object}  dto.FPAPIResponse
// @Failure      404  {object}  dto.FPAPIResponse
// @Failure      500  {object}  dto.FPAPIResponse
// @Router       /user/add [POST]
func (u *User) AddUser(c *gin.Context) {
	var user dto.FPUser

	if err := c.ShouldBindJSON(&user); err != nil {
		_err := helper.ValidateStruct(err)
		if _err != "" {
			c.JSON(http.StatusBadRequest, dto.FPAPIResponse{Message: _err, StatusCode: http.StatusBadRequest})
			return
		}
	}

	resp, err := u.svc.AddUser(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.FPAPIResponse{Message: err.Error(), StatusCode: http.StatusBadRequest})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// UpdateUser godoc
// @Summary      This API Updates a User
// @Description  This API Updates a User
// @Tags         User
// @Accept       json
// @Produce      json
// @Param user body dto.FPUser true "Existing User"
// @Success      200  {object}  dto.FPUser
// @Failure      400  {object}  dto.FPAPIResponse
// @Failure      404  {object}  dto.FPAPIResponse
// @Failure      500  {object}  dto.FPAPIResponse
// @Router       /user/update [POST]
func (u *User) UpdateUser(c *gin.Context) {
	var user dto.FPUser

	if err := c.ShouldBindJSON(&user); err != nil {
		_err := helper.ValidateStruct(err)
		if _err != "" {
			c.JSON(http.StatusBadRequest, dto.FPAPIResponse{Message: _err, StatusCode: http.StatusBadRequest})
			return
		}
	}

	resp, err := u.svc.UpdateUser(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.FPAPIResponse{Message: err.Error(), StatusCode: http.StatusBadRequest})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// FindAllUser godoc
// @Summary      This API Finds All Users
// @Description  This API Finds All Users
// @Tags         User
// @Accept       json
// @Produce      json
// @Success      200  {object}  []models.FPUser
// @Failure      400  {object}  dto.FPAPIResponse
// @Failure      404  {object}  dto.FPAPIResponse
// @Failure      500  {object}  dto.FPAPIResponse
// @Router       /user [GET]
func (u *User) FindAllUser(c *gin.Context) {
	resp, err := u.svc.FindAllUser()
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.FPAPIResponse{Message: err.Error(), StatusCode: http.StatusBadRequest})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// DeleteUser godoc
// @Summary      This API deletes a User
// @Description  This API deletes a User
// @Tags         User
// @Accept       json
// @Produce      json
// @Param userId path string true "The id of a user"
// @Success      200  {object}  dto.FPAPIResponse
// @Failure      400  {object}  dto.FPAPIResponse
// @Failure      404  {object}  dto.FPAPIResponse
// @Failure      500  {object}  dto.FPAPIResponse
// @Router       /user/delete/{userId} [DELETE]
func (u *User) DeleteUser(c *gin.Context) {
	userId := c.Param("userId")
	err := u.svc.DeleteUser(userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.FPAPIResponse{Message: err.Error(), StatusCode: http.StatusBadRequest})
		return
	}
	c.JSON(http.StatusOK, dto.FPAPIResponse{Message: "Success", StatusCode: http.StatusOK})
}

// FindUser godoc
// @Summary      This API Finds User by ID
// @Description  This API Finds User by ID
// @Tags         User
// @Accept       json
// @Produce      json
// @Param userId path string true "The id of a user"
// @Success      200  {object}  models.FPUser
// @Failure      400  {object}  dto.FPAPIResponse
// @Failure      404  {object}  dto.FPAPIResponse
// @Failure      500  {object}  dto.FPAPIResponse
// @Router       /user/find/{userId} [GET]
func (u *User) FindUser(c *gin.Context) {
	userId := c.Param("userId")
	resp, err := u.svc.FindUser(userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.FPAPIResponse{Message: err.Error(), StatusCode: http.StatusBadRequest})
		return
	}
	c.JSON(http.StatusOK, resp)
}
