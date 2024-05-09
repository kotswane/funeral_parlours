package models

type FPPremium struct {
	FPPremiumID          string `gorm:"column:fp_premium_id;primaryKey"`
	FPPremiumName        string `gorm:"column:fp_premium_name"`
	FPPremiumDescription string `gorm:"column:fp_premium_description"`
	FPPremiumServiceID   string `gorm:"column:fp_premium_sp_id"`
	FPPremiumStatus      string `gorm:"column:fp_premium_status"`
}

func (m FPPremium) TableName() string {
	return "fp_premiums"
}
