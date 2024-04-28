package serviceprovider

import (
	"funeral_parlour/models"
	"funeral_parlour/models/dto"
	"funeral_parlour/repository/serviceprovider"
)

func NewServiceProvider(svc serviceprovider.ServiceProviderRepository) ServiceProviderService {
	return &ServiceProvider{svc: svc}
}

type ServiceProvider struct {
	svc serviceprovider.ServiceProviderRepository
}

func (s ServiceProvider) AddServiceProvider(serviceProvider dto.FPServiceProvider) (models.FPServiceProvider, error) {
	return s.svc.AddServiceProvider(models.FPServiceProvider{
		FPServiceProviderName:          serviceProvider.FPServiceProviderName,
		FPServiceProviderFspNumber:     serviceProvider.FPServiceProviderFspNumber,
		FPServiceProviderAddress:       serviceProvider.FPServiceProviderAddress,
		FPServiceProviderEmail:         serviceProvider.FPServiceProviderEmail,
		FPServiceProviderContactNumber: serviceProvider.FPServiceProviderContactNumber,
		FPServiceProviderStatus:        serviceProvider.FPServiceProviderStatus,
	})
}

func (s ServiceProvider) UpdateServiceProvider(serviceProvider dto.FPServiceProvider) (models.FPServiceProvider, error) {
	return s.svc.UpdateServiceProvider(models.FPServiceProvider{
		FPServiceProviderID:            serviceProvider.FPServiceProviderID,
		FPServiceProviderName:          serviceProvider.FPServiceProviderName,
		FPServiceProviderFspNumber:     serviceProvider.FPServiceProviderFspNumber,
		FPServiceProviderAddress:       serviceProvider.FPServiceProviderAddress,
		FPServiceProviderEmail:         serviceProvider.FPServiceProviderEmail,
		FPServiceProviderContactNumber: serviceProvider.FPServiceProviderContactNumber,
		FPServiceProviderStatus:        serviceProvider.FPServiceProviderStatus,
	})
}

func (s ServiceProvider) FindAllServiceProvider() ([]models.FPServiceProvider, error) {
	return s.svc.FindAllServiceProvider()
}

func (s ServiceProvider) DeleteServiceProvider(serviceProviderId int) error {
	return s.svc.DeleteServiceProvider(serviceProviderId)
}

func (s ServiceProvider) FindServiceProvider(serviceProviderId int) (models.FPServiceProvider, error) {
	return s.svc.FindServiceProvider(serviceProviderId)
}
