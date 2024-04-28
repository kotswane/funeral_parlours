package models

type FPServiceProvider struct {
	FPServiceProviderID            uint   `json:"sp_id"`
	FPServiceProviderName          string `json:"sp_name"`
	FPServiceProviderFspNumber     string `json:"sp_fsp_number"`
	FPServiceProviderAddress       string `json:"sp_address"`
	FPServiceProviderEmail         string `json:"sp_email"`
	FPServiceProviderContactNumber string `json:"sp_contact_number"`
	FPServiceProviderStatus        int    `json:"sp_status"`
}

func (m FPServiceProvider) TableName() string {
	return "service_provider"
}
