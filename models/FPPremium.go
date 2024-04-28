package models

type FPPremium struct {
	FPPremiumID          uint   `gorm:"column:fp_premium_id:primaryKey;autoIncrement"`
	FPPremiumName        string `gorm:"column:fp_premium_name"`
	FPPremiumDescription string `gorm:"column:fp_premium_description"`
	FPPremiumServiceID   int64  `gorm:"column:fp_premium_sp_id"`
	FPPremiumStatus      int    `gorm:"column:fp_premium_status"`
}

func (m FPPremium) TableName() string {
	return "fp_premiums"
}
