package userrole

import (
	"funeral_parlour/models"
	"funeral_parlour/models/dto"
	"funeral_parlour/repository/userrole"
)

type UserRole struct {
	svc userrole.UserRoleRepository
}

func NewUserRole(svc userrole.UserRoleRepository) UserRoleService {
	return &UserRole{svc: svc}
}

func (u UserRole) AddUserRole(userRole dto.FPUserRole) (models.FPUserRole, error) {
	return u.svc.AddUserRole(models.FPUserRole{
		FPUserRoleUsrId:  userRole.FPUserRoleUsrId,
		FPUserRoleRoleId: userRole.FPUserRoleRoleId,
	})
}

func (u UserRole) UpdateUserRole(userRole dto.FPUserRole) (models.FPUserRole, error) {
	return u.svc.UpdateUserRole(models.FPUserRole{
		FPUserRoleId:     userRole.FPUserRoleId,
		FPUserRoleUsrId:  userRole.FPUserRoleUsrId,
		FPUserRoleRoleId: userRole.FPUserRoleRoleId,
	})
}

func (u UserRole) FindAllUserRole() ([]models.FPUserRole, error) {
	return u.svc.FindAllUserRole()
}

func (u UserRole) DeleteUserRole(userRoleId int) error {
	return u.svc.DeleteUserRole(userRoleId)
}

func (u UserRole) FindUserRole(userRoleId int) (models.FPUserRole, error) {
	return u.svc.FindUserRole(userRoleId)
}
