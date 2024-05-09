package userrole

import (
	"funeral_parlour/models"
)

type UserRoleRepository interface {
	AddUserRole(userRole models.FPUserRole) (models.FPUserRole, error)
	UpdateUserRole(userRole models.FPUserRole) (models.FPUserRole, error)
	FindAllUserRole() ([]models.FPUserRole, error)
	DeleteUserRole(userRoleId string) error
	FindUserRole(userRoleId string) (models.FPUserRole, error)
}
