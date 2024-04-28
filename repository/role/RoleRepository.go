package role

import "funeral_parlour/models"

type RoleRepository interface {
	AddRole(role models.FPRoles) (models.FPRoles, error)
	UpdateRole(role models.FPRoles) (models.FPRoles, error)
	FindAllRoles() ([]models.FPRoles, error)
	DeleteRole(statusId int) error
	FindRole(statusId int) (models.FPRoles, error)
}
