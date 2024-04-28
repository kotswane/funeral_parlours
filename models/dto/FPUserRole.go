package dto

type FPUserRole struct {
	FPUserRoleId     uint  `json:"fp_usr_role_id"`
	FPUserRoleUsrId  int64 `json:"fp_usr_id"`
	FPUserRoleRoleId uint  `json:"fp_role_id"`
}
