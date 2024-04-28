package dto

type FPServiceProvider struct {
	FPServiceProviderID            uint   `json:"fp_service_provider_id"`
	FPServiceProviderName          string `json:"sp_name"`
	FPServiceProviderFspNumber     string `json:"sp_fsp_number"`
	FPServiceProviderAddress       string `json:"sp_address"`
	FPServiceProviderEmail         string `json:"sp_email"`
	FPServiceProviderContactNumber string `json:"sp_contact_number"`
	FPServiceProviderStatus        int    `json:"sp_status"`
}
