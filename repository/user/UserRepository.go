package user

import (
	"funeral_parlour/models"
	"funeral_parlour/models/dto"
)

type UserRepository interface {
	AddUser(user models.FPUser) (models.FPUser, error)
	UpdateUser(user models.FPUser, id string) (models.FPUser, error)
	DeleteUser(userId string) error
	FindAllUser() ([]models.FPUser, error)
	FindUser(userId string) (models.FPUser, error)
	UpdatePassword(password dto.FPPassword) error
}
