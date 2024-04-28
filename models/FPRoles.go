package models

type FPRoles struct {
	FPRoleID   uint   `gorm:"column:fp_role_id;primaryKey;autoIncrement"`
	FPRoleName string `gorm:"column:fp_role_name"`
}

func (m FPRoles) TableName() string {
	return "fp_roles"
}
