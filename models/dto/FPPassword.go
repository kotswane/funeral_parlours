package dto

type FPPassword struct {
	UserID          string `json:"userId"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmPassword"`
}
