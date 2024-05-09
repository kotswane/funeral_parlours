package models

type FPRoles struct {
	FPRoleID   string `gorm:"column:fp_role_id;primaryKey"`
	FPRoleName string `gorm:"column:fp_role_name"`
}

func (m FPRoles) TableName() string {
	return "fp_roles"
}
