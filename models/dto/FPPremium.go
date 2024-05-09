package dto

type FPCreatePremium struct {
	FPPremiumName        string `json:"fp_premium_amount" validate:"required"`
	FPPremiumDescription string `json:"fp_premium_description" validate:"required"`
	FPPremiumServiceID   string `json:"fp_premium_sp_id" validate:"required,uuid4"`
	FPPremiumStatus      string `json:"fp_premium_status" validate:"required,uuid4"`
}

type FPUpdatePremium struct {
	FPPremiumID          string `json:"fp_premium_id" validate:"required,uuid4"`
	FPPremiumName        string `json:"fp_premium_amount" validate:"required"`
	FPPremiumDescription string `json:"fp_premium_description" validate:"required"`
	FPPremiumServiceID   string `json:"fp_premium_sp_id" validate:"required,uuid4"`
	FPPremiumStatus      string `json:"fp_premium_status" validate:"required,uuid4"`
}
