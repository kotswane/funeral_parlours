package role

import (
	"funeral_parlour/models"

	"gorm.io/gorm"
)

type Role struct {
	db *gorm.DB
}

func NewRole(db *gorm.DB) RoleRepository {
	return &Role{db: db}
}

func (r Role) AddRole(role models.FPRoles) (models.FPRoles, error) {
	response := r.db.Create(role)
	if response.Error != nil {
		return models.FPRoles{}, response.Error
	}
	return role, nil
}

func (r Role) UpdateRole(role models.FPRoles) (models.FPRoles, error) {

	if err := r.db.Model(&models.FPRoles{}).Where("fp_role_id = ?", role.FPRoleID).Update("fp_role_name", role.FPRoleName).Error; err != nil {
		return models.FPRoles{}, err
	}
	return role, nil
}

func (r Role) FindAllRoles() ([]models.FPRoles, error) {
	var roleList []models.FPRoles
	if err := r.db.Find(&roleList).Error; err != nil {
		return roleList, err
	}
	return roleList, nil
}

func (r Role) DeleteRole(statusId string) error {
	if err := r.db.Where("fp_role_id = ?", statusId).Delete(&models.FPRoles{}).Error; err != nil {
		return err
	}
	return nil
}

func (r Role) FindRole(statusId string) (models.FPRoles, error) {
	var role models.FPRoles
	if err := r.db.Where("fp_role_id = ?", statusId).First(&role).Error; err != nil {
		return models.FPRoles{}, err
	}
	return role, nil
}
