package models

type FPServiceProvider struct {
	FPServiceProviderID            string `gorm:"primaryKey;column:sp_id"`
	FPServiceProviderName          string `gorm:"column:sp_name"`
	FPServiceProviderFspNumber     string `gorm:"column:sp_fsp_number"`
	FPServiceProviderAddress       string `gorm:"column:sp_address"`
	FPServiceProviderEmail         string `gorm:"column:sp_email"`
	FPServiceProviderContactNumber string `gorm:"column:sp_contact_number"`
	FPServiceProviderStatus        string `gorm:"column:sp_status"`
}

func (m FPServiceProvider) TableName() string {
	return "service_provider"
}
