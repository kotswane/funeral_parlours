package user

import (
	"funeral_parlour/models"
	"funeral_parlour/models/dto"
)

type UserRepository interface {
	AddUser(user models.FPUser) (models.FPUser, error)
	UpdateUser(user models.FPUser, id uint) (models.FPUser, error)
	DeleteUser(userId int) error
	FindAllUser() ([]models.FPUser, error)
	FindUser(userId int) (models.FPUser, error)
	UpdatePassword(password dto.FPPassword) error
}
