package userrole

import (
	"funeral_parlour/models"
)

type UserRoleRepository interface {
	AddUserRole(userRole models.FPUserRole) (models.FPUserRole, error)
	UpdateUserRole(userRole models.FPUserRole) (models.FPUserRole, error)
	FindAllUserRole() ([]models.FPUserRole, error)
	DeleteUserRole(userRoleId int) error
	FindUserRole(userRoleId int) (models.FPUserRole, error)
}
