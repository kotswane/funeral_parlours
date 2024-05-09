package userrole

import (
	"funeral_parlour/models"

	"gorm.io/gorm"
)

type UserRole struct {
	db *gorm.DB
}

func NewUserRole(db *gorm.DB) UserRoleRepository {
	return &UserRole{db: db}
}

func (u UserRole) AddUserRole(userRole models.FPUserRole) (models.FPUserRole, error) {
	results := u.db.Create(userRole)
	if results.Error != nil {
		return models.FPUserRole{}, results.Error
	}
	return userRole, nil
}

func (u UserRole) UpdateUserRole(userRole models.FPUserRole) (models.FPUserRole, error) {
	if err := u.db.Model(&models.FPUserRole{}).Where("fp_usr_id = ?", userRole.FPUserRoleUsrId).Update("fp_usr_role_id", userRole.FPUserRoleUsrId).Error; err != nil {
		return models.FPUserRole{}, err
	}
	return userRole, nil
}

func (u UserRole) FindAllUserRole() ([]models.FPUserRole, error) {
	var userRoleList []models.FPUserRole
	if err := u.db.Find(&userRoleList).Error; err != nil {
		return userRoleList, err
	}
	return userRoleList, nil
}

func (u UserRole) DeleteUserRole(userRoleId string) error {
	if err := u.db.Where("fp_usr_role_id = ?", userRoleId).Delete(&models.FPUserRole{}).Error; err != nil {
		return err
	}
	return nil
}

func (u UserRole) FindUserRole(userRoleId string) (models.FPUserRole, error) {
	var userRole models.FPUserRole
	if err := u.db.Where("fp_usr_role_id = ?", userRoleId).First(&userRole).Error; err != nil {
		return userRole, err
	}
	return userRole, nil
}
