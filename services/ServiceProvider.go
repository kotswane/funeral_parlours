package services

import (
	"errors"
	"funeral_parlour/models"
	"funeral_parlour/models/dto"
	"funeral_parlour/repository/serviceprovider"
	"funeral_parlour/services/helper"

	"github.com/google/uuid"
)

type ServiceProvider struct {
	repo serviceprovider.ServiceProviderRepository
}

func NewServiceProvider(repo serviceprovider.ServiceProviderRepository) ServiceProviderService {
	return &ServiceProvider{repo: repo}
}

func (s ServiceProvider) AddServiceProvider(serviceProvider dto.FPServiceProvider) (models.FPServiceProvider, error) {
	return s.repo.AddServiceProvider(models.FPServiceProvider{
		FPServiceProviderID:            uuid.NewString(),
		FPServiceProviderName:          serviceProvider.FPServiceProviderName,
		FPServiceProviderFspNumber:     serviceProvider.FPServiceProviderFspNumber,
		FPServiceProviderAddress:       serviceProvider.FPServiceProviderAddress,
		FPServiceProviderEmail:         serviceProvider.FPServiceProviderEmail,
		FPServiceProviderContactNumber: serviceProvider.FPServiceProviderContactNumber,
		FPServiceProviderStatus:        serviceProvider.FPServiceProviderStatus,
	})
}

func (s ServiceProvider) UpdateServiceProvider(serviceProvider dto.FPServiceProvider) (models.FPServiceProvider, error) {
	isValidGuid := helper.IsValidUUID(serviceProvider.FPServiceProviderID)
	if !isValidGuid {
		return models.FPServiceProvider{}, errors.New("invalid ServiceProviderID")
	}
	return s.repo.UpdateServiceProvider(models.FPServiceProvider{
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
	return s.repo.FindAllServiceProvider()
}

func (s ServiceProvider) DeleteServiceProvider(serviceProviderId string) error {
	isValidGuid := helper.IsValidUUID(serviceProviderId)
	if !isValidGuid {
		return errors.New("invalid serviceProviderId")
	}
	return s.repo.DeleteServiceProvider(serviceProviderId)
}

func (s ServiceProvider) FindServiceProvider(serviceProviderId string) (models.FPServiceProvider, error) {
	isValidGuid := helper.IsValidUUID(serviceProviderId)
	if !isValidGuid {
		return models.FPServiceProvider{}, errors.New("invalid serviceProviderId")
	}
	return s.repo.FindServiceProvider(serviceProviderId)
}
