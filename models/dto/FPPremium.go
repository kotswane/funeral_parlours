package dto

type FPPremium struct {
	FPPremiumID          uint   `json:"fp_premium_id"`
	FPPremiumName        string `json:"fp_premium_name"`
	FPPremiumDescription string `json:"fp_premium_description"`
	FPPremiumServiceID   int64  `json:"fp_premium_sp_id"`
	FPPremiumStatus      int    `json:"fp_premium_status"`
}
