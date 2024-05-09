package models

type FPUserRole struct {
	FPUserRoleId     string `gorm:"column:fp_usr_role_id;primaryKey"`
	FPUserRoleUsrId  string `gorm:"column:fp_usr_id;not null"`
	FPUserRoleRoleId string `gorm:"column:fp_role_id;not null"`
}

func (m FPUserRole) TableName() string {
	return "fp_usr_role"
}
