package userrole

import (
	"funeral_parlour/models"
	"funeral_parlour/models/dto"
)

type UserRoleService interface {
	AddUserRole(userRole dto.FPUserRole) (models.FPUserRole, error)
	UpdateUserRole(userRole dto.FPUserRole) (models.FPUserRole, error)
	FindAllUserRole() ([]models.FPUserRole, error)
	DeleteUserRole(userRoleId string) error
	FindUserRole(userRoleId string) (models.FPUserRole, error)
}
