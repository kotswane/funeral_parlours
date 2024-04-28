package models

type FPUserRole struct {
	FPUserRoleId     uint  `gorm:"column:fp_usr_role_id;primaryKey;autoIncrement"`
	FPUserRoleUsrId  int64 `gorm:"column:fp_usr_id;not null"`
	FPUserRoleRoleId uint  `gorm:"column:fp_role_id;not null"`
}

func (m FPUserRole) TableName() string {
	return "fp_usr_role"
}
