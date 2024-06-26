package models

type FPUser struct {
	FPUserID          string `gorm:"column:usr_id;primaryKey"`
	FPUserFirstName   string `gorm:"column:urs_firstname"`
	FPUserMiddleName  string `gorm:"column:usr_middle_name"`
	FPUserSurname     string `gorm:"column:usr_surname"`
	FPUserIDNumber    string `gorm:"column:usr_idnumber"`
	FPUserAddress     string `gorm:"column:usr_address"`
	FPUserPhoneNumber string `gorm:"column:usr_phonenumber"`
	FPUserEmail       string `gorm:"column:usr_email"`
	FPUserPassword    string `gorm:"column:usr_password"`
	FPUserRole        string `gorm:"column:usr_role"`
	FPUserStatus      string `gorm:"column:usr_status"`
	FPUserSpID        string `gorm:"column:usr_sp_id"`
	FPUSERUsername    string `gorm:"column:usr_username"`
}

func (m FPUser) TableName() string {
	return "fp_users"
}
