package role

import (
	"funeral_parlour/models"
	"funeral_parlour/models/dto"
	"funeral_parlour/repository/role"
)

type Role struct {
	svc role.RoleRepository
}

func NewRole(svc role.RoleRepository) RoleService {
	return &Role{svc: svc}
}

func (r Role) AddRole(role dto.FPRoles) (models.FPRoles, error) {
	return r.svc.UpdateRole(models.FPRoles{
		FPRoleName: role.FPRoleName,
	})
}

func (r Role) UpdateRole(role dto.FPRoles) (models.FPRoles, error) {
	return r.svc.UpdateRole(models.FPRoles{
		FPRoleID:   role.FPStatusID,
		FPRoleName: role.FPRoleName,
	})
}

func (r Role) FindAllRole() ([]models.FPRoles, error) {
	return r.svc.FindAllRoles()
}

func (r Role) DeleteRole(roleId int) error {
	return r.svc.DeleteRole(roleId)
}

func (r Role) FindRole(roleId int) (models.FPRoles, error) {
	return r.svc.FindRole(roleId)
}
