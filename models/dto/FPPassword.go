package dto

type FPPassword struct {
	UserID          uint   `json:"userId"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmPassword"`
}
