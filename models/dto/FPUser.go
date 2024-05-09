package dto

type FPUser struct {
	FPUserID          string `json:"usr_id"`
	FPUserFirstName   string `json:"urs_firstname"`
	FPUserMiddleName  string `json:"usr_middle_name"`
	FPUserSurname     string `json:"usr_surname"`
	FPUserIDNumber    string `json:"usr_idnumber"`
	FPUserAddress     string `json:"usr_address"`
	FPUserPhoneNumber string `json:"usr_phonenumber"`
	FPUserEmail       string `json:"usr_email"`
	FPUserPassword    string `json:"usr_password"`
	FPUserRole        string `json:"usr_role"`
	FPUserStatus      string `json:"usr_status"`
	FPUserSpID        string `json:"usr_sp_id"`
	FPUSERUsername    string `json:"usr_username"`
}
