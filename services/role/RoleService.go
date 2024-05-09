package role

import (
	"funeral_parlour/models"
	"funeral_parlour/models/dto"
)

type RoleService interface {
	AddRole(role dto.FPRoles) (models.FPRoles, error)
	UpdateRole(role dto.FPRoles) (models.FPRoles, error)
	FindAllRole() ([]models.FPRoles, error)
	DeleteRole(roleId string) error
	FindRole(roleId string) (models.FPRoles, error)
}
