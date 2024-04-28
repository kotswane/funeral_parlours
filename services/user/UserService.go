package user

import (
	"funeral_parlour/models"
	"funeral_parlour/models/dto"
)

type UserService interface {
	AddUser(user dto.FPUser) (models.FPUser, error)
	UpdateUser(user dto.FPUser) (models.FPUser, error)
	FindAllUser() ([]models.FPUser, error)
	DeleteUser(userId int) error
	FindUser(userId int) (models.FPUser, error)
}
