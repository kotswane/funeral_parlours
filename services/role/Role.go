package role

import (
	"errors"
	"funeral_parlour/models"
	"funeral_parlour/models/dto"
	"funeral_parlour/repository/role"
	"funeral_parlour/services/helper"
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
	isValidGuid := helper.IsValidUUID(role.FPStatusID)
	if !isValidGuid {
		return models.FPRoles{}, errors.New("invalid premiumId")
	}
	return r.svc.UpdateRole(models.FPRoles{
		FPRoleID:   role.FPStatusID,
		FPRoleName: role.FPRoleName,
	})
}

func (r Role) FindAllRole() ([]models.FPRoles, error) {
	return r.svc.FindAllRoles()
}

func (r Role) DeleteRole(roleId string) error {
	isValidGuid := helper.IsValidUUID(roleId)
	if !isValidGuid {
		return errors.New("invalid premiumId")
	}
	return r.svc.DeleteRole(roleId)
}

func (r Role) FindRole(roleId string) (models.FPRoles, error) {
	isValidGuid := helper.IsValidUUID(roleId)
	if !isValidGuid {
		return models.FPRoles{}, errors.New("invalid premiumId")
	}
	return r.svc.FindRole(roleId)
}
