package userrole

import (
	"errors"
	"funeral_parlour/models"
	"funeral_parlour/models/dto"
	"funeral_parlour/repository/userrole"
	"funeral_parlour/services/helper"
)

type UserRole struct {
	repo userrole.UserRoleRepository
}

func NewUserRole(repo userrole.UserRoleRepository) UserRoleService {
	return &UserRole{repo: repo}
}

func (u UserRole) AddUserRole(userRole dto.FPUserRole) (models.FPUserRole, error) {
	return u.repo.AddUserRole(models.FPUserRole{
		FPUserRoleUsrId:  userRole.FPUserRoleUsrId,
		FPUserRoleRoleId: userRole.FPUserRoleRoleId,
	})
}

func (u UserRole) UpdateUserRole(userRole dto.FPUserRole) (models.FPUserRole, error) {
	isValidGuid := helper.IsValidUUID(userRole.FPUserRoleId)
	if !isValidGuid {
		return models.FPUserRole{}, errors.New("invalid userRoleId")
	}
	return u.repo.UpdateUserRole(models.FPUserRole{
		FPUserRoleId:     userRole.FPUserRoleId,
		FPUserRoleUsrId:  userRole.FPUserRoleUsrId,
		FPUserRoleRoleId: userRole.FPUserRoleRoleId,
	})
}

func (u UserRole) FindAllUserRole() ([]models.FPUserRole, error) {
	return u.repo.FindAllUserRole()
}

func (u UserRole) DeleteUserRole(userRoleId string) error {
	isValidGuid := helper.IsValidUUID(userRoleId)
	if !isValidGuid {
		return errors.New("invalid userRoleId")
	}
	return u.repo.DeleteUserRole(userRoleId)
}

func (u UserRole) FindUserRole(userRoleId string) (models.FPUserRole, error) {
	isValidGuid := helper.IsValidUUID(userRoleId)
	if !isValidGuid {
		return models.FPUserRole{}, errors.New("invalid userRoleId")
	}
	return u.repo.FindUserRole(userRoleId)
}
